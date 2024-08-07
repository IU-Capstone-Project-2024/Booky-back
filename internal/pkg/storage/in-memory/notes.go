package inmemory

import (
	"booky-back/internal/pkg/models"
	"booky-back/internal/pkg/storage"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NoteStorage struct {
	notes map[string]*models.Note
}

func NewNoteStorage() *NoteStorage {
	return &NoteStorage{
		notes: make(map[string]*models.Note),
	}
}

func (s *NoteStorage) CreateNote(note *models.Note) (*models.Note, error) {
	note.ID = uuid.New().String()
	note.CreatedAt = timestamppb.Now()

	s.notes[note.ID] = note
	return note, nil
}

func (s *NoteStorage) GetNote(id string) (*models.Note, error) {
	note, ok := s.notes[id]
	if !ok {
		return nil, fmt.Errorf("note with id %s was not found: %w", id, storage.ErrNotFound)
	}

	return note, nil
}

func (s *NoteStorage) UpdateNote(note *models.Note) (*models.Note, error) {
	_, ok := s.notes[note.ID]
	if !ok {
		return nil, fmt.Errorf("note with id %s was not found: %w", note.ID, storage.ErrNotFound)
	}

	s.notes[note.ID] = note
	s.notes[note.ID].UpdatedAt = timestamppb.Now()

	return note, nil
}

func (s *NoteStorage) DeleteNote(id string) error {
	_, ok := s.notes[id]
	if !ok {
		return fmt.Errorf("note with id %s was not found: %w", id, storage.ErrNotFound)
	}

	delete(s.notes, id)
	return nil
}

func (s *NoteStorage) ListNotes(courseID string) ([]*models.Note, error) {
	noteList := make([]*models.Note, 0, len(s.notes))
	for _, n := range s.notes {
		if n.CourseID == courseID {
			noteList = append(noteList, n)
		}
	}
	return noteList, nil
}
