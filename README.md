# Gin Server API

基于Gin+GORM+Redis的企业级后端API服务，用于开发APP端、小程序端、H5端和后台管理端的接口。

## 项目特点

- 基于Gin框架，性能卓越
- 使用GORM进行数据库操作，支持多种数据库
- 集成Redis缓存，提高系统性能
- JWT认证，保障API安全
- 统一的响应格式和错误处理
- 完善的日志记录
- Swagger文档自动生成
- 支持跨域请求
- 优雅关闭服务

## 项目结构

```
├── api                 # API层
│   ├── controllers     # 控制器
│   ├── middlewares     # 中间件
│   └── routes          # 路由
├── config              # 配置
├── docs                # 文档
├── internal            # 内部包
│   ├── models          # 数据模型
│   ├── repositories    # 数据仓库
│   └── services        # 业务服务
├── pkg                 # 公共包
│   ├── logger          # 日志
│   ├── response        # 响应
│   └── utils           # 工具
├── scripts             # 脚本
├── config.yaml         # 配置文件
├── go.mod              # Go模块文件
├── go.sum              # Go依赖校验文件
└── main.go             # 主程序入口
```

## 快速开始

### 环境要求

- Go 1.18+
- MySQL 5.7+
- Redis 6.0+

### 安装依赖

```bash
go mod tidy
```

### 配置

修改`config.yaml`文件，配置数据库和Redis连接信息：

```yaml
app:
  name: gin-server-api
  mode: development
  port: 8080
  version: 1.0.0

database:
  driver: mysql
  host: localhost
  port: 3306
  username: root
  password: password
  database: gin_server
  charset: utf8mb4

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

jwt:
  secret: "your-secret-key"
  expiration: 86400 # 24小时
```

### 生成Swagger文档

```bash
swag init
```

### 运行

```bash
go run main.go
```

访问 http://localhost:8080/swagger/index.html 查看API文档。

## API示例

### 用户注册

```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","email":"test@example.com"}'
```

### 用户登录

```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'
```

### 获取用户信息

```bash
curl -X GET http://localhost:8080/api/v1/users/info \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## 开发指南

### 添加新的控制器

1. 在`api/controllers`目录下创建新的控制器文件
2. 在`api/routes/router.go`中注册路由

### 添加新的模型

1. 在`internal/models`目录下创建新的模型文件
2. 在服务启动时自动迁移数据库表结构

### 添加新的服务

1. 在`internal/services`目录下创建新的服务文件
2. 在控制器中使用服务

## 部署

### Docker部署

创建Dockerfile：

```dockerfile
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-server-api .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gin-server-api .
COPY --from=builder /app/config.yaml .

EXPOSE 8080

CMD ["./gin-server-api"]
```

构建和运行Docker容器：

```bash
docker build -t gin-server-api .
docker run -p 8080:8080 gin-server-api
```

## 许可证

[MIT](LICENSE)
