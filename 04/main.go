package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":1234", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		wf := Workflow{
			steps: []Step{
				WebhookTrigger{Request: request},
				StepPython{},
				RemoteRequest{
					URL:    "{{python.remote_url}}",
					Method: "{{trigger.method}}",
					Body:   "{{trigger.body}}",
					Query:  "hello={{trigger.query.hello}}",
				},
			},
		}

		err := wf.Execute(request.Context())
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		return
	}))
	if err != nil {
		fmt.Printf("error serving: %v", err)
	}

	os.Exit(0)
}
