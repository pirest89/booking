package requests

import (
	"github.com/google/uuid"
	"time"
)

type CreateFlightRequest struct {
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Status        string    `json:"status"`
	Date          time.Time `json:"date"`
	AvailableSlot int64     `json:"availableSlot" binding:"required"`
}

type UpdateFlightRequest struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Status        string    `json:"status"`
	Date          time.Time `json:"date"`
	AvailableSlot int64     `json:"availableSlot" binding:"required"`
}

type ListFlightRequest struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Status        string    `json:"status"`
	Date          time.Time `json:"date"`
	AvailableSlot int64     `json:"availableSlot" binding:"required"`
}
