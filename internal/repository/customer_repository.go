package repository

import "github.com/Jirayut-l/turtle-badminton-api/internal/domain"

// นี่คือ "สัญญา" (Contract) ที่ชั้น Usecase จะเรียกใช้
// โดยไม่สนว่าข้างหลังบ้านจะใช้ GORM, SQLC หรืออะไรก็ตาม

// CustomerRepository
type CustomerRepository interface {
	Create(customer *domain.Customer) error
	FindByID(id uint) (*domain.Customer, error)
	FindByPhone(phone string) (*domain.Customer, error)
	Update(customer *domain.Customer) error
}
