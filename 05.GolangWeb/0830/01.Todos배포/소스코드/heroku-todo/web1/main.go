package main

import (
	"log"
	"net/http"

	"web/web1/app"
)

func main() {
	port := os.Getenv("PORT")
	m := app.MakeHandler("./test.db")
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":"+port, m)
	if err != nil {
		panic(err)
	}
}
