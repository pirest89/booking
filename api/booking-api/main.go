package main

import (
	"bookings-mock/api/booking-api/handlers"
	"bookings-mock/middleware"
	"bookings-mock/pb"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	bookingConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	bookingServiceClient := pb.NewBookingServiceClient(bookingConn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	h := handlers.NewBookingHandler(bookingServiceClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/bookings", h.CreateBooking)
	gr.GET("/bookings/:id", h.ViewBooking)
	gr.POST("/bookings/cancel", h.CancelBooking)

	//Listen and serve
	http.ListenAndServe(":3333", g)
}
