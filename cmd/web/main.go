package main

import (
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./internal/views/index.html")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request){}

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/up", healthCheckHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
