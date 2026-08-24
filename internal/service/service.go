//Вся логика приложения живет тут.
//Валидация, генерация самого короткого кода, срок жизни ссылки.

package service

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"net/url"
	"time"

	"github.com/unkwzx/url-shortener/internal/models"
	"github.com/unkwzx/url-shortener/internal/repository"
)

// Ошибки
var ErrInvalidURL = errors.New("невалидная ссылка")
var ErrLinkExpired = errors.New("ссылка протухла")

// алфавит для генерации кода и длина кода
const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const length = 8

type LinkService struct {
	repo repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) GetOriginalURL(ctx context.Context, code string) (string, error) {
	// вызываем GetByCode
	link, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return "", err
	}
	// проверка на ссылку без срока жизни и ее протухание
	if link.ExpiresAt != nil && time.Now().After(*link.ExpiresAt) {
		return "", ErrLinkExpired
	}
	// возвращаем оригинальный урл
	return link.OriginalURL, nil

}

func (s *LinkService) CreateLink(ctx context.Context, originalURL string, ttlHours int) (*models.Link, error) {
	var expiresAt *time.Time
	// парсинг ссылки и проверка схемы
	parsedURL, err := url.ParseRequestURI(originalURL)
	if err != nil {
		return nil, err
	}
	if parsedURL.Scheme != "https" && parsedURL.Scheme != "http" {
		return nil, ErrInvalidURL
	}
	//генерация кода
	code, _ := generateCode()
	//жизнь ссылки
	if ttlHours > 0 {
		t := time.Now().Add(time.Duration(ttlHours) * time.Hour)
		expiresAt = &t
	}
	//сборка
	link := models.Link{Code: code, OriginalURL: originalURL, ExpiresAt: expiresAt}

	err = s.repo.Create(ctx, &link)
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func generateCode() (code string, err error) {
	result := make([]byte, length)
	// максмимально число для rand.Int
	maxval := big.NewInt(int64(len(charset)))

	for i := range result {
		//выбирает случайное число в диапазоне 0-max
		n, err := rand.Int(rand.Reader, maxval)
		if err != nil {
			return "", err
		}
		//заполняет i-й случайным с charset
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}
