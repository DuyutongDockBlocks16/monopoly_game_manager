package service

import (
    "database/sql"
    "log"
    "app-backend/database"
)

type YourData struct {
    ID   int    `json:"id"`
	LandPrice  int    `json:"land_price"`
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
