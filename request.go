package gortifacts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type getParams struct {
	token string
}

func get[T interface{}](endpoint string, params getParams) (T, error) {
	var payload T

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return payload, fmt.Errorf("http.NewRequest: %w", err)
	}
	req.Header.Add(headerAccept())

	if params.token != "" {
		req.Header.Add(headerAuthorization(params.token))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return payload, fmt.Errorf("http.DefaultClient.Do: %w", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return payload, fmt.Errorf("json.Decode: %w", err)
	}

	return payload, nil
}

type postParams struct {
	token string
	body  interface{}
}

func post[T interface{}](endpoint string, params postParams) (T, int, error) {
	var payload T

	jsonBody, err := json.Marshal(params.body)
	if err != nil {
		return payload, 400, fmt.Errorf("json.Marshal: %w", err)
	}

	body := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		return payload, 500, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Add(headerAuthorization(params.token))
	req.Header.Add(headerAccept())
	req.Header.Add(headerContentType())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if res == nil {
			return payload, 500, fmt.Errorf("http.DefaultClient.Do: %w", err)
		} else {
			return payload, res.StatusCode, fmt.Errorf("http.DefaultClient.Do: %w", err)
		}

	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return payload, 500, fmt.Errorf("json.Decode: %w", err)
	}

	return payload, res.StatusCode, nil
}

func headerAuthorization(token string) (string, string) {
	return "Authorization", fmt.Sprintf("Bearer %s", token)
}

func headerAccept() (string, string) {
	return "Accept", "application/json"
}

func headerContentType() (string, string) {
	return "Content-Type", "application/json"
}
