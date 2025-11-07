package http

import (
	"net/http"

	"github.com/Jirayut-l/turtle-badminton-api/internal/handler/http/dto"
	"github.com/Jirayut-l/turtle-badminton-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

// RestringingHandler
type RestringingHandler struct {
	uc *usecase.RestringingUsecase
}

func NewRestringingHandler(uc *usecase.RestringingUsecase) *RestringingHandler {
	return &RestringingHandler{uc: uc}
}

func (h *RestringingHandler) LogNewJob(c *gin.Context) {
	var req dto.LogJobRequest // <-- ใช้ DTO จาก package ใหม่
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job, err := h.uc.LogNewJob(req.CustomerID, req.ProductID, req.Tension, req.PriceCharged, req.Notes)
	if err != nil {
		if err.Error() == "สินค้าหมดสต็อก (product out of stock)" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()}) // 409 Conflict
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// ใช้ DTO Mapper: แปลง Entity (with preloads) -> DTO
	c.JSON(http.StatusCreated, dto.NewRestringingJobResponse(job)) // <-- ใช้ Mapper จาก DTO
}
