#1 этам билдер
FROM golang:1.26.5-alpine AS builder
WORKDIR /app
#сначала прокинул go.mod
COPY go.mod go.sum ./
RUN go mod download
# . все из корня в /app
COPY . .
#отключаем с - зависимсости, бинарник под линукс,
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/api

#2 этап
FROM alpine:latest
WORKDIR /app
#копируем бинарник сюда
COPY --from=builder /app/server .
COPY --from=builder /app/migrations ./migrations
#порт
EXPOSE 8080
# команда запуска
CMD ["./server"]
