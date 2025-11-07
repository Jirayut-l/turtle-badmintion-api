package domain

import "time"

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
