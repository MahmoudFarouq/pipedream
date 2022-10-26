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
				RemoteRequest{
					URL:    "https://webhook.site/bbdb4afb-4a75-4b45-9270-f75b2ef93454",
					Method: "POST",
					Body:   "mahmoud",
					Query:  "",
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
