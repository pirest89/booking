package handlers

import (
	"bookings-mock/grpc/booking-grpc/models"
	"bookings-mock/grpc/booking-grpc/repositories"
	"bookings-mock/pb"
	"context"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
	bookingRepositories repositories.BookingRepositories
}

func NewBookingHandler(jobRepository repositories.BookingRepositories) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepositories: jobRepository,
	}, nil
}

func (h *BookingHandler) CreateBooking(ctx context.Context, req *pb.BookingRequest) (*pb.Booking, error) {
	booking := &models.Booking{}
	err := copier.Copy(&booking, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.bookingRepositories.CreateBooking(ctx, booking)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Booking{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
func (h *BookingHandler) ViewBooking(ctx context.Context, req *pb.ViewBookingRequest) (*pb.Booking, error) {
	booking := &models.Booking{}
	err := copier.Copy(&booking, &req)
	if err != nil {
		return nil, err
	}

	id, _ := uuid.FromBytes([]byte(req.BookingId))
	res, err := h.bookingRepositories.ViewBooking(ctx, id)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Booking{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
func (h *BookingHandler) CancelBooking(ctx context.Context, req *pb.CancelBookingRequest) (*pb.Empty, error) {
	booking := &models.Booking{}
	err := copier.Copy(&booking, &req)
	if err != nil {
		return nil, err
	}

	id, _ := uuid.FromBytes([]byte(req.BookingId))
	_, err = h.bookingRepositories.CancelBooking(ctx, id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
