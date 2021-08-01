package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}
func getuserInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "User Id:89")
}

//NewHandler make a new myapp Handler
func NewHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", userHandler)
	mux.HandleFunc("/users/89", getuserInfoHandler)

	return mux
}
