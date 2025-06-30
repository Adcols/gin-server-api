package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:50;not null;unique" json:"username"`
	Password  string         `gorm:"size:100;not null" json:"-"` // 不返回密码
	Nickname  string         `gorm:"size:50" json:"nickname"`
	Email     string         `gorm:"size:100;unique" json:"email"`
	Phone     string         `gorm:"size:20;unique" json:"phone"`
	Avatar    string         `gorm:"size:255" json:"avatar"`
	Gender    int8           `gorm:"default:0" json:"gender"` // 0:未知 1:男 2:女
	Status    int8           `gorm:"default:1" json:"status"` // 0:禁用 1:启用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 创建前的钩子
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 可以在这里进行密码加密等操作
	return nil
}

// BeforeUpdate 更新前的钩子
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	// 可以在这里进行密码加密等操作
	return nil
}
