package routes

import (
	"net/http"

	"github.com/Adcols/gin-server-api/api"
	"github.com/Adcols/gin-server-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	// 创建路由
	r := gin.New()

	// 使用中间件
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API路由组
	api := r.Group("/api/v1")

	// 用户控制器
	userController := controllers.NewUserController()

	// 用户路由
	userRoutes := api.Group("/users")
	{
		// 公开路由
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)

		// 需要认证的路由
		userAuthRoutes := userRoutes.Group("/")
		userAuthRoutes.Use(middleware.JWTAuth())
		{
			userAuthRoutes.GET("/info", userController.GetUserInfo)
			userAuthRoutes.PUT("/info", userController.UpdateUserInfo)
			userAuthRoutes.PUT("/password", userController.UpdatePassword)
			userAuthRoutes.GET("/:id", userController.GetUserByID)
			userAuthRoutes.DELETE("/:id", userController.DeleteUser)
		}
	}

	return r
}
