package note_repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"go-flutter-bootcamp/helper"
	"go-flutter-bootcamp/models"
	"gorm.io/gorm"
	"log"
	"time"
)

func New(db *gorm.DB) Repository {
	err := db.AutoMigrate(&models.NoteEntity{})
	if err != nil {
		log.Fatal(err)
	}
	return SqlRepository{
		db: db,
	}
}

type SqlRepository struct {
	db *gorm.DB
}

func (s SqlRepository) Update(ctx context.Context, ownerId string, noteId string, req *models.UpdateNotes) error {
	query := "UPDATE note SET title = ?, content = ?, updated_at = ? WHERE owner_id = ? and id = ?"
	updatedAt := time.Now().UnixMilli()
	tx := s.db.WithContext(ctx).Exec(query, req.Title, req.Content, updatedAt, ownerId, noteId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s SqlRepository) Create(ctx context.Context, ownerId string, req *models.CreateNoteRequest) error {
	data, _ := helper.TypeConverter[models.NoteEntity](&req)
	data.OwnerId = ownerId
	data.Id = uuid.NewString()
	data.CreatedAt = time.Now().UnixMilli()
	data.UpdatedAt = time.Now().UnixMilli()
	tx := s.db.WithContext(ctx).Model(models.NoteEntity{}).Create(&data)
	return tx.Error
}

func (s SqlRepository) Delete(ctx context.Context, ownerId string, noteId string) error {
	tx := s.db.WithContext(ctx).Exec("DELETE FROM note WHERE id = ? AND owner_id = ?", noteId, ownerId)
	return tx.Error
}

func (s SqlRepository) GetAll(ctx context.Context, ownerId string, req *models.GetNoteRequest) (*[]models.NoteResponse, error) {
	notes := make([]models.NoteResponse, 0)
	offset := (req.Page - 1) * req.Limit
	query := "SELECT * FROM note WHERE owner_id = @ownerId"
	if req.Page > 0 {
		query = "SELECT * FROM note WHERE owner_id = @ownerId LIMIT @limit OFFSET @offset"
	}
	tx := s.db.WithContext(ctx).Raw(query, sql.Named("ownerId", ownerId), sql.Named("limit", req.Limit), sql.Named("offset", offset)).Scan(&notes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &notes, nil
}
