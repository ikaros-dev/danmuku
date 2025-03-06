# 使用官方的 Go 镜像作为基础镜像
FROM golang:1.21-alpine AS builder

# 安装 SQLite 开发库
RUN apk add --no-cache gcc musl-dev sqlite-dev

# 设置工作目录
WORKDIR /app

# 复制 Go 模块文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 启用 CGO 并构建程序
RUN CGO_ENABLED=1 go build -o danmuku .

# 使用轻量级的 Alpine 镜像作为运行环境
FROM alpine:latest

# 安装 SQLite 运行时库
RUN apk add --no-cache sqlite

# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/danmuku .

# 复制配置文件或其他资源
COPY config ./config

# 设置容器启动命令
CMD ["./danmuku"]
