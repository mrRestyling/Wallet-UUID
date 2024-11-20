package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wallet/config"
	"wallet/internal/handlers"
	"wallet/internal/service"
	"wallet/internal/storage"

	"github.com/joho/godotenv"
)

func main() {

	// .env (для локальной сборки)
	if err := godotenv.Load(); err != nil {
		log.Println("err load .env", err)
	}

	// База данных
	connDB := storage.ConnectDB()
	db := storage.New(connDB)

	// Сервисный слой
	serv := service.New(db)

	// Хендлеры
	h := handlers.New(serv)
	h.SetRoutes()

	// Запуск сервера в отдельной горутине
	go h.E.Start(config.Host() + ":" + config.Port())

	// Graceful Shutdown
	// Канал, ожидающий сигнала системы
	GS := make(chan os.Signal, 1)
	signal.Notify(GS, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-GS

	// Закрытие Базы данных
	if err := db.Db.Close(); err != nil {
		log.Fatal("err close DB", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.E.Shutdown(ctx)
	if err != nil {
		log.Fatal("err shutdown serv", err)
	}

	log.Println("Server stop")

}
