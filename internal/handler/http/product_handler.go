package http

import (
	"net/http"

	"github.com/Jirayut-l/turtle-badminton-api/internal/handler/http/dto"
	"github.com/Jirayut-l/turtle-badminton-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

// ProductHandler
type ProductHandler struct {
	uc *usecase.ProductUsecase
}

func NewProductHandler(uc *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.uc.ListAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ใช้ DTO Mapper: แปลง []Entity -> []DTO ก่อนส่ง
	c.JSON(http.StatusOK, dto.NewProductListResponse(products)) // <-- ใช้ Mapper จาก DTO
}

// CreateProduct เพิ่มสินค้าใหม่ (เพื่อเติมสต็อก)
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest // <-- ใช้ DTO จาก package ใหม่
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.uc.AddNewProduct(
		req.SKU,
		req.Brand,
		req.Model,
		req.QuantityInStock,
		req.CostPrice,
		req.SellingPrice,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ใช้ DTO Mapper: แปลง Entity -> DTO ก่อนส่ง
	c.JSON(http.StatusCreated, dto.NewProductResponse(product)) // <-- ใช้ Mapper จาก DTO
}
