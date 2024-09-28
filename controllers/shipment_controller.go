package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/yashbalyan08/shipment-app/models"
	"github.com/yashbalyan08/shipment-app/utils"
)

// CreateShipment handler
func CreateShipment(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var shipment models.Shipment
	err := json.NewDecoder(r.Body).Decode(&shipment)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = shipment.CreateShipment(db)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create shipment")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, shipment)
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
