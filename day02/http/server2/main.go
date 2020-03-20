package main

import (
	"encoding/xml"
	"log"
	"net/http"
)

// XMLResquest description
type XMLResquest struct {
	Input string `json:"sample" xml:"test" mypkg:"tttt"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/json", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		v := XMLResquest{}

		d := xml.NewDecoder(r.Body)
		//		d.DisallowUnknownFields()

		err := d.Decode(&v)
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s := []rune(v.Input)
		for i := 0; i < len(s)/2; i++ {
			s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
		}
		/**s := strings.Builder{}
		for i := len(v.Input) - 1; i >= 0; i-- {
			s.WriteByte(v.Input[i])
		}**/
		v.Input = string(s)

		a := xml.NewEncoder(w)
		//		a.SetIndent("", "    ")
		w.Header().Set("Content-Type", "text/xml")
		a.Encode(v)

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
