package main

import (
	"log/slog"
	"os"
)

func main() {

	//Создаю логгер и ставлю его по умолчанию
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

}
