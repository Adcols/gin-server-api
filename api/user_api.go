package controllers

import (
	"github.com/Adcols/gin-server-api/internal/model"
	"github.com/Adcols/gin-server-api/internal/service"
	"github.com/Adcols/gin-server-api/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserController 用户控制器
type UserController struct {
	userService *services.UserService
}

// NewUserController 创建用户控制器
func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body models.User true "用户信息"
// @Success 200 {object} response.Response
// @Router /api/v1/users/register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var user model.User

	// 绑定请求参数
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.ParameterError(ctx, err.Error())
		return
	}

	// 参数验证
	if user.Username == "" || user.Password == "" {
		response.ParameterError(ctx, "用户名和密码不能为空")
		return
	}

	// 注册用户
	err := c.userService.Register(&user)
	if err != nil {
		response.FailWithMessage(ctx, response.CodeFailed, err.Error(), nil)
		return
	}

	response.Success(ctx, gin.H{"user_id": user.ID})
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body map[string]string true "登录信息"
// @Success 200 {object} response.Response
// @Router /api/v1/users/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	// 绑定请求参数
	var params struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ParameterError(ctx, err.Error())
		return
	}

	// 登录验证
	token, user, err := c.userService.Login(params.Username, params.Password)
	if err != nil {
		response.FailWithMessage(ctx, response.CodeFailed, err.Error(), nil)
		return
	}

	response.Success(ctx, gin.H{
		"token": token,
		"user":  user,
	})
}

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response
// @Router /api/v1/users/info [get]
func (c *UserController) GetUserInfo(ctx *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, nil)
		return
	}

	// 获取用户信息
	user, err := c.userService.GetUserByID(userID.(uint))
	if err != nil {
		response.FailWithMessage(ctx, response.CodeFailed, err.Error(), nil)
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
// @Param data body models.User true "用户信息"
// @Success 200 {object} response.Response
// @Router /api/v1/users/info [put]
func (c *UserController) UpdateUserInfo(ctx *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, nil)
		return
	}

	// 绑定请求参数
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.ParameterError(ctx, err.Error())
		return
	}

	// 设置用户ID
	user.ID = userID.(uint)

	// 更新用户信息
	err := c.userService.UpdateUser(&user)
	if err != nil {
		response.FailWithMessage(ctx, response.CodeFailed, err.Error(), nil)
		return
	}

	response.Success(ctx, nil)
}

// UpdatePassword 更新用户密码
// @Summary 更新用户密码
// @Description 更新用户密码接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body map[string]string true "密码信息"
// @Success 200 {object} response.Response
// @Router /api/v1/users/password [put]
func (c *UserController) UpdatePassword(ctx *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		response.Unauthorized(ctx, nil)
		return
	}

	// 绑定请求参数
	var params struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ParameterError(ctx, err.Error())
		return
	}

	// 更新密码
	err := c.userService.UpdatePassword(userID.(uint), params.OldPassword, params.NewPassword)
	if err != nil {
		response.FailWithMessage(ctx, response.CodeFailed, err.Error(), nil)
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
// @Success 200 {object} response.Response
// @Router /api/v1/users/{id} [get]
func (c *UserController) GetUserByID(ctx *gin.Context) {
	// 获取路径参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParameterError(ctx, "无效的用户ID")
		return
	}

	// 获取用户信息
	user, err := c.userService.GetUserByID(uint(id))
	if err != nil {
		response.FailWithMessage(ctx, response.CodeFailed, err.Error(), nil)
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
func (c *UserController) DeleteUser(ctx *gin.Context) {
	// 获取路径参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.ParameterError(ctx, "无效的用户ID")
		return
	}

	// 删除用户
	err = c.userService.DeleteUser(uint(id))
	if err != nil {
		response.FailWithMessage(ctx, response.CodeFailed, err.Error(), nil)
		return
	}

	response.Success(ctx, nil)
}
