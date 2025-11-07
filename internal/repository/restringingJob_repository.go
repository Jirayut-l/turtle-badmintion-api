package repository

import (
	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
	"gorm.io/gorm"
)

// RestringingJobRepository
type RestringingJobRepository interface {
	// Create ต้องรับ gorm.DB (tx) เพื่อให้ทำงานใน Transaction ได้
	Create(tx *gorm.DB, job *domain.RestringingJob) error
	ListByCustomerID(customerID uint) ([]domain.RestringingJob, error)
}
