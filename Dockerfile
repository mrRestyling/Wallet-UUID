FROM golang:latest

# Установка рабочей директории
WORKDIR /app

# Копирование всего проекта
COPY . .

# Установка зависимостей
RUN go mod tidy
RUN go mod vendor

# Сборка приложения
RUN go build -o walletAPP cmd/app/main.go

# Установка команды по умолчанию
CMD ["/app/walletAPP"]

