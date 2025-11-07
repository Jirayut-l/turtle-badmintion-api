package dto

import "github.com/Jirayut-l/turtle-badminton-api/internal/domain"

// --- Request DTOs ---

// CreateProductRequest คือ Request DTO
// ใช้สำหรับ Bind JSON ที่ส่งเข้ามาตอนสร้างสินค้า
type CreateProductRequest struct {
	SKU             string  `json:"sku" binding:"required"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	QuantityInStock int     `json:"quantity_in_stock" binding:"required"`
	CostPrice       float64 `json:"cost_price"`
	SellingPrice    float64 `json:"selling_price"`
}

// --- Response DTOs ---

// ProductResponse
type ProductResponse struct {
	ID              uint    `json:"id"`
	SKU             string  `json:"sku"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	QuantityInStock int     `json:"quantity_in_stock"`
	SellingPrice    float64 `json:"selling_price"` // <-- สังเกตว่าเราไม่ส่ง CostPrice (ราคาทุน) ออกไป
}

// NewProductResponse (Mapper)
func NewProductResponse(p *domain.Product) ProductResponse {
	return ProductResponse{
		ID:              p.ID,
		SKU:             p.SKU,
		Brand:           p.Brand,
		Model:           p.Model,
		QuantityInStock: p.QuantityInStock,
		SellingPrice:    p.SellingPrice,
	}
}

// NewProductListResponse (Mapper for lists)
func NewProductListResponse(products []domain.Product) []ProductResponse {
	responses := make([]ProductResponse, len(products))
	for i, p := range products {
		responses[i] = NewProductResponse(&p)
	}
	return responses
}
