package service

import (
	"github.com/Adcols/gin-server-api/internal/model/entity"
	"github.com/Adcols/gin-server-api/internal/model/request"
	"github.com/Adcols/gin-server-api/internal/model/response"
	"github.com/Adcols/gin-server-api/internal/repo"
	"github.com/Adcols/gin-server-api/pkg/errors"
	"github.com/Adcols/gin-server-api/pkg/utils"
)

// UserService 用户服务
type UserService struct {
	UserRepo *repo.UserRepo
}

// Register 用户注册
func (s *UserService) Register(req *request.RegisterRequest) (*response.UserResponse, error) {
	// 创建用户模型
	user := &entity.User{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    &req.Phone,
		Avatar:   req.Avatar,
		Gender:   uint32(req.Gender),
	}

	// 创建用户（包含业务逻辑验证）
	err := s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	// 转换为响应对象
	resp := &response.UserResponse{
		ID:       uint(user.ID),
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    *user.Phone,
		Avatar:   user.Avatar,
		Gender:   int8(user.Gender),
		Status:   int8(user.Status),
		// CreatedAt: *user.CreatedAt,
		// UpdatedAt: *user.UpdatedAt,
	}

	return resp, nil
}

// Login 用户登录
func (s *UserService) Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	// 查询用户
	user, err := s.UserRepo.FindByUsername(req.Username)
	if err != nil {
		if _, ok := err.(*errors.Error); ok && err.Error() == errors.ErrNotFound.Error() {
			return nil, errors.New(errors.UserNotFound, nil)
		}
		return nil, err
	}

	// 检查用户状态
	if user.Status == 0 {
		return nil, errors.New(errors.UserDisabled, nil)
	}

	// 验证密码
	err = s.UserRepo.VerifyPassword(user, req.Password)
	if err != nil {
		return nil, errors.New(errors.PasswordIncorrect, nil)
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(uint(user.ID), user.Username)
	if err != nil {
		return nil, errors.New(errors.InternalServerError, err)
	}

	// 转换为响应对象
	resp := &response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			ID:        uint(user.ID),
			Username:  user.Username,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Phone:     *user.Phone,
			Avatar:    user.Avatar,
			Gender:    int8(user.Gender),
			Status:    int8(user.Status),
			CreatedAt: *user.CreatedAt,
			UpdatedAt: *user.UpdatedAt,
		},
	}

	return resp, nil
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id uint32) (*response.UserResponse, error) {
	// 获取用户信息
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		if _, ok := err.(*errors.Error); ok && err.Error() == errors.ErrNotFound.Error() {
			return nil, errors.New(errors.UserNotFound, nil)
		}
		return nil, errors.New(errors.InternalServerError, err)
	}

	// 转换为响应对象
	resp := &response.UserResponse{
		ID:        uint(user.ID),
		Username:  user.Username,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Phone:     *user.Phone,
		Avatar:    user.Avatar,
		Gender:    int8(user.Gender),
		Status:    int8(user.Status),
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}

	return resp, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(id uint32, req *request.UpdateUserRequest) (*response.UserResponse, error) {
	// 查询用户
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		if _, ok := err.(*errors.Error); ok && err.Error() == errors.ErrNotFound.Error() {
			return nil, errors.New(errors.UserNotFound, nil)
		}
		return nil, errors.New(errors.InternalServerError, err)
	}

	// 更新用户信息
	user.Nickname = req.Nickname
	user.Email = req.Email
	user.Phone = &req.Phone
	user.Avatar = req.Avatar
	user.Gender = uint32(req.Gender)

	// 保存更新（包含业务逻辑验证）
	err = s.UserRepo.Update(user)
	if err != nil {
		return nil, errors.New(errors.InternalServerError, err)
	}

	// 转换为响应对象
	resp := &response.UserResponse{
		ID:        uint(user.ID),
		Username:  user.Username,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Phone:     *user.Phone,
		Avatar:    user.Avatar,
		Gender:    int8(user.Gender),
		Status:    int8(user.Status),
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}

	return resp, nil
}

// UpdatePassword 更新用户密码
func (s *UserService) UpdatePassword(id uint32, req *request.UpdatePasswordRequest) error {
	return s.UserRepo.UpdatePassword(id, req.OldPassword, req.NewPassword)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id uint32) error {
	return s.UserRepo.Delete(id)
}
