package model

import (
	"time"

	"gorm.io/gorm"
)

// Chat 聊天会话模型
type Chat struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID     uint   `gorm:"not null" json:"user_id"`
	ReceiverID uint   `gorm:"not null" json:"receiver_id"` // 对方ID（用户/智能体/员工）
	Type       string `gorm:"size:20;not null" json:"type"` // user/agent/employee
	LastMessage string `gorm:"type:text" json:"last_message"`
	UnreadCount int    `gorm:"default:0" json:"unread_count"`
}

// TableName 指定表名
func (Chat) TableName() string {
	return "chats"
}

// Message 消息模型
type Message struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`

	ChatID     uint   `gorm:"not null" json:"chat_id"`
	SenderID   uint   `gorm:"not null" json:"sender_id"`
	Content    string `gorm:"type:text;not null" json:"content"`
	Type       string `gorm:"size:20;default:'text'" json:"type"` // text/image/file
	Status     int    `gorm:"default:1" json:"status"` // 1:已发送 2:已读
}

// TableName 指定表名
func (Message) TableName() string {
	return "messages"
}
