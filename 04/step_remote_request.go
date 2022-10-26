package main

import (
	"context"
	"net/http"
	"strings"
)

type RemoteRequest struct {
	URL    string
	Method string
	Body   string
	Query  string
}

func (r RemoteRequest) Execute(ctx context.Context, environment Environment) (Output, error) {
	url, err := environment.Render(r.URL)
	if err != nil {
		return nil, err
	}

	method, err := environment.Render(r.Method)
	if err != nil {
		return nil, err
	}

	body, err := environment.Render(r.Body)
	if err != nil {
		return nil, err
	}

	query, err := environment.Render(r.Query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url+"?"+query, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r RemoteRequest) Name() string {
	return "remote-request"
}
