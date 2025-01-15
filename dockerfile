FROM golang:latest AS builder
COPY . /app
RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app
RUN go vet && go mod tidy && go mod vendor && go build -o gin main.go

RUN apk add build-essential ffmpeg

FROM python:3.9.21-bookworm
# RUN pip config set global.index-url https://mirrors4.tuna.tsinghua.edu.cn/pypi/web/simple
pip install --upgrade pip
RUN pip install openai-whisper --break-system-packages


ENTRYPOINT ["/app/gin"]


# docker run  -dit --name whisper -v C:\Users\zen\Github\whisper-web\videos:/videos -p 8192:9001 python:3.9.21-bookworm bash
# docker run  --name ytdlp -v C:\Users\zen\Githea\ytdlp-web\videos:/videos -p 8192:9001 --rm ytdlp:latest
# docker build --debug -t ytdlp:latest .