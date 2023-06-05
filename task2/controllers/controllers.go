package controllers

import (
	"encoding/json"
	"net/http"
	handlers "task2/handlers"
	"task2/models"
	utils "task2/utils"
)

func GetSpotsController(w http.ResponseWriter, r *http.Request) {
	var response = handlers.SpotsResponse{}
	var spots []models.Spot

	latitude := utils.QueryParser(w, r, "latitude")
	longitude := utils.QueryParser(w, r, "longitude")
	radius := utils.QueryParser(w, r, "radius")
	if latitude == -1 || longitude == -1 || radius == -1 {
		w.WriteHeader(http.StatusBadRequest)
		response = handlers.SpotsResponse{Type: "400", Spots: spots, Message: "Incorrect params! Check latitude, longitude and radius. Make sure they are correct."}
		json.NewEncoder(w).Encode(response)
		return
	}

	locationType := r.URL.Query().Get("type")
	if locationType != "circle" && locationType != "square" {
		w.WriteHeader(http.StatusBadRequest)
		response = handlers.SpotsResponse{Type: "400", Spots: spots, Message: "Incorrect params! Valid types are 'circle' and 'square'."}
		json.NewEncoder(w).Encode(response)
		return
	}

	spots, err := utils.FindSpots(latitude, longitude, radius, locationType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = handlers.SpotsResponse{Type: "400", Spots: spots, Message: "Error getting spots!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	response = handlers.SpotsResponse{Spots: spots}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response = handlers.SpotsResponse{Type: "200", Spots: spots, Message: "Spots have been found!"}
	json.NewEncoder(w).Encode(response)
}
