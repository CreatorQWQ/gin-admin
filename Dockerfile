# 使用官方 Go 镜像作为基础（多阶段构建，体积小）
FROM golang:1.22-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 先下载依赖（缓存优化）
COPY go.mod go.sum ./
RUN go mod download

# 复制全部代码
COPY . .

# 编译二进制（生产优化：静态链接、无 CGO）
RUN CGO_ENABLED=0 GOOS=linux go build -o gin-admin ./main.go

# 第二阶段：运行时镜像（极小 alpine）
FROM alpine:latest

# 安装 ca-certificates（如果需要 HTTPS）
RUN apk --no-cache add ca-certificates

WORKDIR /app

# 从 builder 阶段复制编译好的二进制
COPY --from=builder /app/gin-admin .

# 暴露端口
EXPOSE 8080

# 运行命令
CMD ["./gin-admin"]