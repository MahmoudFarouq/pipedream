package main

import (
	"context"
)

type Output map[string]interface{}

type Step interface {
	Name() string
	Execute(context.Context, Environment) (Output, error)
}

type Workflow struct {
	steps []Step
}

func (s Workflow) Execute(ctx context.Context) error {
	env := NewEnvironment()

	for _, step := range s.steps {
		out, err := step.Execute(ctx, env)
		if err != nil {
			return err
		}

		env = env.Consume(step.Name(), out)
	}

	return nil
}
