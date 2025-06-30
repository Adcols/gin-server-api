package services

import (
	"errors"

	models "github.com/Adcols/gin-server-api/internal/model"
	repositories "github.com/Adcols/gin-server-api/internal/repository"
	"github.com/Adcols/gin-server-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 用户服务
type UserService struct {
	userRepo *repositories.UserRepository
}

// NewUserService 创建用户服务
func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepository(),
	}
}

// Register 用户注册
func (s *UserService) Register(user *models.User) error {
	// 检查用户名是否已存在
	exists, err := s.userRepo.IsUsernameExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	if user.Email != "" {
		exists, err := s.userRepo.IsEmailExists(user.Email)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("邮箱已存在")
		}
	}

	// 检查手机号是否已存在
	if user.Phone != "" {
		exists, err := s.userRepo.IsPhoneExists(user.Phone)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("手机号已存在")
		}
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// 创建用户
	return s.userRepo.Create(user)
}

// Login 用户登录
func (s *UserService) Login(username, password string) (string, *models.User, error) {
	// 查询用户
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, errors.New("用户不存在")
		}
		return "", nil, err
	}

	// 检查用户状态
	if user.Status == 0 {
		return "", nil, errors.New("用户已被禁用")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", nil, errors.New("密码错误")
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	return user, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(user *models.User) error {
	// 更新用户信息，不更新密码
	return s.userRepo.Update(user)
}

// UpdatePassword 更新用户密码
func (s *UserService) UpdatePassword(id uint, oldPassword, newPassword string) error {
	// 查询用户
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	// 验证旧密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("旧密码错误")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	return s.userRepo.UpdatePassword(id, string(hashedPassword))
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
