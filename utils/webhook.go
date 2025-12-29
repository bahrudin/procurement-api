package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func SendWebhook(payload interface{}) error {

	url := os.Getenv("WEBHOOK_URL")
	if url == "" {
		return nil // webhook optional
	}

	body, _ := json.Marshal(payload)

	_, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	return err
}
