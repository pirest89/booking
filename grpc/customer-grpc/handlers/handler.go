package handlers

import (
	"bookings-mock/grpc/customer-grpc/models"
	"bookings-mock/grpc/customer-grpc/repositories"
	"bookings-mock/pb"
	"context"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
	customerRepositories repositories.CustomerRepositories
}

func NewCustomerHandler(customerRepositories repositories.CustomerRepositories) (*CustomerHandler, error) {
	return &CustomerHandler{
		customerRepositories: customerRepositories,
	}, nil
}

func (h *CustomerHandler) Create(ctx context.Context, req *pb.Customer) (*pb.Customer, error) {
	customer := &models.Customer{}
	err := copier.Copy(&customer, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.customerRepositories.CreateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Customer{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
func (h *CustomerHandler) Update(ctx context.Context, req *pb.Customer) (*pb.Customer, error) {
	customer := &models.Customer{}
	err := copier.Copy(&customer, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.customerRepositories.UpdateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Customer{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
func (h *CustomerHandler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.Empty, error) {
	customer := &models.Customer{}
	err := copier.Copy(&customer, &req)
	if err != nil {
		return nil, err
	}

	id, _ := uuid.FromBytes([]byte(req.Id))
	res, err := h.customerRepositories.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Customer{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (h *CustomerHandler) BookingHistory(context.Context, *pb.BookingHistoryRequest) (*pb.BookingHistoryResponse, error) {

	return nil, nil
}
