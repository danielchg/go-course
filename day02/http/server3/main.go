package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Resquest description
type Resquest struct {
	Input string `json:"sample" xml:"test" mypkg:"tttt"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		type decodeFunc interface {
			Decode(interface{}) error
		}
		type encodeFunc interface {
			Encode(interface{}) error
		}

		var (
			decodeFn func(io.Reader) decodeFunc
			encodeFn func(io.Writer) encodeFunc
		)

		switch ct := req.Header.Get("Content-Type"); ct {
		case "application/json":
			decodeFn = func(r io.Reader) decodeFunc{
				return json.NewDecoder(r)
			}
		case "application/xml":
			decodeFn = func (r io.Reader) decodeFunc {
				return xml.NewDecoder(r)
			}
		default:
			http.Error(w, fmt.Sprintf("Invalid format %q", ct), http.StatusBadRequest)
			return
		}
		ac := req.Header.Get("Accept")
		switch ac {
		case "application/json":
			encodeFn = func(w io.Writer) encodeFunc{
				return json.NewEncoder(w)
			}
		case "application/xml":
			encodeFn = func(w io.Writer) encodeFunc {
				return xml.NewEncoder(w)
			}
		default:
			http.Error(w, fmt.Sprintf("Invalid format %q", ac), http.StatusBadRequest)
			return
		}

		v := Resquest{}

		d := decodeFn(req.Body)
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

		a := encodeFn(w)
		//		a.SetIndent("", "    ")
		w.Header().Set("Content-Type", ac)
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
