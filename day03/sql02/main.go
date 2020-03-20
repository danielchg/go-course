package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResquest description
type JSONResquest struct {
	Input string `json:"sample"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/json", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		v := JSONResquest{}

		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		err := d.Decode(&v)
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("%v", v)
		log.Printf("%+v", v)
		log.Printf("%#v", v)
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Cannot create server: ", err)
	}
}
