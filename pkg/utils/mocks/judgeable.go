package mocks

import (
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/problems"
)


type Judgeable struct {
	FMemoryLimit func() int
	FTimeLimit func() int
	FChecker func() problems.Checker
	FInputOutputFiles func() (string, string)
	FLanguages func() []language.Language
	FStatusSkeleton func(testset string) (*problems.Status, error)
	FFiles func() []problems.File
	FGetTaskType func() problems.TaskType
}

func (mj *Judgeable) MemoryLimit() int {
	if mj.FMemoryLimit == nil {
		panic("not implemented")
	}

	return mj.FMemoryLimit()
}

func (mj *Judgeable) TimeLimit() int {
	if mj.FTimeLimit == nil {
		panic("not implemented") 
	}

	return mj.FTimeLimit()
}

func (mj *Judgeable) Checker() problems.Checker {
	if mj.FChecker == nil {
		panic("not implemented")
	}

	return mj.Checker()
}

func (mj *Judgeable) InputOutputFiles() (string, string) {
	if mj.FInputOutputFiles == nil {
		panic("not implemented")
	}

	return mj.FInputOutputFiles()
}

func (mj *Judgeable) Languages() []language.Language {
	if mj.FLanguages == nil {
		panic("not implemented")
	}

	return mj.FLanguages()
}

func (mj *Judgeable) StatusSkeleton(testset string) (*problems.Status, error) {
	if mj.FStatusSkeleton == nil {
		panic("not implemented")
	}

	return mj.FStatusSkeleton(testset)
}

func (mj *Judgeable) Files() []problems.File {
	if mj.FFiles == nil {
		panic("not implemented")
	}

	return mj.FFiles()
}

func (mj *Judgeable) GetTaskType() problems.TaskType {
	if mj.FGetTaskType == nil {
		panic("not implemented")
	}

	return mj.FGetTaskType()
}

