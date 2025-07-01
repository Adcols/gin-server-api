package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Adcols/gin-server-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 全局日志对象
var Logger *zap.Logger
var SugaredLogger *zap.SugaredLogger

// InitLogger 初始化日志
func InitLogger() {
	// 创建日志目录
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			fmt.Printf("创建日志目录失败: %v\n", err)
			os.Exit(1)
		}
	}

	// 设置日志级别
	var level zapcore.Level
	if config.GlobalConfig.App.Mode == "production" {
		level = zapcore.InfoLevel
	} else {
		level = zapcore.DebugLevel
	}

	// 创建日志文件
	logFilePath := filepath.Join(logDir, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")))
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("打开日志文件失败: %v\n", err)
		os.Exit(1)
	}

	// 设置编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置输出
	var core zapcore.Core

	// 开发模式同时输出到控制台和文件
	if config.GlobalConfig.App.Mode == "development" {
		// 控制台输出
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleOutput := zapcore.AddSync(os.Stdout)

		// 文件输出
		fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
		fileOutput := zapcore.AddSync(logFile)

		// 多输出
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleOutput, level),
			zapcore.NewCore(fileEncoder, fileOutput, level),
		)
	} else {
		// 生产环境只输出到文件
		fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
		fileOutput := zapcore.AddSync(logFile)
		core = zapcore.NewCore(fileEncoder, fileOutput, level)
	}

	// 创建Logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	SugaredLogger = Logger.Sugar()

	Info("日志初始化成功")
}

// CloseLogger 关闭日志
func CloseLogger() {
	if Logger != nil {
		Logger.Sync()
	}
}

// Debug 调试日志
func Debug(format string, v ...interface{}) {
	SugaredLogger.Debugf(format, v...)
}

// Info 信息日志
func Info(format string, v ...interface{}) {
	SugaredLogger.Infof(format, v...)
}

// Warn 警告日志
func Warn(format string, v ...interface{}) {
	SugaredLogger.Warnf(format, v...)
}

// Error 错误日志
func Error(format string, v ...interface{}) {
	SugaredLogger.Errorf(format, v...)
}

// Fatal 致命错误日志
func Fatal(format string, v ...interface{}) {
	SugaredLogger.Fatalf(format, v...)
}