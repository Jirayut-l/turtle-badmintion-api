package database

import (
	"errors"

	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
	"github.com/Jirayut-l/turtle-badminton-api/internal/repository"
	"gorm.io/gorm"
)

// GormProductRepository
type gormProductRepository struct {
	db *gorm.DB
}

// NewGormProductRepository คือ constructor
func NewGormProductRepository(db *gorm.DB) repository.ProductRepository {
	return &gormProductRepository{db: db}
}

func (r *gormProductRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *gormProductRepository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *gormProductRepository) ListAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *gormProductRepository) DecrementStock(tx *gorm.DB, productID uint, amount int) error {
	// ใช้ tx (Transaction) ที่ส่งมาจาก Usecase
	// ใช้ GORM Expression เพื่อป้องกัน Race Condition
	result := tx.Model(&domain.Product{}).
		Where("id = ? AND quantity_in_stock >= ?", productID, amount).
		Update("quantity_in_stock", gorm.Expr("quantity_in_stock - ?", amount))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("ไม่สามารถตัดสต็อกได้ (อาจจะสต็อกไม่พอ หรือ ID ผิด)")
	}
	return nil
}
