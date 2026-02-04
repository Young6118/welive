package model

import (
	"time"

	"gorm.io/gorm"
)

// Note 笔记模型
type Note struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Title    string `gorm:"size:200;not null" json:"title"`
	Content  string `gorm:"type:text;not null" json:"content"`
	AuthorID uint   `gorm:"not null" json:"author_id"`
	Category string `gorm:"size:50" json:"category"`
	Tags     string `gorm:"size:500" json:"tags"` // JSON格式存储
	Status   int    `gorm:"default:1" json:"status"`

	Author User `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
}

// TableName 指定表名
func (Note) TableName() string {
	return "notes"
}

// NoteCategory 笔记分类模型
type NoteCategory struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name   string `gorm:"size:50;not null" json:"name"`
	UserID uint   `gorm:"not null" json:"user_id"`
	Sort   int    `gorm:"default:0" json:"sort"`
}

// TableName 指定表名
func (NoteCategory) TableName() string {
	return "note_categories"
}
