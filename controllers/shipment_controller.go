package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yashbalyan08/shipment-app/models"
	"github.com/yashbalyan08/shipment-app/utils"
)

// CreateShipment handler
func CreateShipment(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var shipment models.Shipment
	err := json.NewDecoder(r.Body).Decode(&shipment)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error decoding JSON: %v", err) // Add this log for more detail
		return
	}

	query := `INSERT INTO shipments (id, shipment_id, client, pickup_location, delivery_location, cargo_type, cargo_weight, special_instructions, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err = db.Exec(query, uuid.New(), shipment.ShipmentID, shipment.ClientName, shipment.PickupLocation, shipment.DeliveryLocation, shipment.CargoType, shipment.CargoWeight, shipment.SpecialInstructions, time.Now(), time.Now())
	if err != nil {
		http.Error(w, "Failed to create shipment", http.StatusInternalServerError)
		log.Printf("Error inserting into database: %v", err) // Log the actual error here
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Shipment created successfully"})
}

// GetAllShipments handler
func GetAllShipments(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	shipments, err := models.GetAllShipments(db)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve shipments")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, shipments)
}
