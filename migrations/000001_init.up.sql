-- табла для коротких ссылок
CREATE TABLE IF NOT EXISTS links (
    id           BIGSERIAL PRIMARY KEY,
    code         VARCHAR(10) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at   TIMESTAMPTZ,
    clicks       BIGINT NOT NULL DEFAULT 0
);

-- индекс
CREATE INDEX IF NOT EXISTS idx_links_code ON links(code);

-- табла для аналитики
CREATE TABLE IF NOT EXISTS link_clicks (
    id         BIGSERIAL PRIMARY KEY,
    link_id    BIGINT NOT NULL REFERENCES links(id) ON DELETE CASCADE,
    clicked_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- индекс для быстрой группировки кликов по ссылке
CREATE INDEX IF NOT EXISTS idx_link_clicks_link_id ON link_clicks(link_id);