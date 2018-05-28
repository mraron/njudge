package judge

import (
	"github.com/mraron/njudge/utils/problems"
	"net/http"
	"encoding/json"
	"bytes"
	"time"
	"errors"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"github.com/satori/go.uuid"
	"os"
)

type Status struct {
	Test string
	Status problems.Status
	Time time.Time
}

type Callbacker interface {
	Callback(string, problems.Status) error
}

type HTTPCallback struct {
	url string
}

func NewHTTPCallback(url string) HTTPCallback {
	return HTTPCallback{url}
}

func (h HTTPCallback) Callback(test string, status problems.Status) error {
	raw := Status{ test ,status, time.Now()}

	buf := &bytes.Buffer{}

	data := json.NewEncoder(buf)
	err := data.Encode(raw)
	if err != nil {
		return err
	}

	resp, err := http.Post(h.url, "application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprint("Callback error: ", resp.Status, resp.Body))
	}

	return nil
}

func Judge(p problems.Problem, src string, lang language.Language, sandbox language.Sandbox, c Callbacker) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	f, err := os.Create("/tmp/judge_"+id.String())
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(src))
	if err != nil {
		return err
	}

	f.Close()

	f, err = os.Open("/tmp/judge_"+id.String())
	if err != nil {
		return err
	}


	stderr := bytes.Buffer{}
	bin, err := p.Compile(sandbox, lang, f, &stderr)

	if err != nil {
		st := problems.Status{}
		st.Compiled = false
		st.CompilerOutput = stderr.String()

		return c.Callback("", st)
	}

	var (
		testNotifier = make(chan string)
		statusNotifier = make(chan problems.Status)
		ran = make(chan bool)
		st problems.Status
	)

	go func() {
		st, err = p.Run(sandbox, lang, bin, testNotifier, statusNotifier)
		ran <- true
		close(ran)
	}()

	run := true
	for run {
		select {
		case test := <-testNotifier:
			status := <-statusNotifier

			err = c.Callback(test, status)

			if err != nil {
				return err
			}
		case <-ran:
			run = false
			break

		}
	}


	os.Remove("/tmp/judge_"+id.String())

	return c.Callback("", st)
}