package language

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var ISOLATE_ROOT = getEnv("ISOLATE_ROOT", "/var/local/lib/isolate/")

type Sandbox interface {
	Init() error

	Pwd() string

	CreateFile(string, io.Reader) error
	GetFile(string) (io.Reader, error)
	MakeExecutable(string) error

	SetMaxProcesses(int) Sandbox
	Env() Sandbox
	TimeLimit(time.Duration) Sandbox
	MemoryLimit(int) Sandbox
	Stdin(io.Reader) Sandbox
	Stderr(io.Writer) Sandbox
	Stdout(io.Writer) Sandbox
	WorkingDirectory(string) Sandbox
	Run(string, bool) error

	Status() Status

	Cleanup() error
}

type IsolateSandbox struct {
	id   int
	argv []string

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
	wdir   string

	st Status

	sync.Mutex
}

func NewIsolateSandbox(id int) Sandbox {
	return &IsolateSandbox{id: id}
}

func (s *IsolateSandbox) Pwd() string {
	return ISOLATE_ROOT + strconv.Itoa(s.id) + "/box/"
}

func (s *IsolateSandbox) GetFile(name string) (io.Reader, error) {
	f, err := ioutil.ReadFile(filepath.Join(s.Pwd(), name))
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(f), nil
}

func (s *IsolateSandbox) Init() error {
	s.Lock()

	s.argv = make([]string, 0)
	s.stdin = nil
	s.stdout = nil
	s.wdir = ""
	s.st = Status{}

	err := exec.Command("isolate", "--cg", "-b", strconv.Itoa(s.id), "--init").Run()

	return err
}

func (s *IsolateSandbox) CreateFile(name string, r io.Reader) error {
	f, err := os.Create(ISOLATE_ROOT + strconv.Itoa(s.id) + "/box/" + name)
	if err != nil {
		return err
	}

	if _, err := io.Copy(f, r); err != nil {
		f.Close()
		return err
	}

	return f.Close()
}

func (s *IsolateSandbox) MakeExecutable(name string) error {
	return os.Chmod(ISOLATE_ROOT+strconv.Itoa(s.id)+"/box/"+name, 0777)
}

func (s *IsolateSandbox) SetMaxProcesses(num int) Sandbox {
	if num < 0 {
		s.argv = append(s.argv, "--processes=100")
	} else {
		s.argv = append(s.argv, "--processes="+strconv.Itoa(num))
	}

	return s
}

func (s *IsolateSandbox) Env() Sandbox {
	s.argv = append(s.argv, "-e")
	return s
}

func (s *IsolateSandbox) TimeLimit(tl time.Duration) Sandbox {
	tl = tl / time.Millisecond
	s.argv = append(s.argv, fmt.Sprintf("--time=%d.%d", tl/1000, tl%1000))

	return s
}

func (s *IsolateSandbox) MemoryLimit(ml int) Sandbox {
	s.argv = append(s.argv, "--cg-mem="+strconv.Itoa(ml), "--mem="+strconv.Itoa(ml))
	return s
}

func (s *IsolateSandbox) Verbose() Sandbox {
	s.argv = append(s.argv, "-v")
	return s
}

func (s *IsolateSandbox) Stdin(reader io.Reader) Sandbox {
	s.stdin = reader
	return s
}

func (s *IsolateSandbox) Stdout(writer io.Writer) Sandbox {
	s.stdout = writer
	return s
}

func (s *IsolateSandbox) Stderr(writer io.Writer) Sandbox {
	s.stderr = writer
	return s
}

func (s *IsolateSandbox) WorkingDirectory(wd string) Sandbox {
	s.wdir = wd
	return s
}

func (s *IsolateSandbox) Run(prg string, needStatus bool) error {
	var (
		err error
		f   *os.File
		cmd *exec.Cmd
		str string
		st  int
	)

	splt := strings.Split(prg, " ")

	s.argv = append([]string{"--cg", "--cg-timing", "-b", strconv.Itoa(s.id), "-M", "/tmp/metafile" + strconv.Itoa(s.id)}, s.argv...)
	s.argv = append(s.argv, "--run", "--")
	s.argv = append(s.argv, splt...)

	fmt.Println(s.argv, "!!!!")

	stderr := &bytes.Buffer{}

	cmd = exec.Command("isolate", s.argv...)

	cmd.Stdin = s.stdin
	cmd.Stdout = s.stdout
	cmd.Stderr = stderr
	cmd.Dir = s.wdir

	if err = cmd.Run(); err != nil {
		if !needStatus {
			s.stderr.Write(stderr.Bytes())
			return err
		}

		str, st = stderr.String(), -1
		fmt.Sscanf(str, "Caught fatal signal %d", &st)

		if st == -1 {
			fmt.Println("caught shiet", stderr.String())
			s.st.Verdict = VERDICT_XX
		} else if st == 2 || st == 127 {
			//fmt.Println(st)
			s.st.Verdict = VERDICT_ML
		} else { //signal 8 -> division by zero
			s.st.Verdict = VERDICT_RE
		}
	} else {
		s.st.Verdict = VERDICT_OK
	}

	s.stderr.Write([]byte(str))

	if !needStatus {
		return nil
	}

	memorySum := 0

	if f, err = os.Open("/tmp/metafile" + strconv.Itoa(s.id)); err != nil {
		s.st.Verdict = VERDICT_XX
		fmt.Println("can't open shiet")
		return err
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lst := strings.Split(sc.Text(), ":")

		if lst[0] == "max-rss" || lst[0] == "cg-mem" {
			s.st.Memory, _ = strconv.Atoi(lst[1])
			memorySum += s.st.Memory
		} else if lst[0] == "time" {
			tmp, _ := strconv.ParseFloat(lst[1], 32)
			s.st.Time = time.Duration(tmp*1000) * time.Millisecond
		} else if lst[0] == "status" {
			switch lst[1] {
			case "TO":
				s.st.Verdict = VERDICT_TL
			case "SG":
				s.st.Verdict = VERDICT_RE
			}
		}
	}

	s.st.Memory = memorySum

	return nil
}

func (s *IsolateSandbox) Status() Status {
	return s.st
}

func (s *IsolateSandbox) Cleanup() error {
	err := exec.Command("isolate", "--cg", "-b", strconv.Itoa(s.id), "--cleanup").Run()
	s.Unlock()

	return err
}
