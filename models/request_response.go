package models

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HubRequest struct {
	Name        string `validate:"empty=false" json:"name"`
	GeoLocation string `validate:"empty=false" json:"geo_location"`
}

type TeamRequest struct {
	Name     string `validate:"empty=false" json:"name"`
	TeamType string `validate:"empty=false" json:"team_type"`
	HubID    int64  `json:"hub_id"`
}

type UserRequest struct {
	Role   string `validate:"empty=false"  json:"role"`
	Email  string `validate:"empty=false & format=email" json:"email"`
	TeamID int64  `json:"team_id"`
}

type HubAssign struct {
	TeamID int64 `validate:"nil=false & gte=0" json:"team_id"`
	HubID  int64 `validate:"nil=false & gte=0" json:"hub_id"`
}

type TeamAssign struct {
	UserID int64 `validate:"nil=false & gte=0" json:"user_id"`
	TeamID int64 `validate:"nil=false & gte=0" json:"team_id"`
}
