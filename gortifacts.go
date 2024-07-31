package gortifacts

import (
	"fmt"
	"time"

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

// GetStatusResponse represents GetStatus API response.
type GetStatusResponse struct {
	Data struct {
		Status           string `json:"status"`
		Version          string `json:"version"`
		CharactersOnline int    `json:"characters_online"`
		Announcments     struct {
			Message   string    `json:"message"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"announcments"`
		LastWipe time.Time `json:"last_wipe"`
		NextWipe time.Time `json:"next_wipe"`
	} `json:"data"`
}

// GetStatus returns current service status.
func (a *App) GetStatus() (GetStatusResponse, error) {
	res, err := getRequest[GetStatusResponse](endpoints.Status(), getRequestParams{})
	if err != nil {
		return GetStatusResponse{}, fmt.Errorf("failed to get server status: %w", err)
	}

	return res, nil
}
