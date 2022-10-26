package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":1234", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		s := WebhookTrigger{Request: request}

		out, err := s.Execute()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		b, err := json.Marshal(out)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		_, err = writer.Write(b)
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
