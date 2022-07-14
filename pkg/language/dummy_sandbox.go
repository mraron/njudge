package language

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type DummySandbox struct {
	logger *log.Logger
	tmpdir string
	env    []string
	tl     time.Duration

	stdin          io.Reader
	stdout, stderr io.Writer

	workingDir string
}

func NewDummySandbox() *DummySandbox {
	return &DummySandbox{}
}

func (s DummySandbox) Id() string {
	return s.tmpdir
}

func (s *DummySandbox) Init(logger *log.Logger) error {
	var err error
	if s.tmpdir, err = ioutil.TempDir("", "dummysandbox"); err != nil {
		return err
	}

	s.workingDir = s.tmpdir
	s.logger = logger
	return nil
}

func (s DummySandbox) Pwd() string {
	return s.tmpdir
}

func (s *DummySandbox) CreateFile(name string, r io.Reader) error {
	filename := filepath.Join(s.tmpdir, name)
	s.logger.Print("Creating file ", filename)

	f, err := os.Create(filename)
	if err != nil {
		s.logger.Print("Error occurred while creating file ", err)
		return err
	}

	if _, err := io.Copy(f, r); err != nil {
		s.logger.Print("Error occurred while populating it with its content: ", err)
		f.Close()
		return err
	}

	return f.Close()
}

func (s DummySandbox) GetFile(name string) (io.Reader, error) {
	f, err := ioutil.ReadFile(filepath.Join(s.Pwd(), name))
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(f), nil
}

func (s DummySandbox) MakeExecutable(name string) error {
	filename := filepath.Join(s.Pwd(), name)

	err := os.Chmod(filename, 0777)
	s.logger.Print("Making executable: ", filename, " error: ", err)

	return err
}

func (s *DummySandbox) SetMaxProcesses(i int) Sandbox {
	return s
}

func (s *DummySandbox) Env() Sandbox {
	s.env = os.Environ()
	return s
}

func (s *DummySandbox) SetEnv(env string) Sandbox {
	s.env = append(s.env, env+"="+os.Getenv(env))
	return s
}

func (s *DummySandbox) AddArg(string) Sandbox {
	return s
}

func (s *DummySandbox) TimeLimit(tl time.Duration) Sandbox {
	s.tl = tl
	return s
}

func (s *DummySandbox) MemoryLimit(int) Sandbox {
	return s
}

func (s *DummySandbox) Stdin(reader io.Reader) Sandbox {
	s.stdin = reader
	return s
}

func (s *DummySandbox) Stderr(writer io.Writer) Sandbox {
	s.stderr = writer
	return s
}

func (s *DummySandbox) Stdout(writer io.Writer) Sandbox {
	s.stdout = writer
	return s
}

func (s *DummySandbox) MapDir(x string, y string, i []string, b bool) Sandbox {
	return s
}

func (s *DummySandbox) WorkingDirectory(dir string) Sandbox {
	s.workingDir = dir
	return s
}

func (s *DummySandbox) Verbose() Sandbox {
	return s
}

func (s *DummySandbox) Run(prg string, needStatus bool) (Status, error) {
	cmd := exec.Command("bash", "-c", prg)
	cmd.Stdin = s.stdin
	cmd.Stdout = s.stdout
	cmd.Stderr = s.stderr
	cmd.Dir = s.workingDir
	cmd.Env = append(s.env, "PATH="+os.Getenv("PATH")+":"+s.tmpdir)

	var (
		st               Status
		errKill, errWait error
		finish           = make(chan bool, 1)
		wg               sync.WaitGroup
	)

	st.Verdict = VERDICT_OK

	start := time.NewTimer(s.tl)
	if err := cmd.Start(); err != nil {
		st.Verdict = VERDICT_XX
		return st, err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		errWait = cmd.Wait()
		finish <- true
	}()

	select {
	case <-start.C:
		st.Verdict = VERDICT_TL
		if errKill = cmd.Process.Kill(); errKill != nil {
			st.Verdict = VERDICT_XX
		}
	case <-finish:
	}

	wg.Wait()

	if errWait != nil && (strings.HasPrefix(errWait.Error(), "exit status") || strings.HasPrefix(errWait.Error(), "signal:")) {
		if st.Verdict == VERDICT_OK {
			st.Verdict = VERDICT_RE
		}
		errWait = nil
	}

	if errWait != nil {
		return st, errWait
	}

	return st, errKill
}

func (s *DummySandbox) Cleanup() error {
	return os.RemoveAll(s.tmpdir)
}
