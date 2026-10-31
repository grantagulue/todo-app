package main

import (
	"log"
	"net/http"
	"github.com/grant1x/todo/internals/database"
	"github.com/grant1x/todo/internals/handlers"
)

func main() {
	db := database.New()
	defer db.Pool.Close() 

	h := handlers.New(db)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /getnotes", h.HandleFetchNotes)
	mux.HandleFunc("POST /createnotes", h.HandleCreateNotes)
	mux.HandleFunc("GET /", h.HandleDefault)

	log.Fatal(http.ListenAndServe(":8080", mux))
}