package problems

import (
	"errors"
	"github.com/labstack/gommon/log"
)

var taskTypes []TaskType

func init() {
	taskTypes = make([]TaskType, 0)
}

func RegisterTaskType(taskType TaskType) {
	log.Print("registered", taskType.Name())
	taskTypes = append(taskTypes, taskType)
}

func GetTaskType(name string) TaskType {
	for _, tt := range taskTypes {
		if tt.Name() == name {
			return tt
		}
	}

	panic(errors.New("no such task type: " + name))

	return nil
}
