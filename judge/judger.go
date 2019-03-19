package judge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"github.com/satori/go.uuid"
	"log"
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

//@TODO add separated logger for problem

func Judge(logger *log.Logger, p problems.Problem, src []byte, lang language.Language, sandbox language.Sandbox, c Callbacker) error {
	logger.Print("Started Judge")

	id, err := uuid.NewV4()
	if err != nil {
		logger.Print("Error while generating uuid")
		return err
	}

	filename := "/tmp/judge_" + id.String()
	logger.Print("Creating tempfile", filename)

	f, err := os.Create(filename)
	if err != nil {
		logger.Print("Error while creating:", err)
		return err
	}

	_, err = f.Write([]byte(src))
	if err != nil {
		logger.Print("Error while writing data:", err)
		return err
	}

	f.Close()

	logger.Print("Opening tempfile")
	f, err = os.Open("/tmp/judge_" + id.String())
	if err != nil {
		logger.Print("Error while opening:", err)
		return err
	}

	sandbox.Init(logger)
	defer sandbox.Cleanup()

	logger.Print("Getting tasktype")
	tt := problems.GetTaskType(p.TaskTypeName())

	stderr := bytes.Buffer{}

	logger.Print("Trying to compile")
	bin, err := tt.Compile(p, sandbox, lang, f, &stderr)

	if err != nil {
		logger.Print("Compile got error: ", err)
		st := problems.Status{}
		st.Compiled = false
		st.CompilerOutput = err.Error() + stderr.String()
		return c.Callback("", st, true)
	}

	var (
		testNotifier   = make(chan string)
		statusNotifier = make(chan problems.Status)
		ran            = make(chan bool)
		st             problems.Status
	)

	go func() {
		st, err = tt.Run(p, sandbox, lang, bin, testNotifier, statusNotifier)
		ran <- true
		close(ran)
	}()

	run := true
	for run {
		select {
		case test := <-testNotifier:
			status := <-statusNotifier

			err2 := c.Callback(test, status, false)

			if err2 != nil {
				logger.Print("Error while calling callback", err2)
				return err
			}
		case <-ran:
			run = false
			break

		}
	}

	if err == nil {
		logger.Print("Succesful judging! removing tempfile and calling back for the last time...")
	} else {
		logger.Print("Got error! removing tempfile and calling back for the last time... error is", err)
	}

	os.Remove("/tmp/judge_" + id.String())

	return c.Callback("", st, true)
}
