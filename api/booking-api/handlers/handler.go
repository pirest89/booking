package handlers

import (
	"bookings-mock/api/booking-api/requests"
	"bookings-mock/api/booking-api/responses"
	"bookings-mock/pb"
	"github.com/jinzhu/copier"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookingHandler interface {
	CreateBooking(c *gin.Context)
	ViewBooking(c *gin.Context)
	CancelBooking(c *gin.Context)
}

type bookingHandler struct {
	bookingClient pb.BookingServiceClient
}

func (h *bookingHandler) ViewBooking(c *gin.Context) {
	id := c.Param("id")
	pReq := &pb.ViewBookingRequest{
		BookingId: id,
	}

	pRes, err := h.bookingClient.ViewBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingResponse{}
	copier.Copy(dto, pRes)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *bookingHandler) CancelBooking(c *gin.Context) {
	id := c.Param("id")
	pReq := &pb.CancelBookingRequest{
		BookingId: id,
	}

	pRes, err := h.bookingClient.CancelBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingResponse{}
	copier.Copy(dto, pRes)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *bookingHandler) CreateBooking(c *gin.Context) {
	req := requests.CreateBookingRequest{}

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

	pReq := &pb.BookingRequest{}

	pRes, err := h.bookingClient.CreateBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingResponse{}
	copier.Copy(dto, pRes)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func NewBookingHandler(bookingServiceClient pb.BookingServiceClient) BookingHandler {
	return &bookingHandler{
		bookingClient: bookingServiceClient,
	}
}
