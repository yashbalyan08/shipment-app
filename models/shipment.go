package models

import (
	"database/sql"
	"log"
	"time"
)

type Shipment struct {
	ID                  int64     `json:"id"`
	ShipmentID          string    `json:"shipment_id"`
	Client              string    `json:"client"`
	PickupLocation      string    `json:"pickup_location"`
	DeliveryLocation    string    `json:"delivery_location"`
	CargoType           string    `json:"cargo_type"`
	CargoWeight         float64   `json:"cargo_weight"`
	SpecialInstructions string    `json:"special_instructions"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// Insert a new shipment into the database
func (s *Shipment) CreateShipment(db *sql.DB) error {
	query := `
        INSERT INTO shipments (shipment_id, client, pickup_location, delivery_location, cargo_type, cargo_weight, special_instructions, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id`

	err := db.QueryRow(query, s.ShipmentID, s.Client, s.PickupLocation, s.DeliveryLocation, s.CargoType, s.CargoWeight, s.SpecialInstructions, time.Now(), time.Now()).Scan(&s.ID)
	if err != nil {
		log.Printf("Failed to insert shipment: %v", err)
		return err
	}

	return nil
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
		err = rows.Scan(&s.ID, &s.ShipmentID, &s.Client, &s.PickupLocation, &s.DeliveryLocation, &s.CargoType, &s.CargoWeight, &s.SpecialInstructions, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		shipments = append(shipments, s)
	}

	return shipments, nil
}
