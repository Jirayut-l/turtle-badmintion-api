package domain

import "time"

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
