package sandbox

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

	"github.com/mraron/njudge/pkg/language"
)

type Dummy struct {
	logger *log.Logger
	tmpdir string
	env    []string
	tl     time.Duration

	stdin          io.Reader
	stdout, stderr io.Writer

	workingDir string
}

func NewDummy() *Dummy {
	return &Dummy{}
}

func (s Dummy) Id() string {
	return s.tmpdir
}

func (s *Dummy) Init(logger *log.Logger) error {
	var err error
	if s.tmpdir, err = os.MkdirTemp("", "dummysandbox"); err != nil {
		return err
	}

	s.workingDir = s.tmpdir
	s.logger = logger
	return nil
}

func (s Dummy) Pwd() string {
	return s.tmpdir
}

func (s *Dummy) CreateFile(name string, r io.Reader) error {
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

func (s Dummy) GetFile(name string) (io.Reader, error) {
	f, err := ioutil.ReadFile(filepath.Join(s.Pwd(), name))
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(f), nil
}

func (s Dummy) MakeExecutable(name string) error {
	filename := filepath.Join(s.Pwd(), name)

	err := os.Chmod(filename, 0777)
	s.logger.Print("Making executable: ", filename, " error: ", err)

	return err
}

func (s *Dummy) SetMaxProcesses(i int) language.Sandbox {
	return s
}

func (s *Dummy) Env() language.Sandbox {
	s.env = os.Environ()
	return s
}

func (s *Dummy) SetEnv(env string) language.Sandbox {
	s.env = append(s.env, env+"="+os.Getenv(env))
	return s
}

func (s *Dummy) AddArg(string) language.Sandbox {
	return s
}

func (s *Dummy) TimeLimit(tl time.Duration) language.Sandbox {
	s.tl = tl
	return s
}

func (s *Dummy) MemoryLimit(int) language.Sandbox {
	return s
}

func (s *Dummy) Stdin(reader io.Reader) language.Sandbox {
	s.stdin = reader
	return s
}

func (s *Dummy) Stderr(writer io.Writer) language.Sandbox {
	s.stderr = writer
	return s
}

func (s *Dummy) Stdout(writer io.Writer) language.Sandbox {
	s.stdout = writer
	return s
}

func (s *Dummy) MapDir(x string, y string, i []string, b bool) language.Sandbox {
	return s
}

func (s *Dummy) WorkingDirectory(dir string) language.Sandbox {
	s.workingDir = dir
	return s
}

func (s *Dummy) Verbose() language.Sandbox {
	return s
}

func (s *Dummy) Run(prg string, needStatus bool) (language.Status, error) {
	cmd := exec.Command("bash", "-c", prg)
	cmd.Stdin = s.stdin
	cmd.Stdout = s.stdout
	cmd.Stderr = s.stderr
	cmd.Dir = s.workingDir
	cmd.Env = append(s.env, "PATH="+os.Getenv("PATH")+":"+s.tmpdir)

	var (
		st               language.Status
		errKill, errWait error
		finish           = make(chan bool, 1)
		wg               sync.WaitGroup
	)

	st.Verdict = language.VerdictOK

	start := time.NewTimer(s.tl)
	if err := cmd.Start(); err != nil {
		st.Verdict = language.VerdictXX
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
		st.Verdict = language.VerdictTL
		if errKill = cmd.Process.Kill(); errKill != nil {
			st.Verdict = language.VerdictXX
		}
	case <-finish:
	}

	wg.Wait()

	if errWait != nil && (strings.HasPrefix(errWait.Error(), "exit status") || strings.HasPrefix(errWait.Error(), "signal:")) {
		if st.Verdict == language.VerdictOK {
			st.Verdict = language.VerdictRE
		}
		errWait = nil
	}

	if errWait != nil {
		return st, errWait
	}

	return st, errKill
}

func (s *Dummy) Cleanup() error {
	return os.RemoveAll(s.tmpdir)
}
