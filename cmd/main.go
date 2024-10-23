package main

import (
	"context"
	"crossword/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

// gracefulShutdown корректно завершает работу сервера
func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Канал для получения системных сигналов
	quit := make(chan os.Signal, 1)

	// Подписываемся на SIGINT и SIGTERM
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Блокируем выполнение до получения сигнала
	<-quit
	log.Println("Получен сигнал завершения. Остановка сервера...")

	// Устанавливаем таймаут для корректного завершения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := apiServer.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %v", err)
	}

	close(done)
}

func main() {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Создаем сервер
	apiServer := server.NewServer(":8080")

	// Канал для уведомления о завершении работы
	done := make(chan bool, 1)

	// Запуск сервера в отдельной горутине
	go func() {
		if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()
	log.Println("Сервер запущен на http://localhost:8080")

	// Настраиваем ловушку для сигналов завершения (Ctrl+C, SIGTERM)
	gracefulShutdown(apiServer, done)

	// Ожидаем завершения
	<-done
	log.Println("Сервер завершил работу.")
}
