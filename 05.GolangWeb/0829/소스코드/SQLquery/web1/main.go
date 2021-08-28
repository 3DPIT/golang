package main

import (
	"log"
	"net/http"

	"web/web1/app"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":3003", m)
	if err != nil {
		panic(err)
	}
}
