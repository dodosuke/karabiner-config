package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT must be set")
		port = "8080"
	}

	http.HandleFunc("/json", json)

	http.ListenAndServe(":"+port, nil)
}

func json(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open("config.json")
	if e != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
