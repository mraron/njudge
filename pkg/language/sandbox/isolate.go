package sandbox

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/mraron/njudge/pkg/language"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var IsolateRoot = getEnv("ISOLATE_ROOT", "/var/local/lib/isolate/")

type Isolate struct {
	id   int
	argv []string

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
	wdir   string

	st language.Status

	logger *log.Logger
}

func NewIsolate(id int) language.Sandbox {
	return &Isolate{id: id}
}

func (s *Isolate) Id() string {
	return "isolate" + strconv.Itoa(s.id)
}

func (s *Isolate) Pwd() string {
	return filepath.Join(IsolateRoot, strconv.Itoa(s.id), "box")
}

func (s *Isolate) GetFile(name string) (io.Reader, error) {
	f, err := os.ReadFile(filepath.Join(s.Pwd(), name))
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(f), nil
}

func (s *Isolate) ClearArguments() {
	s.argv = make([]string, 0)
	s.stdin = nil
	s.stdout = nil
	s.wdir = ""
	s.st = language.Status{}
}

func (s *Isolate) init() error {
	if err := s.Cleanup(); err != nil { //cleanup because the previous invocation might not have cleaned up
		return err
	}

	args := []string{"--cg", "-b", strconv.Itoa(s.id), "--init"}
	s.logger.Print("Running init: isolate with args ", args)

	err := exec.Command("isolate", args...).Run()
	return err
}

func (s *Isolate) Init(l *log.Logger) error {
	s.ClearArguments()
	s.logger = l

	return s.init()
}

func (s *Isolate) getPathToFile(name string) string {
	return filepath.Join(s.Pwd(), name)
}

func (s *Isolate) CreateFile(name string, r io.Reader) error {
	filename := s.getPathToFile(name)
	s.logger.Print("Creating file ", filename)

	if err := syscall.Unlink(filename); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		s.logger.Print("Error occurred while creating file ", err)
		return err
	}

	if _, err := io.Copy(f, r); err != nil {
		s.logger.Print("Error occurred while populating it with its content: ", err)
		return multierr.Combine(err, f.Close())
	}

	return f.Close()
}

func (s *Isolate) MakeExecutable(name string) error {
	filename := s.getPathToFile(name)

	err := os.Chmod(filename, 0755)
	s.logger.Print("Making executable: ", filename, " error: ", err)

	return err
}

func (s *Isolate) SetMaxProcesses(num int) language.Sandbox {
	if num < 0 {
		s.argv = append(s.argv, "--processes=100")
	} else {
		s.argv = append(s.argv, "--processes="+strconv.Itoa(num))
	}

	return s
}

func (s *Isolate) Env() language.Sandbox {
	s.argv = append(s.argv, "--full-env")
	return s
}

func (s *Isolate) SetEnv(e string) language.Sandbox {
	s.argv = append(s.argv, fmt.Sprintf("--env=%s", e))
	return s
}

func (s *Isolate) AddArg(a string) language.Sandbox {
	s.argv = append(s.argv, a)
	return s
}

func (s *Isolate) TimeLimit(tl time.Duration) language.Sandbox {
	tl = tl / time.Millisecond
	s.argv = append(s.argv, fmt.Sprintf("--time=%d.%d", tl/1000, tl%1000))
	s.argv = append(s.argv, fmt.Sprintf("--wall-time=%d.%d", (2*tl+1000)/1000, (2*tl+1000)%1000))

	return s
}

func (s *Isolate) MemoryLimit(ml int) language.Sandbox {
	s.argv = append(s.argv, "--cg-mem="+strconv.Itoa(ml))
	return s
}

func (s *Isolate) Verbose() language.Sandbox {
	s.argv = append(s.argv, "-v")
	return s
}

func (s *Isolate) Stdin(reader io.Reader) language.Sandbox {
	s.stdin = reader
	return s
}

func (s *Isolate) Stdout(writer io.Writer) language.Sandbox {
	s.stdout = writer
	return s
}

func (s *Isolate) Stderr(writer io.Writer) language.Sandbox {
	s.stderr = writer
	return s
}

func (s *Isolate) MapDir(src string, dest string, opts []string, checkExists bool) language.Sandbox {
	if checkExists {
		if _, err := os.Stat(src); os.IsNotExist(err) {
			return s
		}
	}

	format := fmt.Sprintf("--dir=%s=%s", src, dest)
	for _, opt := range opts {
		format += ":" + opt
	}
	s.argv = append(s.argv, format)

	return s
}

func (s *Isolate) WorkingDirectory(wd string) language.Sandbox {
	s.wdir = wd
	return s
}

func (s *Isolate) Run(args string, needStatus bool) (language.Status, error) {
	var (
		cmd *exec.Cmd
		err error

		metafile     *os.File
		metafileName = "/tmp/metafile" + strconv.Itoa(s.id)
	)

	s.MapDir("/etc/alternatives", "/etc/alternatives", []string{}, true)
	s.MapDir("/languages", "/languages", []string{"maybe"}, true)
	splitted := strings.Split(args, " ")
	s.argv = append([]string{"--cg", "--cg-timing", "-b", strconv.Itoa(s.id), "-M", metafileName}, s.argv...)
	s.argv = append(s.argv, "--run", "--")
	s.argv = append(s.argv, splitted...)
	defer s.ClearArguments()

	s.logger.Print("Running isolate with args ", s.argv)

	pb := NewPrefixBuffer(s.stderr, 2048)

	cmd = exec.Command("isolate", s.argv...)
	cmd.Stdin = s.stdin
	cmd.Stdout = s.stdout
	cmd.Stderr = pb
	cmd.Dir = s.wdir

	if err = cmd.Run(); err != nil {
		s.logger.Print("Command exited with non-zero exit code: ", err)

		if !needStatus {
			return s.st, err
		}

		str, st := string(pb.Prefix()), -1
		if strings.Contains(str, "Caught fatal signal") {
			fmt.Sscanf(str, "Caught fatal signal %d", &st)
		} else if strings.Contains(str, "Exited with error status") {
			fmt.Sscanf(str[strings.Index(str, "Exited with error status"):], "Exited with error status %d", &st)
		} else {
			s.logger.Print("unknown error status format: ", str)
		}

		if st == -1 {
			s.st.Verdict = language.VerdictXX
		} else { //eg. signal 8/136?? -> division by zero
			s.st.Verdict = language.VerdictRE
		}
	} else {
		s.logger.Print("Command exited successfully")
		s.logger.Printf("stderr of process: %q", string(pb.Prefix()))

		s.st.Verdict = language.VerdictOK
	}

	if !needStatus {
		return s.st, nil
	}

	memorySum := 0

	if metafile, err = os.Open(metafileName); err != nil {
		s.st.Verdict = language.VerdictXX
		s.logger.Print("Can't open metafile ", metafileName, " error is: ", err)

		return s.st, err
	}
	defer metafile.Close()

	s.logger.Print("Parsing metafile")

	sc := bufio.NewScanner(metafile)
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
				s.st.Verdict = language.VerdictTL
			case "SG":
				s.st.Verdict = language.VerdictRE
			}
		}
	}

	if err = sc.Err(); err != nil {
		s.st.Verdict = language.VerdictXX
		s.logger.Print("Error scanning metafile", err)

		return s.st, err
	}

	s.logger.Print("Calculated memory usage ", memorySum, "KiB")
	s.st.Memory = memorySum

	return s.st, nil
}

func (s *Isolate) Cleanup() error {
	args := []string{"--cg", "-b", strconv.Itoa(s.id), "--cleanup"}

	s.logger.Print("Executing cleanup with args ", args)

	return exec.Command("isolate", args...).Run()
}
