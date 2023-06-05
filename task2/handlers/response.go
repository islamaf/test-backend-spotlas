package spots

import (
	models "task2/models"
)

type SpotsResponse struct {
	Type    string        `json:"code"`
	Spots   []models.Spot `json:"spots"`
	Message string        `json:"message"`
}
