package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CustomerID int64
	FlightID   int64
	Code       string `gorm:"type:varchar(250);not null"`
	Status     string `gorm:"type:varchar(250);not null"`
	BookedDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
