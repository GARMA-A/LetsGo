package main

import (
	"log"
	"net/http"

	"github.com/GARMA-A/LetsGo/simpleServer/backend/cmd/web"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", web.Home)
	mux.HandleFunc("GET /snippet/view/{id}", web.SnippetView)
	mux.HandleFunc("GET /snippet/create", web.SnippetCreate)
	mux.HandleFunc("POST /snippet/create", web.SnippetCreatePost)
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
