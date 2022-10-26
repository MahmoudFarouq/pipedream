package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type WebhookTrigger struct {
	Request *http.Request
}

func (s WebhookTrigger) Name() string {
	return "trigger"
}

func (s WebhookTrigger) Execute(context.Context, Environment) (Output, error) {
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

	return Output{
		"method": s.Request.Method,
		"query":  s.Request.URL.Query(),
		"body":   b,
	}, nil
}
