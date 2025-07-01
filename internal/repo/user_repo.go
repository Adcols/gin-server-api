package repo

import (
	"github.com/Adcols/gin-server-api/internal/model/entity"
	"github.com/Adcols/gin-server-api/internal/query"
	"github.com/Adcols/gin-server-api/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepo 用户仓库
type UserRepo struct {
	db *gorm.DB
	q  *query.Query
}

// NewUserRepo 创建用户仓库实例
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
		q:  query.Use(db),
	}
}

// BeforeCreate 创建用户前的处理
func (r *UserRepo) BeforeCreate(user *entity.User) error {
	// 检查用户名是否已存在
	exists, err := r.IsUsernameExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New(errors.UsernameExists, nil)
	}

	// 检查邮箱是否已存在
	if user.Email != "" {
		exists, err := r.IsEmailExists(user.Email)
		if err != nil {
			return err
		}
		if exists {
			return errors.New(errors.EmailAlreadyExists, nil)
		}
	}

	// 检查手机号是否已存在
	if user.Phone != nil && *user.Phone != "" {
		exists, err := r.IsPhoneExists(*user.Phone)
		if err != nil {
			return err
		}
		if exists {
			return errors.New(errors.PhoneAlreadyExists, nil)
		}
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(errors.PasswordEncryptFail, err)
	}
	user.Password = string(hashedPassword)

	// 设置默认值
	if user.Status == 0 {
		user.Status = 1 // 默认启用
	}

	return nil
}

// BeforeUpdate 更新用户前的处理
func (r *UserRepo) BeforeUpdate(user *entity.User) error {
	// 如果更新了邮箱，检查是否已存在
	if user.Email != "" {
		exists, err := r.q.User.WithContext(r.db.Statement.Context).Where(
			r.q.User.Email.Eq(user.Email),
			r.q.User.ID.Neq(user.ID),
		).First()
		if err == nil && exists != nil {
			return errors.New(errors.EmailAlreadyExists, nil)
		}
	}

	// 如果更新了手机号，检查是否已存在
	if user.Phone != nil && *user.Phone != "" {
		exists, err := r.q.User.WithContext(r.db.Statement.Context).Where(
			r.q.User.Phone.Eq(*user.Phone),
			r.q.User.ID.Neq(user.ID),
		).First()
		if err == nil && exists != nil {
			return errors.New(errors.PhoneAlreadyExists, nil)
		}
	}

	return nil
}

// FindByUsername 根据用户名查找用户
func (r *UserRepo) FindByUsername(username string) (*entity.User, error) {
	user, err := r.q.User.WithContext(r.db.Statement.Context).Where(
		r.q.User.Username.Eq(username),
	).First()
	if err != nil {
		return nil, errors.FromGormError(err)
	}
	return user, nil
}

// FindByEmail 根据邮箱查找用户
func (r *UserRepo) FindByEmail(email string) (*entity.User, error) {
	user, err := r.q.User.WithContext(r.db.Statement.Context).Where(
		r.q.User.Email.Eq(email),
	).First()
	if err != nil {
		return nil, errors.FromGormError(err)
	}
	return user, nil
}

// FindByPhone 根据手机号查找用户
func (r *UserRepo) FindByPhone(phone string) (*entity.User, error) {
	user, err := r.q.User.WithContext(r.db.Statement.Context).Where(
		r.q.User.Phone.Eq(phone),
	).First()
	if err != nil {
		return nil, errors.FromGormError(err)
	}
	return user, nil
}

// FindByID 根据ID查找用户
func (r *UserRepo) FindByID(id uint32) (*entity.User, error) {
	user, err := r.q.User.WithContext(r.db.Statement.Context).Where(
		r.q.User.ID.Eq(id),
	).First()
	if err != nil {
		return nil, errors.FromGormError(err)
	}
	return user, nil
}

// Create 创建用户
func (r *UserRepo) Create(user *entity.User) error {
	// 执行创建前的处理
	if err := r.BeforeCreate(user); err != nil {
		return err
	}

	// 创建用户
	result := r.db.Create(user)
	return errors.FromGormError(result.Error)
}

// Update 更新用户信息（不包括密码）
func (r *UserRepo) Update(user *entity.User) error {
	// 执行更新前的处理
	if err := r.BeforeUpdate(user); err != nil {
		return err
	}

	// 更新用户信息，排除密码字段
	result := r.db.Model(user).Omit("password").Updates(user)
	return errors.FromGormError(result.Error)
}

// UpdatePassword 更新用户密码
func (r *UserRepo) UpdatePassword(id uint32, oldPassword, newPassword string) error {
	// 查找用户
	user, err := r.FindByID(id)
	if err != nil {
		return err
	}

	// 验证旧密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New(errors.PasswordIncorrect, nil)
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(errors.PasswordEncryptFail, err)
	}

	// 更新密码
	result := r.db.Model(user).Update("password", string(hashedPassword))
	return errors.FromGormError(result.Error)
}

// Delete 删除用户
func (r *UserRepo) Delete(id uint32) error {
	result := r.db.Delete(&entity.User{}, id)
	return errors.FromGormError(result.Error)
}

// IsUsernameExists 检查用户名是否存在
func (r *UserRepo) IsUsernameExists(username string) (bool, error) {
	count, err := r.q.User.WithContext(r.db.Statement.Context).Where(
		r.q.User.Username.Eq(username),
	).Count()
	if err != nil {
		return false, errors.FromGormError(err)
	}
	return count > 0, nil
}

// IsEmailExists 检查邮箱是否存在
func (r *UserRepo) IsEmailExists(email string) (bool, error) {
	count, err := r.q.User.WithContext(r.db.Statement.Context).Where(
		r.q.User.Email.Eq(email),
	).Count()
	if err != nil {
		return false, errors.FromGormError(err)
	}
	return count > 0, nil
}

// IsPhoneExists 检查手机号是否存在
func (r *UserRepo) IsPhoneExists(phone string) (bool, error) {
	count, err := r.q.User.WithContext(r.db.Statement.Context).Where(
		r.q.User.Phone.Eq(phone),
	).Count()
	if err != nil {
		return false, errors.FromGormError(err)
	}
	return count > 0, nil
}

// VerifyPassword 验证用户密码
func (r *UserRepo) VerifyPassword(user *entity.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New(errors.PasswordIncorrect, nil)
	}
	return nil
}
