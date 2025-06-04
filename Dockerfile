# Используем официальный образ Golang
FROM golang:1.24-alpine

# Создаём рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod ./
RUN go mod download

# Копируем исходный код
COPY . .

# Запуск тестов по умолчанию (но в docker-compose это будет переопределено)
CMD ["go", "test", "-v", "./..."]
