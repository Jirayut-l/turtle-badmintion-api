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
