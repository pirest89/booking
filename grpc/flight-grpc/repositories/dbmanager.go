package repositories

import (
	"bookings-mock/database"
	"bookings-mock/grpc/flight-grpc/models"
	"bookings-mock/grpc/flight-grpc/requests"
	"context"

	"gorm.io/gorm"
)

type FlightRepositories interface {
	CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	SearchFlight(ctx context.Context, req *requests.SearchFlightRequest) ([]*models.Flight, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Flight{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	flight := models.Flight{}
	if err := m.Updates(model).Error; err != nil {
		return nil, err
	}
	return &flight, nil
}

func (m *dbmanager) SearchFlight(ctx context.Context, req *requests.SearchFlightRequest) ([]*models.Flight, error) {
	flights := []*models.Flight{}
	if req.Level > 0 {
		if err := m.Where("level > ?", req.Level).Find(&flights).Error; err != nil {
			return nil, err
		}
	} else {
		if err := m.Find(&flights).Error; err != nil {
			return nil, err
		}
	}

	return flights, nil
}
