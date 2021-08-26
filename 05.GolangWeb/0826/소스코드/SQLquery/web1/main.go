package main

import (
	"log"
	"net/http"
	"web/web1/app"

	"github.com/urfave/negroni"
)

func main() {

	m := app.MakeHandler()
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":1005", n)
	if err != nil {
		panic(err)
	}
}
