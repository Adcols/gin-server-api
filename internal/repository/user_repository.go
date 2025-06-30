package repositorie

import (
	"github.com/Adcols/gin-server-api/config"
	"github.com/Adcols/gin-server-api/internal/model"
)

// UserRepository 用户仓库
type UserRepository struct{}

// NewUserRepository 创建用户仓库
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// FindByUsername 根据用户名查找用户
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	result := config.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindByEmail 根据邮箱查找用户
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindByPhone 根据手机号查找用户
func (r *UserRepository) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	result := config.DB.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindByID 根据ID查找用户
func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Create 创建用户
func (r *UserRepository) Create(user *model.User) error {
	result := config.DB.Create(user)
	return result.Error
}

// Update 更新用户信息（不包括密码）
func (r *UserRepository) Update(user *model.User) error {
	result := config.DB.Model(user).Omit("password").Updates(user)
	return result.Error
}

// UpdatePassword 更新用户密码
func (r *UserRepository) UpdatePassword(id uint, password string) error {
	result := config.DB.Model(&model.User{}).Where("id = ?", id).Update("password", password)
	return result.Error
}

// Delete 删除用户
func (r *UserRepository) Delete(id uint) error {
	result := config.DB.Delete(&model.User{}, id)
	return result.Error
}

// IsUsernameExists 检查用户名是否存在
func (r *UserRepository) IsUsernameExists(username string) (bool, error) {
	var count int64
	result := config.DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

// IsEmailExists 检查邮箱是否存在
func (r *UserRepository) IsEmailExists(email string) (bool, error) {
	var count int64
	result := config.DB.Model(&model.User{}).Where("email = ?", email).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

// IsPhoneExists 检查手机号是否存在
func (r *UserRepository) IsPhoneExists(phone string) (bool, error) {
	var count int64
	result := config.DB.Model(&model.User{}).Where("phone = ?", phone).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
