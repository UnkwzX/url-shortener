//Вся логика приложения живет тут.
//Валидация, генерация самого короткого кода, срок жизни ссылки.

package service

import (
	"crypto/rand"
	"math/big"
)

// алфавит для генерации кода и длина кода
const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const lenght = 8

func generateCode() (code string, err error) {
	result := make([]byte, lenght)
	// максмимально число для rand.Int
	max := big.NewInt(int64(len(charset)))

	for i, _ := range result {
		//выбирает случайное число в диапазоне 0-max
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		//заполняет i-й случайным с charset
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}
