package repositories

import (
	"bookings-mock/database"
	"bookings-mock/grpc/booking-grpc/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingRepositories interface {
	CreateBooking(ctx context.Context, model *models.Booking) (*models.Booking, error)
	ViewBooking(ctx context.Context, id uuid.UUID) (*models.Booking, error)
	CancelBooking(ctx context.Context, id uuid.UUID) (*models.Booking, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Booking{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) CreateBooking(ctx context.Context, model *models.Booking) (*models.Booking, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) ViewBooking(ctx context.Context, id uuid.UUID) (*models.Booking, error) {
	booking := models.Booking{}
	if err := m.First(&booking, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (m *dbmanager) CancelBooking(ctx context.Context, id uuid.UUID) (*models.Booking, error) {
	booking := models.Booking{}
	if err := m.First(&booking, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}
	booking.Status = "Cancelled"
	if err := m.Updates(booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}
