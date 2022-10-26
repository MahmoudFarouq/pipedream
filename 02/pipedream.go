package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Output map[string]interface{}

type Step interface {
	Execute(context.Context) (Output, error)
}

type WebhookTrigger struct {
	Request *http.Request
}

func (s WebhookTrigger) Execute(context.Context) (Output, error) {
	body, err := io.ReadAll(s.Request.Body)
	if err != nil {
		return nil, err
	}

	var b any

	b = body

	var m map[string]interface{}

	err = json.Unmarshal(body, &m)
	if err == nil {
		b = m
	}

	return map[string]interface{}{
		"method": s.Request.Method,
		"query":  s.Request.URL.Query(),
		"body":   b,
	}, nil
}

type RemoteRequest struct {
	URL    string
	Method string
	Body   string
	Query  string
}

func (r RemoteRequest) Execute(ctx context.Context) (Output, error) {
	req, err := http.NewRequest(r.Method, r.URL, strings.NewReader(r.Body))
	if err != nil {
		return nil, err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

type Workflow struct {
	steps []Step
}

func (s Workflow) Execute(ctx context.Context) error {
	for _, step := range s.steps {
		_, err := step.Execute(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
