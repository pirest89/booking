package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string         `gorm:"type:varchar(250);not null"`
	Address     sql.NullString `gorm:"type:varchar(250)"`
	License     sql.NullString `gorm:"type:varchar(20)"`
	PhoneNumber sql.NullString `gorm:"type:varchar(20)"`
	Email       sql.NullString `gorm:"type:varchar(200)"`
	Password    sql.NullString `gorm:"type:varchar(200)"`
	Active      bool           `gorm:"type:boolean"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
