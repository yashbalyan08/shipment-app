package models

type Driver struct {
	DriverId      string `json:"driverid,omitempty" `
	DriverName    string `json:"drivername,omitempty"`
	DriverLicense string `json:"driverlicense,omitempty"`
	Availabilty   bool   `json:"availabilty,omitempty"`
}
