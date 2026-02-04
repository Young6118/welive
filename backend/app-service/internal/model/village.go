package model

import (
	"time"

	"gorm.io/gorm"
)

// Village 地球村模型
type Village struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string `gorm:"size:100;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Icon        string `gorm:"size:255" json:"icon"`
	Category    string `gorm:"size:50" json:"category"`
	MemberCount int    `gorm:"default:0" json:"member_count"`
	PostCount   int    `gorm:"default:0" json:"post_count"`
	Status      int    `gorm:"default:1" json:"status"`
}

// TableName 指定表名
func (Village) TableName() string {
	return "villages"
}

// VillageMember 村落成员模型
type VillageMember struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	VillageID uint `gorm:"not null" json:"village_id"`
	UserID    uint `gorm:"not null" json:"user_id"`
	Role      int  `gorm:"default:0" json:"role"` // 0:成员 1:管理员 2:创建者
}

// TableName 指定表名
func (VillageMember) TableName() string {
	return "village_members"
}

// Post 帖子模型
type Post struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	VillageID uint   `gorm:"not null" json:"village_id"`
	AuthorID  uint   `gorm:"not null" json:"author_id"`
	Content   string `gorm:"type:text;not null" json:"content"`
	Images    string `gorm:"type:text" json:"images"` // JSON格式存储图片URL
	Likes     int    `gorm:"default:0" json:"likes"`
	Comments  int    `gorm:"default:0" json:"comments"`
	Status    int    `gorm:"default:1" json:"status"`

	Author  User    `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Village Village `gorm:"foreignKey:VillageID" json:"village,omitempty"`
}

// TableName 指定表名
func (Post) TableName() string {
	return "posts"
}

// PostLike 帖子点赞模型
type PostLike struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	PostID uint `gorm:"not null;index" json:"post_id"`
	UserID uint `gorm:"not null;index" json:"user_id"`
}

// TableName 指定表名
func (PostLike) TableName() string {
	return "post_likes"
}
