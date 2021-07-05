package judge

import (
	"bytes"
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/problems"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Status struct {
	Test   string
	Status problems.Status
	Done   bool
	Time   time.Time
}

//@TODO add separated logger for problem

func Judge(logger *log.Logger, p problems.Problem, src []byte, lang language.Language, sandboxProvider *language.SandboxProvider, c Callbacker) error {
	logger.Print("Started Judge")

	f, err := ioutil.TempFile("/tmp", "judge_*")
	if err != nil {
		logger.Print("Error while creating:", err)
		return err
	}

	_, err = f.Write(src)
	if err != nil {
		logger.Print("Error while writing data:", err)
		return err
	}

	fname := f.Name()

	f.Close()

	f, err = os.Open(fname)
	if err != nil {
		logger.Print("Error while reopening file:", err)
		return err
	}
	defer f.Close()

	//@TODO do smth better
	sandboxes := language.NewSandboxProvider()
	for i := 0; i < 2; i++ {
		sandbox, err := sandboxProvider.Get()
		if err != nil {
			logger.Print("Error while getting sandbox: ", err)
			return err
		}
		defer sandboxProvider.Put(sandbox)

		sandbox.Init(logger)
		sandboxes.Put(sandbox)

		defer sandbox.Cleanup()
	}

	logger.Print("Getting tasktype")
	tt := problems.GetTaskType(p.TaskTypeName())

	stderr := bytes.Buffer{}

	logger.Print("Trying to compile")

	compileSandbox := sandboxes.MustGet()
	bin, err := tt.Compile(p, compileSandbox, lang, f, &stderr)
	sandboxes.Put(compileSandbox)

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
		st, err = tt.Run(p, sandboxes, lang, bin, testNotifier, statusNotifier)
		ran <- true
		close(ran)
	}()

	run := true
	for run {
		select {
		case test, ok := <-testNotifier:
			if ok {
				status := <-statusNotifier

				err2 := c.Callback(test, status, false)

				if err2 != nil {
					logger.Print("Error while calling callback", err2)
					return err
				}
			}
		case <-ran:
			run = false
			break

		}
	}

	if err == nil {
		logger.Print("Successful judging! removing tempfile and calling back for the last time...")
	} else {
		logger.Print("Got error! removing tempfile and calling back for the last time... error is", err)
	}

	os.Remove(fname)

	return c.Callback("", st, true)
}
