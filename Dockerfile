FROM golang:1.20.2

WORKDIR /usr/src/app

ENV GOPROXY https://goproxy.cn
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/camera-receiver ./cmd
COPY receiver-config.yml /etc/config/

CMD ["camera-receiver", "--config=/etc/config/receiver-config.yml"]