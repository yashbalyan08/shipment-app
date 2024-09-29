package models

import (
	"database/sql"
	"log"
	"time"
)

type Shipment struct {
	ID                  string `json:"int"`
	ShipmentID          string `json:"shipment_id"`
	ClientName          string `json:"client_name"`
	PickupLocation      string `json:"pickup_location"`
	DeliveryLocation    string `json:"delivery_location"`
	CargoType           string `json:"cargo_type"`
	CargoWeight         int    `json:"cargo_weight"`
	SpecialInstructions string `json:"special_instructions"`
}

// Insert a new shipment into the database
func (s *Shipment) CreateShipment(db *sql.DB) (*Shipment, error) {
	query := `
        INSERT INTO shipments (shipment_id, client, pickup_location, delivery_location, cargo_type, cargo_weight, special_instructions, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	// Add both `created_at` and `updated_at` as timestamps
	err := db.QueryRow(query,
		s.ID,
		s.ShipmentID,
		s.ClientName,
		s.PickupLocation,
		s.DeliveryLocation,
		s.CargoType,
		s.CargoWeight,
		s.SpecialInstructions,
		time.Now(), // created_at
		time.Now(), // updated_at
	).Scan() // Scan the returned id into the struct's ID field

	if err != nil {
		log.Printf("Failed to insert shipment: %v", err)
		return nil, err
	}

	return s, nil
}

// Fetch all shipments from the database
func GetAllShipments(db *sql.DB) ([]Shipment, error) {
	query := `SELECT id, shipment_id, client, pickup_location, delivery_location, cargo_type, cargo_weight, special_instructions, created_at, updated_at FROM shipments`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to fetch shipments: %v", err)
		return nil, err
	}
	defer rows.Close()

	shipments := []Shipment{}
	for rows.Next() {
		var s Shipment
		err = rows.Scan(&s.ShipmentID, &s.ClientName, &s.PickupLocation, &s.DeliveryLocation, &s.CargoType, &s.CargoWeight, &s.SpecialInstructions)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		shipments = append(shipments, s)
	}

	return shipments, nil
}
