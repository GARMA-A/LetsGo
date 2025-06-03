package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HomeversionOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server:", "GO")
	w.Write([]byte("Welcome to the Home Page!"))
}

func ItemViewversionOne(w http.ResponseWriter, r *http.Request) {
	// Extracting the item ID from the URL path
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Viewing Item %d", id)
	w.Write([]byte(msg))
}

func ItemCreateversionOne(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(201) // HTTP 201 Created
	w.Write([]byte("Creating Item!"))
}

func mainversionOne() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", HomeversionOne)
	mux.HandleFunc("GET /item/view/{id}", ItemViewversionOne)
	mux.HandleFunc("POST /item/create", ItemCreateversionOne)

	println("Starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
