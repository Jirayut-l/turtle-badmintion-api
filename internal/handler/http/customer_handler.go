package http

import (
	"net/http"

	"github.com/Jirayut-l/turtle-badminton-api/internal/handler/http/dto"
	"github.com/Jirayut-l/turtle-badminton-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

// CustomerHandler
type CustomerHandler struct {
	uc *usecase.CustomerUsecase
}

func NewCustomerHandler(uc *usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{uc: uc}
}

func (h *CustomerHandler) RegisterCustomer(c *gin.Context) {
	var req dto.CreateCustomerRequest // <-- ใช้ DTO จาก package ใหม่
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.uc.RegisterCustomer(req.Name, req.Phone, req.Notes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ใช้ DTO Mapper: แปลง Entity -> DTO ก่อนส่ง
	c.JSON(http.StatusCreated, dto.NewCustomerResponse(customer)) // <-- ใช้ Mapper จาก DTO
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone query parameter is required"})
		return
	}

	customer, err := h.uc.GetCustomerByPhone(phone)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}

	// ใช้ DTO Mapper: แปลง Entity -> DTO ก่อนส่ง
	c.JSON(http.StatusOK, dto.NewCustomerResponse(customer)) // <-- ใช้ Mapper จาก DTO
}
