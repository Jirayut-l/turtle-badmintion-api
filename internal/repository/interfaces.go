package repository

import (
	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
	"gorm.io/gorm"
)

// นี่คือ "สัญญา" (Contract) ที่ชั้น Usecase จะเรียกใช้
// โดยไม่สนว่าข้างหลังบ้านจะใช้ GORM, SQLC หรืออะไรก็ตาม

// CustomerRepository
type CustomerRepository interface {
	Create(customer *domain.Customer) error
	FindByID(id uint) (*domain.Customer, error)
	FindByPhone(phone string) (*domain.Customer, error)
	Update(customer *domain.Customer) error
}

// ProductRepository
type ProductRepository interface {
	Create(product *domain.Product) error
	FindByID(id uint) (*domain.Product, error)
	ListAll() ([]domain.Product, error)
	// DecrementStock ต้องรับ gorm.DB (tx) เพื่อให้ทำงานใน Transaction ได้
	DecrementStock(tx *gorm.DB, productID uint, amount int) error
}

// RestringingJobRepository
type RestringingJobRepository interface {
	// Create ต้องรับ gorm.DB (tx) เพื่อให้ทำงานใน Transaction ได้
	Create(tx *gorm.DB, job *domain.RestringingJob) error
	ListByCustomerID(customerID uint) ([]domain.RestringingJob, error)
}
