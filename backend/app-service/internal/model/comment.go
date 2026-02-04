package model

import (
	"time"

	"gorm.io/gorm"
)

// Comment 评论模型
type Comment struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	TargetID   uint   `gorm:"not null" json:"target_id"`   // 目标ID
	TargetType string `gorm:"size:20;not null" json:"target_type"` // question/answer/note/post
	Content    string `gorm:"type:text;not null" json:"content"`
	AuthorID   uint   `gorm:"not null" json:"author_id"`
	ParentID   *uint  `gorm:"" json:"parent_id"` // 父评论ID，用于回复
	Likes      int    `gorm:"default:0" json:"likes"`
	Status     int    `gorm:"default:1" json:"status"`

	Author User `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
}

// TableName 指定表名
func (Comment) TableName() string {
	return "comments"
}

// CommentLike 评论点赞模型
type CommentLike struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	CommentID uint `gorm:"not null;index" json:"comment_id"`
	UserID    uint `gorm:"not null;index" json:"user_id"`
}

// TableName 指定表名
func (CommentLike) TableName() string {
	return "comment_likes"
}
