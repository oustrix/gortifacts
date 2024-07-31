package gortifacts

import (
	"fmt"

	"github.com/oustrix/gortifacts/endpoints"
)

// App is an artifacts controller.
type App struct {
	token string
}

// NewApp creates an App instance.
func NewApp(token string) App {
	return App{
		token: token,
	}
}

type GetStatusResponse struct {
	Data struct {
		Status string `json:"status"`
	} `json:"data"`
}

func (a *App) GetStatus() (*GetStatusResponse, error) {
	res, err := GetRequestWithoutToken[GetStatusResponse](endpoints.Status())
	if err != nil {
		return nil, fmt.Errorf("failed to get server status: %w", err)
	}

	return res, nil
}
