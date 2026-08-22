package main

import (
	"log/slog"
	"os"

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

}
