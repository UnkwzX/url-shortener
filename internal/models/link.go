package models

// для временных меток
import "time"

// Link -> запись короткой ссылки в БД
type Link struct {
	ID          int64      `json:"id"`
	Code        string     `json:"code"`
	OriginalURL string     `json:"original_url"`
	CreatedAt   time.Time  `json:"created_at"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	// ExpiresAt - исп. указатель, т.к. указатель передаст nil, если в БД будет NULL. Иначе pgx выпадет с ошибкой.
	Clicks int64 `json:"clicks"`
}

// ClickStat -> стата по дням
type ClickStat struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

type LinkStats struct {
	Code        string      `json:"code"`
	OriginalURL string      `json:"original_url"`
	TotalClicks int64       `json:"total_clicks"`
	CreatedAt   time.Time   `json:"created_at"`
	ByDay       []ClickStat `json:"by_day"`
}

//Тут описываются доменные сущности, которые будут использоваться в проекте не 1 раз.
