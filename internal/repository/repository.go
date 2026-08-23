package repository

import (
	"context"
	"errors"

	"github.com/unkwzx/url-shortener/internal/models"
)

var ErrNotFound = errors.New("ссылка не найдена")

// LinkRepository интерфейс c 3 методами
type LinkRepository interface {

	// Create создает новую ссылку в БД
	Create(ctx context.Context, link *models.Link) error

	// GetByCode ищет ссылку по коду
	GetByCode(ctx context.Context, code string) (*models.Link, error)

	// IncrementCount увеличивает счетчик по коду
	IncrementClicks(ctx context.Context, code string) error
}
