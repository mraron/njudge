package judge

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

type customEnqueuer struct {
	ch chan Response
}

func newCustomEnqueuer(r []Response) *customEnqueuer {
	res := &customEnqueuer{make(chan Response, len(r))}
	for i := range r {
		res.ch <- r[i]
	}
	close(res.ch)

	return res
}

func (c customEnqueuer) Enqueue(context.Context, Submission) (<-chan Response, error) {
	return c.ch, nil
}

func TestPostJudgeStream(t *testing.T) {
	resps := []Response{
		{Test: "1"},
		{Test: "423"},
		{Test: "2", Done: true},
	}

	s := NewHTTPServer(HTTPConfig{"", ""}, newCustomEnqueuer(resps))
	
	buf := bytes.Buffer{}

	if err := json.NewEncoder(&buf).Encode(Submission{Stream: true}); err != nil {
		t.Error(err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", &buf)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/judge")

	if err := s.postJudge(c); err != nil {
		t.Error(err)
	}

	sc := bufio.NewScanner(rec.Body)
	ind := 0
	for sc.Scan() {
		status := SubmissionStatus{}
		if err := json.Unmarshal([]byte(sc.Text()), &status); err != nil {
			return
		}

		if resps[ind].Test != status.Test {
			t.Errorf("%s != %s", resps[ind].Test, status.Test)
		}

		ind ++
	}

	if ind != len(resps) {
		t.Error("wrong number")
	}
}