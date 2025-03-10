# 使用 Go 官方镜像作为构建环境
FROM golang:1.23-alpine AS builder

# 安装 C 工具链和依赖
RUN apk --no-cache add gcc musl-dev


# 设置工作目录
WORKDIR /app

# 复制代码到容器中
COPY . .

# 编译 Go 应用
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -trimpath -o danmuku .

# 使用轻量级 Alpine 镜像作为运行时环境
FROM alpine:latest

# 复制编译好的二进制文件
COPY --from=builder /app/danmuku /danmuku

# 安装运行时依赖
RUN apk --no-cache add ca-certificates libc6-compat

# 设置容器启动命令
CMD ["/danmuku"]
