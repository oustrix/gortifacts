package gortifacts

import (
	"fmt"
	"net/http"
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
		Status           string `json:"status"`            // Server status.
		Version          string `json:"version"`           // Game version.
		CharactersOnline int    `json:"characters_online"` // Current characters online number.
		Announcements    []struct {
			Message   string    `json:"message"`    // Announcement text.
			CreatedAt time.Time `json:"created_at"` // Datetime of the announcement.
		} `json:"announcements"`
		LastWipe string `json:"last_wipe"` // Last server wipe.
		NextWipe string `json:"next_wipe"` // Next server wipe.
	} `json:"data"`
}

// GetStatus return the status of the game server.
func (a *App) GetStatus() (GetStatusResponse, error) {
	res, err := get[GetStatusResponse](endpoints.Status(), getParams{})
	if err != nil {
		return GetStatusResponse{}, fmt.Errorf("failed to get server status: %w", err)
	}

	return res, nil
}

// ActionMoveRequest represents data for ActionMove request.
type ActionMoveRequest struct {
	X int `json:"x"` // The x coordinate of the destination.
	Y int `json:"y"` // The x coordinate of the destination.
}

// ActionMoveResponse represents GetStatus API response.
type ActionMoveResponse struct {
	Data struct {
		Cooldown    Cooldown `json:"cooldown"`
		Destination struct {
			Name    string `json:"name"` // The name of the destination.
			X       int    `json:"x"`    // The x coordinate of the destination.
			Y       int    `json:"y"`    // The y coordinate of the destination.
			Content *struct {
				Type string `json:"type"` // The type of the map content.
				Code string `json:"code"` // The code of the map content.
			} `json:"content"` // Content of the destination.
		} `json:"destination"` // Destination details.
		Character Character `json:"character"` // Character details.
	} `json:"data"`
}

// ActionMove moves a character on the map using the map's X and Y position.
func (a *App) ActionMove(characterName string, req ActionMoveRequest) (ActionMoveResponse, error) {
	res, code, err := post[ActionMoveResponse](
		endpoints.ActionMove(characterName),
		postParams{
			token: a.token,
			body:  req,
		},
	)

	switch code {
	case http.StatusOK:
		// Ok
	case http.StatusUnprocessableEntity, http.StatusBadRequest:
		return ActionMoveResponse{}, ErrorInvalidInputData
	case 486:
		return ActionMoveResponse{}, ErrorCharacterIsLocked
	case 490:
		return ActionMoveResponse{}, ErrorCharacterAtDestination
	case 498:
		return ActionMoveResponse{}, ErrorCharacterNotFound
	case 499:
		return ActionMoveResponse{}, ErrorCharacterCooldown
	case 500:
		return ActionMoveResponse{}, fmt.Errorf("internal error: %w", err)
	default:
		return ActionMoveResponse{}, fmt.Errorf("unexpected code: %d, error: %w", code, err)
	}
	if err != nil {
		return ActionMoveResponse{}, fmt.Errorf("unexpected error while move: %w", err)
	}

	return res, nil
}
