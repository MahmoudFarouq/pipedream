package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Output map[string]interface{}

type WebhookTrigger struct {
	Request *http.Request
}

func (s WebhookTrigger) Execute() (Output, error) {
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
