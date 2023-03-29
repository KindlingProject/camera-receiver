FROM golang:1.20.2-bullseye AS builder
WORKDIR /build
ENV GOPROXY https://goproxy.cn
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o camera-receiver ./cmd

FROM debian:bullseye-slim AS runner
WORKDIR /app
COPY receiver-config.yml /app/
COPY --from=builder /build/camera-receiver /app/
CMD ["/app/camera-receiver", "--config=/app/receiver-config.yml"]