package service

import (
	"app-backend/database"
	"database/sql"
	"log"
)

type YourData struct {
	ID          int    `json:"id"`
	LandPrice   int    `json:"land_price"`
	CountryName string `json:"country_name"`
}

type SampleService struct {
	db *sql.DB
}

func NewPropertyService() *SampleService {
	return &SampleService{
		db: database.GetPool(),
	}
}

func (s *SampleService) GetAllProperties() ([]YourData, error) {
	query := "SELECT id, country_name, land_price FROM properties"

	rows, err := s.db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var results []YourData

	for rows.Next() {
		var data YourData
		if err := rows.Scan(&data.ID, &data.CountryName, &data.LandPrice); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in rows:", err)
		return nil, err
	}

	return results, nil
}

func (s *SampleService) OwnedPropertiesHandler(gameID string) ([]YourData, error) {
	query := `
        SELECT 
            a.id, 
            a.country_name, 
            a.land_price,
            b.house_number,
            b.property_status
        FROM properties a
        LEFT JON 
        (
            SELECT
                house_number,
                property_status
            FROM properties_in_game
            WHERE game_id = $1
            
        ) b 
        ON a.id = b.property_id
    `

	rows, err := s.db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var results []YourData

	for rows.Next() {
		var data YourData
		if err := rows.Scan(&data.ID, &data.CountryName, &data.LandPrice); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		results = append(results, data)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error in rows:", err)
		return nil, err
	}

	return results, nil
}
