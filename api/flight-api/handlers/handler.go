package handlers

import (
	"bookings-mock/api/flight-api/requests"
	"bookings-mock/api/flight-api/responses"
	"bookings-mock/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FlightHandler interface {
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	ListFlight(c *gin.Context)
}

type flightHandler struct {
	flightClient pb.FlightServiceClient
}

func NewFlightHandler(flightClient pb.FlightServiceClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}

func (h *flightHandler) UpdateFlight(c *gin.Context) {
	req := requests.UpdateFlightRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}
	pReq := &pb.Flight{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		AvailableSlot: req.AvailableSlot,
	}
	pRes, err := h.flightClient.Update(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		ID:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		AvailableSlot: pRes.AvailableSlot,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *flightHandler) ListFlight(c *gin.Context) {
	req := requests.ListFlightRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}
	pReq := &pb.ListFlightRequest{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		AvailableSlot: req.AvailableSlot,
	}
	pRes, err := h.flightClient.List(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dtos := make([]*responses.FlightResponse, 0)

	for _, v := range pRes.Flights {
		dto := &responses.FlightResponse{
			ID:            v.Id,
			Name:          v.Name,
			From:          v.From,
			To:            v.To,
			AvailableSlot: v.AvailableSlot,
		}

		dtos = append(dtos, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dtos,
	})
}

func (h *flightHandler) CreateFlight(c *gin.Context) {
	req := requests.CreateFlightRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}
	pReq := &pb.Flight{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		AvailableSlot: req.AvailableSlot,
	}
	pRes, err := h.flightClient.Create(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		ID:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		AvailableSlot: pRes.AvailableSlot,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
