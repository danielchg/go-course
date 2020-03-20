package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		log.Fatalf("can not connect to URL: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("error calling to URL: %v", res.StatusCode)
	}

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal("error reading the request body")
	}
}
