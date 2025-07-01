package request

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Password string `json:"password" binding:"required" example:"password123"`
	Nickname string `json:"nickname" example:"John Doe"`
	Email    string `json:"email" binding:"omitempty,email" example:"john@example.com"`
	Phone    string `json:"phone" binding:"omitempty" example:"13800138000"`
	Avatar   string `json:"avatar" example:"https://example.com/avatar.jpg"`
	Gender   int8   `json:"gender" example:"1"` // 0:未知 1:男 2:女
}

// LoginRequest 用户登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// UpdatePasswordRequest 更新密码请求
type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" example:"oldpassword123"`
	NewPassword string `json:"new_password" binding:"required" example:"newpassword123"`
}

// UpdateUserRequest 更新用户信息请求
type UpdateUserRequest struct {
	Nickname string `json:"nickname" example:"John Doe"`
	Email    string `json:"email" binding:"omitempty,email" example:"john@example.com"`
	Phone    string `json:"phone" binding:"omitempty" example:"13800138000"`
	Avatar   string `json:"avatar" example:"https://example.com/avatar.jpg"`
	Gender   int8   `json:"gender" example:"1"` // 0:未知 1:男 2:女
}
