# Устанавливаем базовый образ для сборки
FROM golang:1.18 as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для скачивания зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости Go
RUN go mod download

# Копируем остальные файлы в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o myapp .

# Используем минимальный базовый образ для запуска приложения
FROM debian:buster-slim

# Устанавливаем необходимые зависимости
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию для финального контейнера
WORKDIR /app

# Копируем скомпилированное приложение из предыдущего этапа
COPY --from=builder /app/myapp .

# Указываем команду для запуска приложения
CMD ["./myapp"]