package main

import (
	"log"
	"net/http"

	"github.com/yourusername/notes-microservice/internal/handlers"
	"github.com/yourusername/notes-microservice/internal/storage"
)



func main() {
	storage := storage.NewMemoryStorage()

	notesHandler := handlers.NewNotesHandler(storage)

	http.HandleFunc("/notes", notesHandler.GetNotes)
	http.HandleFunc("/notes/create", notesHandler.CreateNote)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
