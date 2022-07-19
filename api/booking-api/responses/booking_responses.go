package responses

import "time"

type BookingResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	FlightID   string `json:"flight_id"`
	Code       string `json:"code"`
	Status     string `json:"status"`
	BookedDate time.Time
}
