package routes

import (
	"database/sql"
	"net/http"

	"github.com/yashbalyan08/shipment-app/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shipments", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateShipment(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/shipments", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllShipments(db, w, r)
	}).Methods("GET")

	return router
}
