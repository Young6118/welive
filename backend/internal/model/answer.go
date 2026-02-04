package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Answer struct {
	ID         string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	QuestionID string         `gorm:"type:varchar(36);not null;index" json:"question_id"`
	UserID     string         `gorm:"type:varchar(36);not null;index" json:"user_id"`
	Content    string         `gorm:"type:text;not null" json:"content"`
	IsAI       bool           `gorm:"default:false" json:"is_ai"`
	LikeCount  int            `gorm:"default:0" json:"like_count"`
	Status     int            `gorm:"default:1" json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	User       User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (a *Answer) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return nil
}
