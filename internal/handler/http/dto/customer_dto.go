package dto

import (
	"time"

	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
)

// --- Request DTOs ---

// CreateCustomerRequest คือ Request DTO
// ใช้สำหรับ Bind JSON ที่ส่งเข้ามาตอนสร้างลูกค้า
type CreateCustomerRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Notes string `json:"notes"`
}

// --- Response DTOs ---

// CustomerResponse คือ struct ที่เราจะส่งกลับไปใน API
type CustomerResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}

// NewCustomerResponse เป็น "Mapper"
// แปลง domain.Customer (Entity) -> CustomerResponse (DTO)
func NewCustomerResponse(c *domain.Customer) CustomerResponse {
	return CustomerResponse{
		ID:        c.ID,
		Name:      c.Name,
		Phone:     c.Phone,
		Notes:     c.Notes,
		CreatedAt: c.CreatedAt,
	}
}
