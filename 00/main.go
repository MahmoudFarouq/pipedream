package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":1234", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			return
		}

		_, err = writer.Write(body)
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
