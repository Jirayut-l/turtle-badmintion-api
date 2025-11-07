package repository

import (
	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
	"gorm.io/gorm"
)

// ProductRepository
type ProductRepository interface {
	Create(product *domain.Product) error
	FindByID(id uint) (*domain.Product, error)
	ListAll() ([]domain.Product, error)
	// DecrementStock ต้องรับ gorm.DB (tx) เพื่อให้ทำงานใน Transaction ได้
	DecrementStock(tx *gorm.DB, productID uint, amount int) error
}
