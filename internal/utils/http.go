package utils

import (
	"encoding/json"
	"net/http"
)

func GetData[T any](endpoint string, result T) (T, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return result, err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
