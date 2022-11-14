package models

import (
	"gorm.io/gorm"
	"time"
)

type NoteEntity struct {
	Id        string `json:"id,omitempty" gorm:"primaryKey,column:id"`
	Title     string `json:"title,omitempty" gorm:"column:title"`
	Content   string `json:"content,omitempty" gorm:"column:content"`
	OwnerId   string `json:"owner_id,omitempty" gorm:"column:owner_id"`
	CreatedAt int64  `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt int64  `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (r NoteEntity) TableName() string {
	return "note"
}

func (r NoteEntity) BeforeCreate(*gorm.DB) error {
	r.CreatedAt = time.Now().UnixMilli()
	return nil
}

type CreateNoteRequest struct {
	Title   string `json:"title,omitempty" validate:"required" form:"title"`
	Content string `json:"content,omitempty" validate:"required" form:"content"`
}

type NoteResponse struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type UpdateNotes struct {
	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
}
