package repository

import (
	"context"
	"errors"

	"github.com/unkwzx/url-shortener/internal/models"
)

var ErrNotFound = errors.New("ссылка не найдена")

// LinkRepository интерфейс c 3 методами
type LinkRepository interface {
	// Метод Create() создает новую ссылку в БД
	Create(ctx context.Context, link *models.Link) error
	// метод GetByCode() ищет ссылку по коду
	GetByCode(ctx context.Context, code string) (*models.Link, error)
	// метод IncrementCount() увеличивает счетчик по коду
	IncrementCount(ctx context.Context, code string) error
}
