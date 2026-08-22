package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/unkwzx/url-shortener/internal/models"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Create привязан к PostgresRepository, принимает контекст и ссылку, вставляет в БД code, original_url, expires_at. RETURNING возвращает сгеню БД id, created_at
func (r *PostgresRepository) Create(ctx context.Context, link *models.Link) error {
	return r.db.QueryRow(ctx, "INSERT INTO links (code, original_url, expires_at) VALUES ($1, $2, $3) RETURNING id, created_at", link.Code, link.OriginalURL, link.ExpiresAt).Scan(&link.ID, &link.CreatedAt)
}

// GetByCode привязан к PostgresRepository, принимает контекст и код, возвращает поля в link.
func (r *PostgresRepository) GetByCode(ctx context.Context, code string) (*models.Link, error) {
	var link models.Link
	err := r.db.QueryRow(ctx, "SELECT id, code, original_url, created_at, expires_at, clicks FROM links WHERE code = $1", code).Scan(&link.ID, &link.Code, &link.OriginalURL, &link.CreatedAt, &link.ExpiresAt, &link.Clicks)
	// Проверка на кастомную ошибку.
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		} else {
			return nil, err
		}
	}
	return &link, nil
}
