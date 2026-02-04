package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Username     string `gorm:"uniqueIndex;size:50;not null" json:"username"`
	PasswordHash string `gorm:"size:255;not null" json:"-"`
	Email        string `gorm:"uniqueIndex;size:100" json:"email"`
	Avatar       string `gorm:"size:255" json:"avatar"`
	Bio          string `gorm:"size:500" json:"bio"`
	Status       int    `gorm:"default:1" json:"status"` // 1:正常 0:禁用
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
