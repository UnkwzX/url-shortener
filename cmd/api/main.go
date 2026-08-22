package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/unkwzx/url-shortener/internal/config"
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

}
