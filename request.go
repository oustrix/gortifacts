package gortifacts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRequest[T interface{}](endpoint string, token string) (*T, error) {
	var payload = new(T)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return payload, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Add(headerAuthorization(token))
	req.Header.Add(headerAccept())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.DefaultClient.Do: %w", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return payload, fmt.Errorf("json.Decode: %w", err)
	}

	return payload, nil
}

func GetRequestWithoutToken[T interface{}](endpoint string) (*T, error) {
	var payload = new(T)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return payload, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Add(headerAccept())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.DefaultClient.Do: %w", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return payload, fmt.Errorf("json.Decode: %w", err)
	}

	return payload, nil
}

func headerAuthorization(token string) (string, string) {
	return "Authorization", fmt.Sprintf("Bearer %s", token)
}

func headerAccept() (string, string) {
	return "Accept", "application/json"
}
