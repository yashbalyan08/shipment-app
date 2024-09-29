package models

type Truck struct {
	TruckId      string `json:"truckid,omitempty"`
	TruckName    string `json:"truckname,omitempty"`
	TruckLicense string `json:"trucklicense,omitempty"`
	Availabilty  bool   `json:"availabilty,omitempty"`
}
