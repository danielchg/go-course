package main

import (
	"fmt"
	"github.com/go-playground/form/v4"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusBadRequest)
			return
		}

		err := r.ParseForm()

		if err != nil {
			http.Error(w, "Invalid Form", http.StatusBadRequest)
			return
		}

		var user struct {
			Name     string   `form:"name"`
			Surname  string   `form:"surname"`
			Username string   `form:"username"`
			Password string   `form:"password"`
			Age      int      `form:"age"`
			Groups   []string `form:"groups"`
		}

		decoder := form.NewDecoder()

		err = decoder.Decode(&user, r.PostForm)

		if err != nil {
			http.Error(w, "Invalid Form", http.StatusBadRequest)
			return
		}

		//user.Name = r.PostFormValue("name")
		//user.Surname = r.PostFormValue("surname")
		//user.Username = r.PostFormValue("username")
		//user.Password = r.PostFormValue("password")
		//user.Age, err = strconv.Atoi(r.PostFormValue("age"))
		//
		//if err != nil {
		//	log.Println(err)
		//	http.Error(w, "Invalid age", http.StatusBadRequest)
		//	return
		//}
		//
		//user.Groups = r.Form["groups"]

		log.Printf("%+v", user)
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	fmt.Println("Listening on port 8000...")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
