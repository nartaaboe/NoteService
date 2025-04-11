package storage

import (
	"github.com/yourusername/notes-microservice/internal/models"
	"sync"
)

type MemoryStorage struct {
	notes []models.Note
	mu    sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{notes: make([]models.Note, 0)}
}

func (s *MemoryStorage) AddNote(note models.Note) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.notes = append(s.notes, note)
}

func (s *MemoryStorage) GetAllNotes() []models.Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.notes
}
