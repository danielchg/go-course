package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sql02/storage"
)

func main() {
	db, err := sql.Open("mysql", "root:my-secret-pw@/test")
	if err != nil {
		log.Fatalf("Connection to DB fail: %v", err)
	}
	s := storage.NewMySQL(db)

	mux := http.NewServeMux()
	const userPrefix = "/user/"
	mux.Handle(userPrefix, http.StripPrefix(userPrefix, router(s)))

	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Cannot create server: ", err)
	}
}

func router(s storage.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var fn func(s storage.Storage, r *http.Request) (interface{}, int)

		switch r.Method {
		case http.MethodGet:
			fn = getUser
		case http.MethodPost:
			fn = createUser
		case http.MethodPut:
			fn = updateUser
		case http.MethodDelete:
			fn = deleteUser
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		result, code := fn(s, r)
		if code != http.StatusOK {
			w.WriteHeader(code)
		}

		e := json.NewEncoder(w)
		e.Encode(result)
	})

}

//Error base to be returned
type Error struct {
	Message string `json:"message"`
}

func createUser(s storage.Storage, r *http.Request) (interface{}, int) {
	d := json.NewDecoder(r.Body)
	u := storage.User{}

	if err := d.Decode(&u); err != nil {
		return Error{"Invalid User"}, http.StatusBadRequest
	}

	switch err := s.CreateUser(&u); err {
	case nil:
		return u, http.StatusCreated
	case storage.ErrDuplicate:
		return Error{"User already exist"}, http.StatusConflict
	default:
		return Error{"Internal Server Error"}, http.StatusInternalServerError
	}
}

func getUser(s storage.Storage, r *http.Request) (interface{}, int) {
	return nil, http.StatusOK
}

func updateUser(s storage.Storage, r *http.Request) (interface{}, int) {
	return nil, http.StatusOK
}

func deleteUser(s storage.Storage, r *http.Request) (interface{}, int) {
	return nil, http.StatusOK
}
