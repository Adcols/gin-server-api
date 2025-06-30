package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Adcols/gin-server-api/config"
	"github.com/Adcols/gin-server-api/routes"
)

// @title Gin Server API
// @version 1.0
// @description 基于Gin+GORM+Redis的企业级后端API服务
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化数据库
	config.InitMySQL()
	defer config.CloseMySQL()

	// 初始化Redis
	config.InitRedis()
	defer config.CloseRedis()

	// 设置路由
	router := routes.SetupRouter()

	// 创建HTTP服务器
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GlobalConfig.App.Port),
		Handler: router,
	}

	// 启动HTTP服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("启动服务器失败: %v", err)
		}
	}()

	slog.Info("服务器启动成功: http://localhost:", config.GlobalConfig.App.Port)

	// 等待中断信号优雅关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("服务器关闭中...")

	// 设置关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭HTTP服务器
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("服务器关闭失败: %v", err)
	}

	slog.Info("服务器已关闭")
}
