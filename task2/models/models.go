package models

import (
	"database/sql"
)

type Spot struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Website     sql.NullString `json:"website"`
	Coordinates string         `json:"coordinates"`
	Description sql.NullString `json:"description"`
	Rating      float64        `json:"rating"`
	Distance    float64        `json:"distance,omitempty"`
}
