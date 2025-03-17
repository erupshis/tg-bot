# Этап 1: Сборка приложения
FROM golang:1.24-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей
COPY . .

# Собираем приложение
RUN go mod tidy CGO_ENABLED=0 GOOS=linux go build -a -o bot ./cmd/main.go

# Этап 2: Создание финального образа
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарный файл из этапа сборки
COPY --from=builder /app/bot bot

EXPOSE 8080
# Указываем команду для запуска
CMD ["./bot"]

