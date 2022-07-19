package main

import (
	"bookings-mock/grpc/flight-grpc/handlers"
	"bookings-mock/grpc/flight-grpc/repositories"
	"bookings-mock/helper"
	"bookings-mock/intercepter"
	"bookings-mock/pb"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := helper.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			intercepter.UnaryServerLoggingIntercepter(logger),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
	)

	flightRepositories, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewFlightHandler(flightRepositories)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterFlightServiceServer(s, h)

	logger.Info("Listen at port: 2225")

	s.Serve(listen)
}
