package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandler() http.Handler {
	r := mux.NewRouter()
	return r
}
