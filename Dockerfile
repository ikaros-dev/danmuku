# 使用 Go 官方镜像作为构建环境
FROM golang:1.23-alpine AS builder

# 安装 SQLite 开发库
RUN apk add --no-cache gcc musl-dev sqlite-dev

# 设置工作目录
WORKDIR /app

# 启用 CGO 并构建程序
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -trimpath -o danmuku .

# 复制 Go 模块文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制代码到容器中
COPY . .

# 使用轻量级 Alpine 镜像作为运行时环境
FROM alpine:latest

# 创建持久化目录
RUN mkdir /data

# 复制编译好的二进制文件到持久化目录
COPY --from=builder /app/danmuku /data/danmuku

# 定义持久化目录
VOLUME /data

# 设置容器启动命令
CMD ["/data/danmuku"]


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
RUN CGO_ENABLED=1 go build -o ikaros-danmuku .

# 使用轻量级的 Alpine 镜像作为运行环境
FROM alpine:latest

# 安装 SQLite 运行时库
RUN apk add --no-cache sqlite

# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/ikaros-danmuku .

# 复制配置文件或其他资源
COPY config ./config

# 设置容器启动命令
CMD ["./ikaros-danmuku"]
