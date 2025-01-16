FROM golang:latest AS builder
# 将文件复制和工作目录设置合并到一个 WORKDIR 中，WORKDIR 会自动创建目录
WORKDIR /app
COPY . /app
# 合并多个 Go 环境设置和构建命令
RUN go env -w GO111MODULE=on && \
    # go env -w GOPROXY=https://goproxy.cn,direct && \
    go vet && go mod tidy && go mod vendor && go build -o gin main.go

FROM python:3.9.21-bookworm
# 合并多个 Python 包管理和系统包管理命令
RUN apt update && apt full-upgrade -y && \
    apt install -y build-essential ffmpeg && \
    pip install --upgrade pip && \
    pip install openai-whisper --break-system-packages
# 复制文件
COPY --from=builder /app/gin /usr/bin/gin
# RUN pip config set global.index-url https://mirrors4.tuna.tsinghua.edu.cn/pypi/web/simple

ENTRYPOINT ["/usr/bin/gin"]