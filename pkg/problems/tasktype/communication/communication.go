package communication

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/tasktype/batch"
	"github.com/mraron/njudge/pkg/problems/tasktype/stub"
	"go.uber.org/multierr"
)

type Communication struct {
	stub.Stub

	RunInteractorF func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (language.Status, error)
	RunUserF       func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (language.Status, error)
}

func New() *Communication {
	c := &Communication{
		stub.New(),
		func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (language.Status, error) {
			return rc.Store["interactorSandbox"].(language.Sandbox).Stdin(utoi).Stdout(itou).TimeLimit(2*tc.TimeLimit).MemoryLimit(1024*1024*1024).Run("interactor inp out", true)
		},
		func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (language.Status, error) {
			return rc.Lang.Run(rc.Sandbox, bytes.NewReader(rc.Binary), itou, utoi, tc.TimeLimit, tc.MemoryLimit)
		},
	}

	c.Batch.InitF = func(rc *batch.RunContext) error {
		interactorSandbox, err := rc.SandboxProvider.Get()
		if err != nil {
			return err
		}

		interactorPath := ""
		for _, f := range rc.Problem.Files() {
			if f.Role == "interactor" {
				interactorPath = f.Path
			}
		}

		f, err := os.Open(interactorPath)
		if err != nil {
			return err
		}
		defer f.Close()

		interactorSandbox.CreateFile("interactor", f)
		interactorSandbox.MakeExecutable("interactor")

		rc.Store["interactorSandbox"] = interactorSandbox
		return nil
	}

	c.Batch.RunF = func(rc *batch.RunContext, g *problems.Group, tc *problems.Testcase) (language.Status, error) {
		interactorSandbox := rc.Store["interactorSandbox"].(language.Sandbox)
		testLocation, answerLocation := tc.InputPath, tc.AnswerPath

		testFile, err := os.Open(testLocation)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return language.Status{}, err
		}
		defer testFile.Close()

		err = interactorSandbox.CreateFile("inp", testFile)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return language.Status{}, err
		}

		answerFile, err := os.Open(answerLocation)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return language.Status{}, err
		}
		defer answerFile.Close()

		os.Remove("/tmp/fifo1" + interactorSandbox.Id())
		os.Remove("/tmp/fifo2" + interactorSandbox.Id())

		err = syscall.Mkfifo(filepath.Join("/tmp", "fifo1"+interactorSandbox.Id()), 0766)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return language.Status{}, err
		}

		err = syscall.Mkfifo(filepath.Join("/tmp", "fifo2"+interactorSandbox.Id()), 0766)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return language.Status{}, err
		}

		fifo1, err := os.OpenFile(filepath.Join("/tmp", "fifo1"+interactorSandbox.Id()), os.O_RDWR, 0766)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return language.Status{}, err
		}
		defer fifo1.Close()

		fifo2, err := os.OpenFile(filepath.Join("/tmp", "fifo2"+interactorSandbox.Id()), os.O_RDWR, 0766)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return language.Status{}, err
		}
		defer fifo2.Close()

		done := make(chan int, 1)

		var errInteractor error
		go func() {
			_, errInteractor = c.RunInteractorF(rc, fifo1, fifo2, g, tc)
			done <- 1
		}()

		res, err := c.RunUserF(rc, fifo1, fifo2, g, tc)
		<-done

		if err != nil {
			return language.Status{}, err
		}

		tc.OutputPath = filepath.Join(interactorSandbox.Pwd(), "out")
		conts, err := ioutil.ReadFile(tc.OutputPath)
		if err != nil {
			return language.Status{}, err
		}

		rc.Stdout.Write(conts)

		return res, multierr.Combine(errInteractor, err)
	}

	c.Batch.CleanupF = func(rc *batch.RunContext) error {
		return rc.Store["interactorSandbox"].(language.Sandbox).Cleanup()
	}

	return c
}

func (b Communication) Name() string {
	return "communication"
}

func init() {
	problems.RegisterTaskType(New())
}
