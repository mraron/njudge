package judge

import (
	"bytes"
	"log"
	"time"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"go.uber.org/multierr"
)

type Status struct {
	Test   string
	Status problems.Status
	Done   bool
	Time   time.Time
}

//@TODO add separated logger for problem

func truncate(s string, to int) string {
	if len(s) < to {
		return s
	}

	return s[:to-1] + "..."
}

func Judge(logger *log.Logger, p problems.Problem, src []byte, lang language.Language, sandboxProvider *language.SandboxProvider, c Callbacker) (st problems.Status, err error) {
	logger.Print("Started Judge")

	//@TODO do smth better
	sandboxes := language.NewSandboxProvider()
	for i := 0; i < 2; i++ {
		var sandbox language.Sandbox
		sandbox, err = sandboxProvider.Get()
		if err != nil {
			logger.Print("Error while getting sandbox: ", err)
			return
		}
		defer sandboxProvider.Put(sandbox)

		err = sandbox.Init(logger)
		if err != nil {
			return
		}
		sandboxes.Put(sandbox)

		defer func(sandbox language.Sandbox) {
			err = multierr.Append(err, sandbox.Cleanup())
		}(sandbox)
	}

	logger.Print("Getting tasktype")
	tt := p.GetTaskType()

	stderr := bytes.Buffer{}

	logger.Print("Trying to compile")

	compileSandbox := sandboxes.MustGet()
	bin, err := tt.Compile(p, compileSandbox, lang, bytes.NewReader(src), &stderr)
	sandboxes.Put(compileSandbox)

	if err != nil {
		logger.Print("Compile got error: ", err)
		st.Compiled = false
		st.CompilerOutput = err.Error() + "\n" + truncate(stderr.String(), 1024)
		err = multierr.Append(err, c.Callback("", st, true))
		return
	}
	st.Compiled = true

	var (
		testNotifier   = make(chan string)
		statusNotifier = make(chan problems.Status)
		ran            = make(chan bool)
		errRun         error
	)

	go func() {
		st, errRun = tt.Run(p, sandboxes, lang, bin, testNotifier, statusNotifier)
		ran <- true
		close(ran)
	}()

	run := true
	for run {
		select {
		case test, ok := <-testNotifier:
			if ok {
				status := <-statusNotifier

				err = c.Callback(test, status, false)
				if err != nil {
					logger.Print("Error while calling callback", err)
					return
				}
			}
		case <-ran:
			run = false
			break

		}
	}
	err = multierr.Combine(err, errRun)

	if err == nil {
		logger.Print("Successful judging! removing tempfile and calling back for the last time...")
	} else {
		logger.Print("Got error! removing tempfile and calling back for the last time... error is", err)
	}

	return
}
