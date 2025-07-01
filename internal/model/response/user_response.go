package response

import "time"

// UserResponse 用户信息响应
type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	Username  string    `json:"username" example:"johndoe"`
	Nickname  string    `json:"nickname" example:"John Doe"`
	Email     string    `json:"email" example:"john@example.com"`
	Phone     string    `json:"phone" example:"13800138000"`
	Avatar    string    `json:"avatar" example:"https://example.com/avatar.jpg"`
	Gender    int8      `json:"gender" example:"1"` // 0:未知 1:男 2:女
	Status    int8      `json:"status" example:"1"` // 0:禁用 1:启用
	CreatedAt time.Time `json:"created_at" example:"2024-01-20 10:00:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-20 10:00:00"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  UserResponse `json:"user"`
}
