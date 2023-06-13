# Используем образ Golang для сборки
FROM golang:1.16 AS builder

WORKDIR /app

# Копируем файлы go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Собираем приложение
RUN go build -o book-service

# Отдельный этап сборки образа для минимизации размера
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем скомпилированное приложение из предыдущего этапа
COPY --from=builder /app/book-service .

# Устанавливаем порт, который будет прослушивать приложение
EXPOSE 50051

# Запускаем приложение
CMD ["./book-service"]
