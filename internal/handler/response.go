package handler

import (
	"encoding/json"
	"net/http"
)

func sendJSON(w http.ResponseWriter, status int, data any) {
	//через .header задаем заголовок, того что шлем json
	w.Header().Set("Content-Type", "application/json")
	//через .WriteHeader отправляем статус код 200, 400...
	w.WriteHeader(status)

	//сереализация data
	json.NewEncoder(w).Encode(data)
}

func sendError(w http.ResponseWriter, status int, message string) {
	mError := map[string]string{"error": message}

	sendJSON(w, status, mError)
}
