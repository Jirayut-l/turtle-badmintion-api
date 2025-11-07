package database

import (
	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
	"github.com/Jirayut-l/turtle-badminton-api/internal/repository"
	"gorm.io/gorm"
)

// GormCustomerRepository
type gormCustomerRepository struct {
	db *gorm.DB
}

// NewGormCustomerRepository คือ constructor
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
