//mux로 라우터 만들어 적용
package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello foo")
}
func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello bar")
}
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})
	http.ListenAndServe(":8080", mux)
}
