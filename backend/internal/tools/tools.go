package tools

import (
	"math/rand"
	"time"
)

const apiKeyLength = 16 

func GenerateApiKey() string {

	rand.NewSource(time.Now().UnixNano())

	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	key := make([]byte, apiKeyLength)

	for i := range key {
		key[i] = characters[rand.Intn(len(characters))]
	}

	return string(key)
}
