package main

import (
	"log"
	"net/http"
	"time"
	"web/web1/decoHandler"
	"web/web1/myapp"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time:", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed time:", time.Since(start).Milliseconds())
}

// func NewHandler() http.Handler {
// 	mux := myapp.NewHandler()
// 	h := decoHandler.NewDecoHandler(mux, logger)
// 	return h
// }
func NewHandler() http.Handler {
	h := myapp.NewHandler()
	h = decoHandler.NewDecoHandler(h, logger)
	h = decoHandler.NewDecoHandler(h, logger2)
	return h
}
func main() {

	mux := NewHandler()
	http.ListenAndServe(":3000", mux)
}
