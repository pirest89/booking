package handlers

import (
	"bookings-mock/grpc/flight-grpc/models"
	"bookings-mock/grpc/flight-grpc/repositories"
	"bookings-mock/grpc/flight-grpc/requests"
	"bookings-mock/pb"
	"context"
	"database/sql"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightHandler struct {
	pb.UnimplementedFlightServiceServer
	flightRepository repositories.FlightRepositories
}

func NewFlightHandler(flightRepositories repositories.FlightRepositories) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepository: flightRepositories,
	}, nil
}

func (h *FlightHandler) Create(ctx context.Context, req *pb.Flight) (*pb.Flight, error) {
	flight := &models.Flight{}
	err := copier.Copy(&flight, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.flightRepository.CreateFlight(ctx, flight)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Flight{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *FlightHandler) Update(ctx context.Context, req *pb.Flight) (*pb.Flight, error) {
	flight := &models.Flight{}
	err := copier.Copy(&flight, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.flightRepository.UpdateFlight(ctx, flight)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Flight{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *FlightHandler) List(ctx context.Context, req *pb.ListFlightRequest) (*pb.ListFlightResponse, error) {
	listFlight := &requests.SearchFlightRequest{}

	err := copier.Copy(&listFlight, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.flightRepository.SearchFlight(ctx, listFlight)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight not found")
		}
		return nil, err
	}

	pRes := &pb.ListFlightResponse{}
	err = copier.Copy(&pRes.Flights, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}
