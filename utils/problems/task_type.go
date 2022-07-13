package problems

import (
	"fmt"
	"io"

	"github.com/mraron/njudge/utils/language"
)

type TaskType interface {
	Name() string
	Compile(Judgeable, language.Sandbox, language.Language, io.Reader, io.Writer) (io.Reader, error)
	Run(Judgeable, *language.SandboxProvider, language.Language, io.Reader, chan string, chan Status) (Status, error)
}

var taskTypes []TaskType

func init() {
	taskTypes = make([]TaskType, 0)
}

func RegisterTaskType(taskType TaskType) {
	taskTypes = append(taskTypes, taskType)
}

func GetTaskType(name string) (TaskType, error) {
	for _, tt := range taskTypes {
		if tt.Name() == name {
			return tt, nil
		}
	}

	return nil, fmt.Errorf("no such task type: %q", name)
}
