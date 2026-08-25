package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/unkwzx/url-shortener/internal/repository"
	"github.com/unkwzx/url-shortener/internal/service"
)

// от клиента
type CreateLinkRequest struct {
	URL      string `json:"url"`
	ttlHours int    `json:"ttl_hours"`
}

// к клиенту
type CreateLinkResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
}

type Handler struct {
	service service.LinkService
}

func NewHandler(service service.LinkService) *Handler {
	return &Handler{service: service}
}

// запрос -> h.service.CreateLink (генерация кода, сборка link) -> s.repo.Create(запись link в БД) -> ответ response
func (h *Handler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var req CreateLinkRequest                   // сюда десериализуем тело запроса
	err := json.NewDecoder(r.Body).Decode(&req) // десериализуем json
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid body request") // откдидываем ошибку
		return                                                      // Выходим
	}
	// вызываем метод service CreateLink, передаем контекст, ссылку, срок жизни
	link, err := h.service.CreateLink(r.Context(), req.URL, req.ttlHours)
	//обрабатываем ошибку
	if err != nil {
		if errors.Is(err, service.ErrInvalidURL) {
			sendError(w, http.StatusBadRequest, err.Error())
			return
		}
		sendError(w, http.StatusInternalServerError, "error create url")
		return
	}
	// формируем ответ
	response := CreateLinkResponse{
		Code:     link.Code,
		ShortURL: fmt.Sprintf("https://localhost:8080/%s", link.Code),
	}
	// отправляем ответ клиенту
	sendJSON(w, http.StatusCreated, response)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	// заюираем код
	code := r.PathValue("code")
	//идем до бд, возвращаем ссылку по переданному коду
	originalURL, err := h.service.GetOriginalURL(r.Context(), code)
	//обрабатываем ошибки
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			sendError(w, http.StatusNotFound, "ссылка не найдена")
			return
		}
		if errors.Is(err, service.ErrLinkExpired) {
			sendError(w, http.StatusGone, "ссылка протухла")
			return
		}
		sendError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	// редирект из библиотеки
	http.Redirect(w, r, originalURL, http.StatusFound)
}

// роуты
func (h *Handler) RegRouters(mux *http.ServeMux) {
	mux.HandleFunc("POST /links", h.CreateLink)
	mux.HandleFunc("GET /{code}", h.Redirect)
}

//Хендлер — это то, что обрабатывает входящий HTTP-запрос и пишет ответ.
//w http.ResponseWriter — сюда пишется ответ (заголовки, статус, тело).
//r *http.Request — сюда приходит вся информация о запросе (метод, URL, заголовки, тело, контекст).
