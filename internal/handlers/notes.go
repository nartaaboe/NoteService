package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yourusername/notes-microservice/internal/models"
	"github.com/yourusername/notes-microservice/internal/storage"
)

type NotesHandler struct {
	storage *storage.MemoryStorage
}

func NewNotesHandler(storage *storage.MemoryStorage) *NotesHandler {
	return &NotesHandler{storage: storage}
}


func (h *NotesHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	note.ID = fmt.Sprintf("%d", len(h.storage.GetAllNotes())+1)
	h.storage.AddNote(note)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func (h *NotesHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	allNotes := h.storage.GetAllNotes()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allNotes)
}
