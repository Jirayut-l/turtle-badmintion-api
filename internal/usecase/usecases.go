package usecase

import (
	"errors"
	"time"

	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
	"github.com/Jirayut-l/turtle-badminton-api/internal/repository"
	"gorm.io/gorm"
)

// CustomerUsecase
type CustomerUsecase struct {
	repo repository.CustomerRepository
}

func NewCustomerUsecase(repo repository.CustomerRepository) *CustomerUsecase {
	return &CustomerUsecase{repo: repo}
}

func (uc *CustomerUsecase) RegisterCustomer(name, phone, notes string) (*domain.Customer, error) {
	customer := &domain.Customer{
		Name:  name,
		Phone: phone,
		Notes: notes,
	}
	err := uc.repo.Create(customer)
	return customer, err
}

func (uc *CustomerUsecase) GetCustomerByPhone(phone string) (*domain.Customer, error) {
	return uc.repo.FindByPhone(phone)
}

// ProductUsecase
type ProductUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) AddNewProduct(sku, brand, model string, quantity int, cost, selling float64) (*domain.Product, error) {
	product := &domain.Product{
		SKU:             sku,
		Brand:           brand,
		Model:           model,
		QuantityInStock: quantity,
		CostPrice:       cost,
		SellingPrice:    selling,
	}
	err := uc.repo.Create(product)
	return product, err
}

func (uc *ProductUsecase) ListAllProducts() ([]domain.Product, error) {
	return uc.repo.ListAll()
}

// RestringingUsecase
// นี่คือ Usecase ที่ซับซ้อนที่สุด เพราะต้อง "ตัดสต็อก" (ระบบตัดสินค้าที่ใช้งานแล้ว)
type RestringingUsecase struct {
	jobRepo     repository.RestringingJobRepository
	productRepo repository.ProductRepository
	db          *gorm.DB // ต้องใช้ db instance เพื่อเริ่ม Transaction
}

func NewRestringingUsecase(jobRepo repository.RestringingJobRepository, productRepo repository.ProductRepository, db *gorm.DB) *RestringingUsecase {
	return &RestringingUsecase{
		jobRepo:     jobRepo,
		productRepo: productRepo,
		db:          db,
	}
}

// LogNewJob คือหัวใจของการ "ตัดสต็อก"
// เราใช้ GORM Transaction เพื่อให้แน่ใจว่า
// 1. การตัดสต็อก
// 2. การสร้าง Job Log
// เกิดขึ้นพร้อมกันทั้งหมด หรือไม่เกิดขึ้นเลย (Atomic)
func (uc *RestringingUsecase) LogNewJob(customerID, productID uint, tension string, price float64, notes string) (*domain.RestringingJob, error) {
	var job *domain.RestringingJob

	// เริ่ม Transaction
	err := uc.db.Transaction(func(tx *gorm.DB) error {
		// 1. ตรวจสอบสินค้าและสต็อก (ภายใน Transaction)
		product, err := uc.productRepo.FindByID(productID)
		if err != nil {
			return err
		}
		if product.QuantityInStock < 1 {
			return errors.New("สินค้าหมดสต็อก (product out of stock)")
		}

		// 2. ตัดสต็อก (ส่ง tx เข้าไปใน repository)
		if err := uc.productRepo.DecrementStock(tx, productID, 1); err != nil {
			return err
		}

		// 3. สร้าง Job Log (ส่ง tx เข้าไปใน repository)
		job = &domain.RestringingJob{
			CustomerID:   customerID,
			ProductID:    productID,
			Tension:      tension,
			PriceCharged: price,
			JobDate:      time.Now(),
			Notes:        notes,
		}
		if err := uc.jobRepo.Create(tx, job); err != nil {
			return err
		}

		// ถ้าทุกอย่างสำเร็จ, Transaction จะ commit
		return nil
	})

	// ถ้า err != nil, Transaction จะ Rollback อัตโนมัติ
	if err != nil {
		return nil, err
	}

	// คืนค่า job ที่สร้างเสร็จ
	// เราต้องโหลด Customer และ Product เข้ามาด้วย
	// (GORM จะไม่โหลด Relation ให้อัตโนมัติถ้าไม่สั่ง)
	uc.db.Preload("Customer").Preload("Product").First(&job, job.ID)
	return job, nil
}
