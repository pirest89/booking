package models

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	Id            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name          string    `gorm:"type:varchar(250);not null"`
	From          string    `gorm:"type:varchar(10)"`
	To            string    `gorm:"type:varchar(10)"`
	Status        string    `gorm:"type:varchar(10)"`
	AvailableSlot int64     `gorm:"type:bigint"'`
	Date          time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
