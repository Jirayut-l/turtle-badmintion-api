package database

import (
	"log"
	"os"
	"time" // เพิ่ม import "time"

	"github.com/Jirayut-l/turtle-badminton-api/internal/domain"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	// อ่านค่า DSN (Data Source Name) จาก Environment
	// เราไม่ต้องการ DB_DRIVER อีกต่อไป
	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set. Please check .env or Docker config.")
		os.Exit(1)
	}

	// Log ระดับ Info
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// ลบ Switch case และเชื่อมต่อ MySQL โดยตรง
	// --- เพิ่ม Retry Logic ---
	var retries int = 5
	for {
		log.Println("Connecting to MySQL Database...")
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)

		if err == nil {
			// เชื่อมต่อสำเร็จ
			break
		}

		retries--
		if retries <= 0 {
			log.Fatalf("Failed to connect to database after retries: %v", err)
			os.Exit(1)
		}

		log.Printf("Failed to connect, retrying in 3 seconds... (%d retries left)", retries)
		time.Sleep(3 * time.Second)
	}
	// --- จบ Retry Logic ---

	// AutoMigrate (อัปเดตจาก ... (existing code) ...
	log.Println("Running Database Migrations...")
	err = db.AutoMigrate(
		&domain.Customer{},
		&domain.Product{},
		&domain.RestringingJob{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}

	log.Println("Database connection successful.")
	return db
}
