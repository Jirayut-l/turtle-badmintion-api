package main

import (
	"log"

	"github.com/Jirayut-l/turtle-badminton-api/internal/handler/http"
	"github.com/Jirayut-l/turtle-badminton-api/internal/infrastructure/database"

	"github.com/Jirayut-l/turtle-badminton-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// โหลด .env file (ถ้ามี)
	// ถ้า .env ไม่มี, มันจะไปอ่านจาก Environment Variables ที่ Docker Compose ตั้งให้
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found, using environment variables")
	// }

	// 1. Setup Database (Infrastructure)
	// InitDatabase() จะอ่าน ENV (DB_DRIVER, DB_DSN) เอง
	db := database.InitDatabase()

	// 2. Setup Repositories (Infrastructure Layer)
	// "ฉีด" (Inject) gorm.DB instance เข้าไปใน GORM-based repositories
	// นี่คือ implementation จริง
	customerRepo := database.NewGormCustomerRepository(db)
	productRepo := database.NewGormProductRepository(db)
	jobRepo := database.NewGormRestringingJobRepository(db)

	// 3. Setup Usecases (Application Logic Layer)
	// "ฉีด" (Inject) repositories (ที่เป็น interface) เข้าไปใน usecases
	customerUC := usecase.NewCustomerUsecase(customerRepo)
	productUC := usecase.NewProductUsecase(productRepo)
	// RestringingUsecase ต้องการ gorm.DB instance เพื่อจัดการ Transactions
	restringingUC := usecase.NewRestringingUsecase(jobRepo, productRepo, db)

	// 4. Setup Handlers (Delivery Layer)
	// "ฉีด" (Inject) usecases เข้าไปใน Gin handlers
	customerHandler := http.NewCustomerHandler(customerUC)
	productHandler := http.NewProductHandler(productUC)
	restringingHandler := http.NewRestringingHandler(restringingUC)

	// 5. Setup Gin Router (Delivery Layer)
	r := gin.Default()

	// Setup Middlewares
	r.Use(gin.Logger())   // Log requests
	r.Use(gin.Recovery()) // Recover จาก panic

	// (Optional: CORS Middleware)
	// ถ้าคุณจะเรียก API นี้จาก Frontend (เช่น React, Vue) ที่อยู่คนละ Domain
	// คุณจะต้องเปิดใช้งาน CORS
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"}, // ใส่ Origin ของ Frontend
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// }))

	// 6. Setup Routes
	// ส่ง handlers ที่เรา "ฉีด" dependencies จนครบแล้ว ไปให้ SetupRoutes
	http.SetupRoutes(r, customerHandler, productHandler, restringingHandler)

	// 7. Run Server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
