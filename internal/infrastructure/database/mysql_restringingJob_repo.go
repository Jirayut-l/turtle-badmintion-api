package database

import (
	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
	"github.com/Jirayut-l/turtle-badminton-api/internal/repository"
	"gorm.io/gorm"
)

// GormRestringingJobRepository
type gormRestringingJobRepository struct {
	db *gorm.DB
}

// NewGormRestringingJobRepository คือ constructor
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
