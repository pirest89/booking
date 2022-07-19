package responses

import "time"

type FlightResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Status        string    `json:"status"`
	Date          time.Time `json:"date"`
	AvailableSlot int64     `json:"availableSlot" binding:"required"`
}
