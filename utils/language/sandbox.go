package language

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//@TODO reimplement this functionality with like logs and etc.

var ISOLATE_ROOT = getEnv("ISOLATE_ROOT", "/var/local/lib/isolate/")

type Sandbox interface {
	ClearArguments()
	Init(*log.Logger) error

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
	Verbose() Sandbox
	Run(string, bool) (Status, error)

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

	logger *log.Logger
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

func (s *IsolateSandbox) ClearArguments() {
	s.argv = make([]string, 0)
	s.stdin = nil
	s.stdout = nil
	s.wdir = ""
	s.st = Status{}
}

func (s *IsolateSandbox) Init(l *log.Logger) error {
	s.ClearArguments()
	s.logger = l

	args := []string{"--cg", "-b", strconv.Itoa(s.id), "--init"}

	s.logger.Print("Running init: isolate with args ", args)

	err := exec.Command("isolate", args...).Run()
	return err
}

func (s *IsolateSandbox) CreateFile(name string, r io.Reader) error {
	filename := ISOLATE_ROOT + strconv.Itoa(s.id) + "/box/" + name
	s.logger.Print("Creating file ", filename)

	f, err := os.Create(filename)
	if err != nil {
		s.logger.Print("Error occured while creating file ", err)
		return err
	}

	if _, err := io.Copy(f, r); err != nil {
		s.logger.Print("Error occured while populating it with its content: ", err)
		f.Close()
		return err
	}

	return f.Close()
}

func (s *IsolateSandbox) MakeExecutable(name string) error {
	filename := ISOLATE_ROOT + strconv.Itoa(s.id) + "/box/" + name

	err := os.Chmod(filename, 0777)
	s.logger.Print("Making executable: ", filename, " error: ", err)

	return err
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
	s.argv = append(s.argv, "--full-env")
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

func (s *IsolateSandbox) Run(prg string, needStatus bool) (Status, error) {
	var (
		err      error
		f        *os.File
		cmd      *exec.Cmd
		str      string
		st       int
		metafile = "/tmp/metafile" + strconv.Itoa(s.id)
	)

	defer s.ClearArguments()

	splt := strings.Split(prg, " ")

	s.argv = append([]string{"--cg", "--cg-timing", "-b", strconv.Itoa(s.id), "-M", metafile}, s.argv...)
	s.argv = append(s.argv, "--run", "--")
	s.argv = append(s.argv, splt...)

	stderr := &bytes.Buffer{}

	s.logger.Print("Running isolate with args ", s.argv)

	cmd = exec.Command("isolate", s.argv...)

	cmd.Stdin = s.stdin
	cmd.Stdout = s.stdout
	cmd.Stderr = stderr
	cmd.Dir = s.wdir

	if err = cmd.Run(); err != nil {
		s.logger.Print("Command exited with non-zero exit code: ", err)

		if !needStatus {
			s.stderr.Write(stderr.Bytes())
			return s.st, err
		}

		str, st = stderr.String(), -1
		if strings.Contains(str, "Caught fatal signal") {
			fmt.Sscanf(str, "Caught fatal signal %d", &st)
		}else if(strings.Contains(str, "Exited with error status")) {
			fmt.Sscanf(str[strings.Index(str, "Exited with error status"):], "Exited with error status %d", &st)
		}else {
			s.logger.Print("unknown error status format: ", str)
		}

		if st == -1 {
			s.st.Verdict = VERDICT_XX
		} else if st == 127 {
			s.st.Verdict = VERDICT_ML
		} else { //eg. signal 8/136?? -> division by zero
			s.st.Verdict = VERDICT_RE
		}
	} else {
		s.logger.Print("Command exited successfully")
		s.logger.Print("stderr of process: ", stderr.String())

		s.st.Verdict = VERDICT_OK
	}

	s.stderr.Write([]byte(str))

	if !needStatus {
		return s.st, nil
	}

	memorySum := 0

	if f, err = os.Open(metafile); err != nil {
		s.st.Verdict = VERDICT_XX
		s.logger.Print("Can't open metafile ", metafile, " error is: ", err)

		return s.st, err
	}

	s.logger.Print("Ok now, getting status")

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s.logger.Print(sc.Text())

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

	s.logger.Print("Calculated memory usage ", memorySum, "KiB")
	s.logger.Print("===============")
	s.st.Memory = memorySum

	return s.st, nil
}

func (s *IsolateSandbox) Cleanup() error {
	args := []string{"--cg", "-b", strconv.Itoa(s.id), "--cleanup"}

	s.logger.Print("Executing cleanup with args ", args)

	return exec.Command("isolate", args...).Run()
}
