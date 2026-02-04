package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	ID          string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	UserID      string         `gorm:"type:varchar(36);not null;index" json:"user_id"`
	Title       string         `gorm:"type:varchar(200);not null" json:"title"`
	Content     string         `gorm:"type:text" json:"content"`
	Category    string         `gorm:"type:varchar(50);index" json:"category"`
	Tags        string         `gorm:"type:varchar(500)" json:"tags"`
	ViewCount   int            `gorm:"default:0" json:"view_count"`
	LikeCount   int            `gorm:"default:0" json:"like_count"`
	AnswerCount int            `gorm:"default:0" json:"answer_count"`
	Status      int            `gorm:"default:1" json:"status"` // 1:正常 2:隐藏 3:删除
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	User        User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (q *Question) BeforeCreate(tx *gorm.DB) error {
	if q.ID == "" {
		q.ID = uuid.New().String()
	}
	return nil
}
