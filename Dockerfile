# 使用 Go 官方镜像作为构建环境
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制代码到容器中
COPY . .

# 编译 Go 应用
RUN go build -ldflags="-s -w" -trimpath -o danmuku .

# 使用轻量级 Alpine 镜像作为运行时环境
FROM alpine:latest

# 创建持久化目录
RUN mkdir /run

# 复制编译好的二进制文件到持久化目录
COPY --from=builder /app/danmuku /run/danmuku

# 定义持久化目录
VOLUME /run

# 设置容器启动命令
CMD ["/run/danmuku"]
