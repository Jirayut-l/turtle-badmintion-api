package domain

import "time"

// Customer (ลูกค้า)
// เก็บข้อมูลพื้นฐานของลูกค้า
type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`             // <-- เพิ่ม type
	Phone     string    `json:"phone" gorm:"type:varchar(20);uniqueIndex;not null"` // <-- แก้ไข: ระบุ type เป็น varchar(20)
	Notes     string    `json:"notes"`                                              // <-- field นี้เป็น string (longtext) ได้ เพราะเราไม่ได้ index
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Product (สินค้า)
// ในที่นี้คือ "เอ็น" (String) จากไฟล์ CSV ของคุณ
type Product struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	SKU             string    `json:"sku" gorm:"type:varchar(100);uniqueIndex;not null"` // <-- แก้ไข: ระบุ type เป็น varchar(100)
	Brand           string    `json:"brand" gorm:"type:varchar(100)"`                    // <-- เพิ่ม type
	Model           string    `json:"model" gorm:"type:varchar(100)"`                    // <-- เพิ่ม type
	QuantityInStock int       `json:"quantity_in_stock" gorm:"not null"`                 // จำนวนคงเหลือในสต็อก
	CostPrice       float64   `json:"cost_price"`                                        // ราคาทุน (จาก CSV)
	SellingPrice    float64   `json:"selling_price"`                                     // ราคาขาย (ที่จะคิดกับลูกค้า)
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// RestringingJob (งานขึ้นเอ็น)
// นี่คือ "ธุรกรรม" ที่จะเชื่อมโยงลูกค้าและสินค้าเข้าด้วยกัน
// และเป็นตัว "ตัดสต็อก" สินค้า
type RestringingJob struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerID   uint      `json:"customer_id"` // เชื่อมโยงไปยัง Customer
	Customer     Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	ProductID    uint      `json:"product_id"` // เชื่อมโยงไปยัง Product (เอ็นที่ใช้)
	Product      Product   `gorm:"foreignKey:ProductID" json:"product"`
	Tension      string    `json:"tension" gorm:"type:varchar(50)"` // <-- เพิ่ม type (เช่น "24 lbs")
	PriceCharged float64   `json:"price_charged"`                   // ราคาที่เก็บเงินลูกค้า
	JobDate      time.Time `json:"job_date"`
	Notes        string    `json:"notes"` // <-- field นี้เป็น string (longtext) ได้
	CreatedAt    time.Time `json:"created_at"`
}
