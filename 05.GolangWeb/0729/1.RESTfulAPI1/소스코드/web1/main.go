package main

import (
	"net/http"
	"web/web1/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
