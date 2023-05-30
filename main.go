package main

import (
	"fmt"
	"go_http_server/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func contentTypeMW(next *mux.Router) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

func main() {
	fmt.Println("Starting server on port 9090")

	router := mux.NewRouter()
	usersRouter := router.PathPrefix("/user").Subrouter()

	usersRouter.HandleFunc("", user.GetAllUsersRoute).Methods("GET")
	usersRouter.HandleFunc("", user.AddUserRoute).Methods("POST")
	
	usersRouter.HandleFunc("/{id}", user.GetOneUserRoute).Methods("GET")
	usersRouter.HandleFunc("/{id}", user.UpdateUserRoute).Methods("PUT")
	usersRouter.HandleFunc("/{id}", user.DeleteUserRoute).Methods("DELETE")

	usersRouter.StrictSlash(true)

	http.HandleFunc("/", contentTypeMW(router))

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("Unable to start server", err)
	}

}
