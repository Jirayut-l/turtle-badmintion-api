package http

import (
	"net/http"

	"github.com/Jirayut-l/turtle-badminton-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

// Gin Handler ทำหน้าที่เป็น "ล่าม"
// แปล HTTP Request -> การเรียกใช้ Usecase
// แปลผลลัพธ์จาก Usecase -> HTTP Response (JSON)

// CustomerHandler
type CustomerHandler struct {
	uc *usecase.CustomerUsecase
}

func NewCustomerHandler(uc *usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{uc: uc}
}

type CreateCustomerRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Notes string `json:"notes"`
}

func (h *CustomerHandler) RegisterCustomer(c *gin.Context) {
	var req CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.uc.RegisterCustomer(req.Name, req.Phone, req.Notes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
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

	c.JSON(http.StatusOK, customer)
}

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
	c.JSON(http.StatusOK, products)
}

// ... (สามารถเพิ่ม Handler สำหรับ AddNewProduct ได้ในทำนองเดียวกัน) ...

// RestringingHandler
type RestringingHandler struct {
	uc *usecase.RestringingUsecase
}

func NewRestringingHandler(uc *usecase.RestringingUsecase) *RestringingHandler {
	return &RestringingHandler{uc: uc}
}

type LogJobRequest struct {
	CustomerID   uint    `json:"customer_id" binding:"required"`
	ProductID    uint    `json:"product_id" binding:"required"`
	Tension      string  `json:"tension" binding:"required"`
	PriceCharged float64 `json:"price_charged" binding:"required"`
	Notes        string  `json:"notes"`
}

func (h *RestringingHandler) LogNewJob(c *gin.Context) {
	var req LogJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job, err := h.uc.LogNewJob(req.CustomerID, req.ProductID, req.Tension, req.PriceCharged, req.Notes)
	if err != nil {
		if err.Error() == "สินค้าหมดสต็อก (product out of stock)" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, job)
}
