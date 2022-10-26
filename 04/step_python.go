package main

import (
	"context"
)

type StepPython struct {
	SourceCode string
}

func (s StepPython) Name() string {
	return "python"
}

func (s StepPython) Execute(ctx context.Context, environment Environment) (Output, error) {
	//TODO implement me
	panic("implement me")
}
