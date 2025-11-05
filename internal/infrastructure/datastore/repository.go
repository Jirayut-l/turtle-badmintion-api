package datastore

import (
	"errors"

	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
	"github.com/Jirayut-l/turtle-badminton-api/internal/repository"
	"gorm.io/gorm"
)

// นี่คือ "ของจริง" ที่ implement interfaces จาก repository
// โดยใช้ GORM

// GormCustomerRepository
type gormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &gormCustomerRepository{db: db}
}

func (r *gormCustomerRepository) Create(customer *domain.Customer) error {
	return r.db.Create(customer).Error
}

func (r *gormCustomerRepository) FindByID(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.First(&customer, id).Error
	return &customer, err
}

func (r *gormCustomerRepository) FindByPhone(phone string) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.Where("phone = ?", phone).First(&customer).Error
	return &customer, err
}

func (r *gormCustomerRepository) Update(customer *domain.Customer) error {
	return r.db.Save(customer).Error
}

// GormProductRepository
type gormProductRepository struct {
	db *gorm.DB
}

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

// GormRestringingJobRepository
type gormRestringingJobRepository struct {
	db *gorm.DB
}

func NewGormRestringingJobRepository(db *gorm.DB) repository.RestringingJobRepository {
	return &gormRestringingJobRepository{db: db}
}

func (r *gormRestringingJobRepository) Create(tx *gorm.DB, job *domain.RestringingJob) error {
	// ใช้ tx (Transaction) ที่ส่งมาจาก Usecase
	return tx.Create(job).Error
}

func (r *gormRestringingJobRepository) ListByCustomerID(customerID uint) ([]domain.RestringingJob, error) {
	var jobs []domain.RestringingJob
	// โหลดข้อมูล Relation (Preload) มาด้วย
	err := r.db.Preload("Product").Where("customer_id = ?", customerID).Find(&jobs).Error
	return jobs, err
}
