FROM golang:1.16.3 as builder

# 设置容器环境变量
ENV GOPROXY=https://goproxy.cn
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY . /app

WORKDIR /app

RUN go get -u github.com/fvbock/endless

RUN go build -ldflags="-s -w" -installsuffix cgo -o go_init

FROM alpine as prod

# 开放端口
EXPOSE 9091

# 创建一个目录
RUN mkdir -p /app/logs

RUN chmod 666 /app/logs

RUN ls /app

COPY --from=builder /app/go_init /app

COPY --from=builder /app/config.toml /app

# 启动
CMD ["/app/go_init", "--config=/app/config.toml"]
