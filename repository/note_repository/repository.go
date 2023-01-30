package note_repository

import (
	"context"
	"go-flutter-bootcamp/models"
)

type Repository interface {
	Create(ctx context.Context, ownerId string, req *models.CreateNoteRequest) error
	Delete(ctx context.Context, ownerId string, noteId string) error
	GetAll(ctx context.Context, ownerId string, req *models.GetNoteRequest) (*[]models.NoteResponse, error)
	Update(ctx context.Context, ownerId string, noteId string, req *models.UpdateNotes) error
}
