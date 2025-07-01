package errors

import (
	"fmt"
	"gorm.io/gorm"
)

// Error 自定义错误
type Error struct {
	Code    int
	Message string
	Err     error
}

// 预定义错误
var (
	ErrNotFound = &Error{
		Code:    NotFound,
		Message: GetMessage(NotFound),
		Err:     nil,
	}
)

// Error 实现error接口
func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap 解包错误
func (e *Error) Unwrap() error {
	return e.Err
}

// New 创建新的错误
func New(code int, err error) *Error {
	return &Error{
		Code:    code,
		Message: GetMessage(code),
		Err:     err,
	}
}

// NewWithMessage 创建带自定义消息的错误
func NewWithMessage(code int, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// FromGormError 将GORM错误转换为自定义错误
func FromGormError(err error) error {
	if err == nil {
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		return New(NotFound, nil)
	}

	return New(InternalServerError, err)
}

// IsCode 检查错误是否为指定错误码
func IsCode(err error, code int) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == code
	}
	return false
}

// GetCode 获取错误码
func GetCode(err error) int {
	if e, ok := err.(*Error); ok {
		return e.Code
	}
	return InternalServerError
}
