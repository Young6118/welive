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

	Title    string `gorm:"size:200;not null;index" json:"title"`
	Content  string `gorm:"type:text;not null" json:"content"`
	AuthorID uint   `gorm:"not null;index" json:"author_id"`
	Category string `gorm:"size:50;index" json:"category"`
	Tags     string `gorm:"size:500" json:"tags"` // JSON格式存储
	Status   int    `gorm:"default:1;index" json:"status"`

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
	UserID uint   `gorm:"not null;index" json:"user_id"`
	Sort   int    `gorm:"default:0" json:"sort"`
}

// NoteLike 笔记点赞模型
type NoteLike struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	NoteID uint `gorm:"not null;index:idx_note_user,unique" json:"note_id"`
	UserID uint `gorm:"not null;index:idx_note_user,unique" json:"user_id"`
}

// TableName 指定表名
func (NoteLike) TableName() string {
	return "note_likes"
}
