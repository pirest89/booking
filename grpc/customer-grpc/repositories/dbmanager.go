package repositories

import (
	"bookings-mock/database"
	"bookings-mock/grpc/customer-grpc/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepositories interface {
	CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	BookingHistory(ctx context.Context, id uuid.UUID) (*models.Customer, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.Customer, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Customer{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) FindByID(ctx context.Context, id uuid.UUID) (*models.Customer, error) {
	model := &models.Customer{}

	if err := m.First(model, id).Error; err != nil {
		return nil, err
	}

	return model, nil

}

func (m *dbmanager) UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Updates(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) BookingHistory(ctx context.Context, id uuid.UUID) (*models.Customer, error) {
	return nil, nil
}
