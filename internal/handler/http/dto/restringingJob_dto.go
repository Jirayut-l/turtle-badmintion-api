package dto

import (
	"time"

	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
)

// --- Request DTOs ---

// LogJobRequest คือ Request DTO
// ใช้สำหรับ Bind JSON ตอนบันทึกงานขึ้นเอ็น
type LogJobRequest struct {
	CustomerID   uint    `json:"customer_id" binding:"required"`
	ProductID    uint    `json:"product_id" binding:"required"`
	Tension      string  `json:"tension" binding:"required"`
	PriceCharged float64 `json:"price_charged"`
	Notes        string  `json:"notes"`
}

// --- Response DTOs ---

// RestringingJobResponse
type RestringingJobResponse struct {
	ID           uint             `json:"id"`
	Tension      string           `json:"tension"`
	PriceCharged float64          `json:"price_charged"`
	JobDate      time.Time        `json:"job_date"`
	Notes        string           `json:"notes"`
	Customer     CustomerResponse `json:"customer"` // <-- เราใช้ DTO ซ้อน DTO
	Product      ProductResponse  `json:"product"`  // <-- เราใช้ DTO ซ้อน DTO
}

// NewRestringingJobResponse (Mapper)
func NewRestringingJobResponse(j *domain.RestringingJob) RestringingJobResponse {
	return RestringingJobResponse{
		ID:           j.ID,
		Tension:      j.Tension,
		PriceCharged: j.PriceCharged,
		JobDate:      j.JobDate,
		Notes:        j.Notes,
		Customer:     NewCustomerResponse(&j.Customer), // แปลง nested struct ด้วย
		Product:      NewProductResponse(&j.Product),   // แปลง nested struct ด้วย
	}
}
