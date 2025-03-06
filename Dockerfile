# 未测试不可用
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -trimpath -o danmuku .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/danmuku .
CMD ["./danmuku"]
