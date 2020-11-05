package problems

import (
	"errors"
)

var taskTypes []TaskType

func init() {
	taskTypes = make([]TaskType, 0)
}

func RegisterTaskType(taskType TaskType) {
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
