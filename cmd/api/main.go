package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/unkwzx/url-shortener/internal/config"
	"github.com/unkwzx/url-shortener/internal/handler"
	"github.com/unkwzx/url-shortener/internal/repository"
	"github.com/unkwzx/url-shortener/internal/service"
)

func main() {

	//Создаю логгер и ставлю его по умолчанию
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//читаем конфиг, обрабатываем ошибку и пишем лог
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Ошибка загрузки конфига!", "error", err)
		os.Exit(1)
	}
	slog.Info("Конфиг успешно загружен", "Port", cfg.ServerPort)

	//Контекст с таймаутом 5с
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Пул соединений с Postgres
	dbPool, err := pgxpool.New(ctx, cfg.DBURL)
	if err != nil {
		slog.Error("Ошибка создания пула", "error", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	// Проверка подключения к БД
	if err := dbPool.Ping(ctx); err != nil {
		slog.Error("Ошибка подключения к БД", "error", err)
		os.Exit(1)
	}
	slog.Info("Успешное подключение к БД")

	repo := repository.NewPostgresRepository(dbPool) //создаем репозиторий
	svc := service.NewLinkService(repo)              // пердаем в сервис репоизторий
	h := handler.NewHandler(svc)                     // передаем в хендлер сервис

	// создаем роутер
	mux := http.NewServeMux()
	h.RegRouters(mux)

	slog.Info("Сервер стартует", "port", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, mux); err != nil {
		slog.Error("Сервер не запустился", "error", err)
		os.Exit(1)
	}
}
