package main

import (
	"bookings-mock/api/flight-api/handlers"
	"bookings-mock/middleware"
	"bookings-mock/pb"
	custom_validator "bookings-mock/validator"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	flightConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	flightClient := pb.NewFlightServiceClient(flightConn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	h := handlers.NewFlightHandler(flightClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("max-slot", custom_validator.FlightSlotValidator(int64(100)))
	}

	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/flights", h.CreateFlight)
	gr.PUT("/flights/:id", h.UpdateFlight)
	gr.GET("/flights", h.ListFlight)

	//Listen and serve
	http.ListenAndServe(":3335", g)
}
