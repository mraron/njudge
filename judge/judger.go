package judge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/utils/problems/polygon"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Test   string
	Status problems.Status
	Done   bool
	Time   time.Time
}

type Callbacker interface {
	Callback(string, problems.Status, bool) error
}

type HTTPCallback struct {
	url string
}

func NewHTTPCallback(url string) HTTPCallback {
	return HTTPCallback{url}
}

func (h HTTPCallback) Callback(test string, status problems.Status, done bool) error {
	raw := Status{test, status, done, time.Now()}

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
	log.Print(p.(polygon.Problem).Judging.TestSet[0])
	id, err := uuid.NewV4()
	if err != nil {
		log.Print(err)
		return err
	}
	log.Print("itt?")
	f, err := os.Create("/tmp/judge_" + id.String())
	if err != nil {
		log.Print(err)
		return err
	}

	_, err = f.Write([]byte(src))
	if err != nil {
		log.Print(err)
		return err
	}

	f.Close()

	f, err = os.Open("/tmp/judge_" + id.String())
	if err != nil {
		log.Print(err)
		return err
	}

	stderr := bytes.Buffer{}
	bin, err := p.Compile(sandbox, lang, f, &stderr)

	if err != nil {
		st := problems.Status{}
		st.Compiled = false
		st.CompilerOutput = stderr.String()

		log.Print(err)
		return c.Callback("", st, true)
	}

	var (
		testNotifier   = make(chan string)
		statusNotifier = make(chan problems.Status)
		ran            = make(chan bool)
		st             problems.Status
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

			err = c.Callback(test, status, false)

			if err != nil {
				log.Print(err)
				return err
			}
		case <-ran:
			run = false
			break

		}
	}

	os.Remove("/tmp/judge_" + id.String())

	return c.Callback("", st, true)
}
