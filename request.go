package gortifacts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type getRequestParams struct {
	token string
}

func getRequest[T interface{}](endpoint string, params getRequestParams) (T, error) {
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

func headerAuthorization(token string) (string, string) {
	return "Authorization", fmt.Sprintf("Bearer %s", token)
}

func headerAccept() (string, string) {
	return "Accept", "application/json"
}
