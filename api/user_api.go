package controllers

import (
	"strconv"

	"github.com/Adcols/gin-server-api/internal/model/request"
	"github.com/Adcols/gin-server-api/internal/service"
	"github.com/Adcols/gin-server-api/internal/wire"
	"github.com/Adcols/gin-server-api/pkg/response"
	"github.com/gin-gonic/gin"
)

// UserApi 用户控制器
type UserApi struct {
	userService *service.UserService
}

// NewUserApi 创建用户控制器
func NewUserApi() *UserApi {
	return &UserApi{
		userService: wire.InitializeUserService(),
	}
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body request.RegisterRequest true "用户注册信息"
// @Success 200 {object} response.Response{data=response.UserResponse}
// @Router /api/v1/users/register [post]
func (c *UserApi) Register(ctx *gin.Context) {
	var req request.RegisterRequest

	// 绑定请求参数
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParameterError(ctx, "无效的请求参数")
		return
	}

	// 注册用户
	user, err := c.userService.Register(&req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, user)
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body request.LoginRequest true "登录信息"
// @Success 200 {object} response.Response{data=response.LoginResponse}
// @Router /api/v1/users/login [post]
func (c *UserApi) Login(ctx *gin.Context) {
	var req request.LoginRequest

	// 绑定请求参数
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParameterError(ctx, "无效的请求参数")
		return
	}

	// 登录
	resp, err := c.userService.Login(&req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, resp)
}

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=response.UserResponse}
// @Router /api/v1/users/info [get]
func (c *UserApi) GetUserInfo(ctx *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx)
		return
	}

	// 获取用户信息
	user, err := c.userService.GetUserByID(userID.(uint32))
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, user)
}

// UpdateUserInfo 更新用户信息
// @Summary 更新用户信息
// @Description 更新用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body request.UpdateUserRequest true "用户信息"
// @Success 200 {object} response.Response{data=response.UserResponse}
// @Router /api/v1/users/info [put]
func (c *UserApi) UpdateUserInfo(ctx *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx)
		return
	}

	// 绑定请求参数
	var req request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParameterError(ctx, "无效的请求参数")
		return
	}

	// 更新用户信息
	user, err := c.userService.UpdateUser(userID.(uint32), &req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, user)
}

// UpdatePassword 更新用户密码
// @Summary 更新用户密码
// @Description 更新用户密码接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body request.UpdatePasswordRequest true "密码信息"
// @Success 200 {object} response.Response
// @Router /api/v1/users/password [put]
func (c *UserApi) UpdatePassword(ctx *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx)
		return
	}

	// 绑定请求参数
	var req request.UpdatePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParameterError(ctx, "无效的请求参数")
		return
	}

	// 更新密码
	err := c.userService.UpdatePassword(userID.(uint32), &req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, nil)
}

// GetUserByID 根据ID获取用户
// @Summary 根据ID获取用户
// @Description 根据ID获取用户接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response{data=response.UserResponse}
// @Router /api/v1/users/{id} [get]
func (c *UserApi) GetUserByID(ctx *gin.Context) {
	// 获取路径参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParameterError(ctx, "无效的用户ID")
		return
	}

	// 获取用户信息
	user, err := c.userService.GetUserByID(uint32(id))
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, user)
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除用户接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response
// @Router /api/v1/users/{id} [delete]
func (c *UserApi) DeleteUser(ctx *gin.Context) {
	// 获取路径参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParameterError(ctx, "无效的用户ID")
		return
	}

	// 删除用户
	err = c.userService.DeleteUser(uint32(id))
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, nil)
}
