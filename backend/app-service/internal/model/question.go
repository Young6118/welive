package model

import (
	"time"

	"gorm.io/gorm"
)

// Question 问题模型
type Question struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Title    string `gorm:"size:200;not null;index" json:"title"`
	Content  string `gorm:"type:text;not null" json:"content"`
	AuthorID uint   `gorm:"not null;index" json:"author_id"`
	Tags     string `gorm:"size:500" json:"tags"` // JSON格式存储
	Likes    int    `gorm:"default:0;index" json:"likes"`
	Views    int    `gorm:"default:0" json:"views"`
	Status   int    `gorm:"default:1;index" json:"status"` // 1:正常 0:删除

	Author User `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
}

// TableName 指定表名
func (Question) TableName() string {
	return "questions"
}

// Answer 回答模型
type Answer struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	QuestionID uint   `gorm:"not null;index" json:"question_id"`
	Content    string `gorm:"type:text;not null" json:"content"`
	AuthorID   uint   `gorm:"not null;index" json:"author_id"`
	Likes      int    `gorm:"default:0" json:"likes"`
	IsAI       bool   `gorm:"default:false" json:"is_ai"` // 是否为AI回答
	Status     int    `gorm:"default:1;index" json:"status"`

	Author   User     `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Question Question `gorm:"foreignKey:QuestionID" json:"question,omitempty"`
}

// TableName 指定表名
func (Answer) TableName() string {
	return "answers"
}

// QuestionLike 问题点赞模型
type QuestionLike struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	QuestionID uint `gorm:"not null;index:idx_question_user,unique" json:"question_id"`
	UserID     uint `gorm:"not null;index:idx_question_user,unique" json:"user_id"`
}

// TableName 指定表名
func (QuestionLike) TableName() string {
	return "question_likes"
}
