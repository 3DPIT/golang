//json인것 알려주기
package main

import (
	"net/http"
	"web/web1/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
