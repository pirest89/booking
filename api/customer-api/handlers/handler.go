package handlers

import (
	"bookings-mock/api/customer-api/requests"
	"bookings-mock/api/customer-api/responses"
	"bookings-mock/pb"
	"github.com/jinzhu/copier"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerHandler interface {
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	ChangePassword(c *gin.Context)
	BookingHistory(c *gin.Context)
}

type customerHandler struct {
	customerServiceClient pb.CustomerServiceClient
}

func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	req := requests.UpdateCustomerRequest{}

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

	pReq := &pb.Customer{}
	err := copier.Copy(&pReq, req)
	if err != nil {
	}
	pRes, err := h.customerServiceClient.Update(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{}
	err = copier.Copy(dto, pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *customerHandler) ChangePassword(c *gin.Context) {
	req := requests.UpdateCustomerRequest{}

	pReq := &pb.ChangePasswordRequest{}
	err := copier.Copy(&pReq, req)
	if err != nil {
	}
	pRes, err := h.customerServiceClient.ChangePassword(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{}
	err = copier.Copy(dto, pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *customerHandler) BookingHistory(c *gin.Context) {
	id := c.Param("id")
	pReq := &pb.BookingHistoryRequest{
		CustomerId: id,
	}
	pRes, err := h.customerServiceClient.BookingHistory(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{}
	err = copier.Copy(dto, pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	req := requests.CreateCustomerRequest{}

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

	pReq := &pb.Customer{}
	err := copier.Copy(&pReq, req)
	if err != nil {
	}
	pRes, err := h.customerServiceClient.Create(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{}
	err = copier.Copy(dto, pRes)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func NewPeopleHandler(customerServiceClient pb.CustomerServiceClient) CustomerHandler {
	return &customerHandler{
		customerServiceClient: customerServiceClient,
	}
}
