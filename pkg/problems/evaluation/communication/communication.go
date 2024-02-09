package communication

import (
	"github.com/mraron/njudge/pkg/language/sandbox"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"github.com/mraron/njudge/pkg/problems"
	"github.com/mraron/njudge/pkg/problems/evaluation/batch"
	"github.com/mraron/njudge/pkg/problems/evaluation/stub"
	"go.uber.org/multierr"
)

type Communication struct {
	stub.Stub

	RunInteractorF func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (sandbox.Status, error)
	RunUserF       func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (sandbox.Status, error)
}

func New() *Communication {
	c := &Communication{
		stub.New(),
		func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (sandbox.Status, error) {
			//return rc.Store["interactorSandbox"].(sandbox.Sandbox).Stdin(utoi).Stdout(itou).TimeLimit(2*tc.TimeLimit).MemoryLimit(1024*1024).Run("interactor inp out", true)
			return sandbox.Status{}, nil
		},
		func(rc *batch.RunContext, utoi, itou *os.File, g *problems.Group, tc *problems.Testcase) (sandbox.Status, error) {
			//return rc.Lang.Run(rc.Sandbox, bytes.NewReader(rc.Binary), itou, utoi, tc.TimeLimit, tc.MemoryLimit)
			return sandbox.Status{}, nil
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

		sandbox.CreateFileFromSource(interactorSandbox, "interactor", f)
		interactorSandbox.MakeExecutable("interactor")

		rc.Store["interactorSandbox"] = interactorSandbox
		return nil
	}

	c.Batch.RunF = func(rc *batch.RunContext, g *problems.Group, tc *problems.Testcase) (sandbox.Status, error) {
		interactorSandbox := rc.Store["interactorSandbox"].(sandbox.Sandbox)
		testLocation, answerLocation := tc.InputPath, tc.AnswerPath

		testFile, err := os.Open(testLocation)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return sandbox.Status{}, err
		}
		defer testFile.Close()

		err = sandbox.CreateFileFromSource(interactorSandbox, "inp", testFile)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return sandbox.Status{}, err
		}

		answerFile, err := os.Open(answerLocation)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return sandbox.Status{}, err
		}
		defer answerFile.Close()

		dir, err := os.MkdirTemp("/tmp", "commtask")
		if err != nil {
			return sandbox.Status{}, err
		}
		err = os.Chmod(dir, 0755)
		if err != nil {
			return sandbox.Status{}, err
		}

		dir = filepath.Base(dir)

		rc.Store["tempdir"] = dir

		err = syscall.Mkfifo(filepath.Join("/tmp", dir, "fifo1"), 0666)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return sandbox.Status{}, err
		}

		err = os.Chmod(filepath.Join("/tmp", dir, "fifo1"), 0666)
		if err != nil {
			return sandbox.Status{}, err
		}

		err = syscall.Mkfifo(filepath.Join("/tmp", dir, "fifo2"), 0666)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return sandbox.Status{}, err
		}

		err = os.Chmod(filepath.Join("/tmp", dir, "fifo2"), 0666)
		if err != nil {
			return sandbox.Status{}, err
		}

		fifo1, err := os.OpenFile(filepath.Join("/tmp", dir, "fifo1"), os.O_RDWR, 0666)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return sandbox.Status{}, err
		}
		defer fifo1.Close()

		fifo2, err := os.OpenFile(filepath.Join("/tmp", dir, "fifo2"), os.O_RDWR, 0666)
		if err != nil {
			tc.VerdictName = problems.VerdictXX
			return sandbox.Status{}, err
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
		tc.MemoryUsed = res.Memory
		tc.TimeSpent = res.Time

		//err = multierr.Combine(err, os.Remove(filepath.Join("/tmp", dir, "fifo1")), os.Remove(filepath.Join("/tmp", dir, "fifo2")))

		if err != nil {
			return sandbox.Status{}, err
		}

		tc.OutputPath = filepath.Join(interactorSandbox.Pwd(), "out")
		conts, err := ioutil.ReadFile(tc.OutputPath)
		if err != nil {
			return sandbox.Status{}, err
		}

		rc.Stdout.Write(conts)

		return res, multierr.Combine(errInteractor, err)
	}

	c.Batch.CleanupF = func(rc *batch.RunContext) error {
		return rc.Store["interactorSandbox"].(sandbox.Sandbox).Cleanup(context.TODO())
	}

	return c
}

func (b Communication) Name() string {
	return "communication"
}
