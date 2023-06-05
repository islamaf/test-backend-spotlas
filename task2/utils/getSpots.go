package utils

import (
	"fmt"
	db "task2/core"
	models "task2/models"
)

func FindSpots(latitude, longitude, radius float64, locationType string) ([]models.Spot, error) {
	db := db.InitDb()

	query := ""
	initQuery := `
			SELECT
				id, name, website, ST_AsText(coordinates), description, rating
			FROM
				"MY_TABLE"
	`

	closeQuery := `
		ORDER BY
		CASE
			WHEN ST_Distance(coordinates::geography, 'SRID=4326;POINT(%f %f)') < 50
			THEN rating
			ELSE ST_Distance(coordinates::geography, 'SRID=4326;POINT(%f %f)')
		END;
	`

	switch locationType {
	case "circle":
		query = fmt.Sprintf(initQuery+`
			WHERE
				ST_DWithin(coordinates::geography, 'SRID=4326;POINT(%f %f)', %f)
		`+closeQuery, longitude, latitude, radius, longitude, latitude, longitude, latitude)
	case "square":
		query = fmt.Sprintf(initQuery+`
			WHERE
				ST_DWithin(coordinates::geography, ST_MakeEnvelope(%f, %f, %f, %f, 4326), 0)
		`+closeQuery, longitude, latitude, longitude-radius, latitude-radius, longitude+radius, latitude+radius, longitude, latitude)
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	spots := []models.Spot{}
	for rows.Next() {
		var spot models.Spot
		err := rows.Scan(&spot.ID, &spot.Name, &spot.Website, &spot.Coordinates, &spot.Description, &spot.Rating)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		spots = append(spots, spot)
	}

	return spots, nil
}
