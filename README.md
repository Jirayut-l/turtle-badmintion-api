**Turtle Badminton API**
ระบบหลังบ้านสำหรับจัดการสต็อกสินค้า (เอ็นแบดมินตัน) และบันทึกข้อมูลลูกค้าร้านแบดมินตัน สร้างด้วยสถาปัตยกรรม Clean Architecture

**🚀 Tech Stack**
- Language: Go (Golang)
- Framework: Gin (HTTP Web Framework)
- Database: MySQL 8.0
- ORM: GORM
- Architecture: Clean Architecture
- Containerization: Docker & Docker Compose
- Testing: Go Test, Testify (Asserts), Mockery (Mocks), SQLMock (DB Mocking)

**🏁 Getting Started**
วิธีที่ง่ายและแนะนำที่สุดในการรันโปรเจกต์นี้คือการใช้ Docker Compose

**1.(ครั้งแรก) สร้างไฟล์ Environment**
โปรเจกต์นี้ใช้ Environment Variables ในการตั้งค่า Database
    1. คัดลอกไฟล์ .env.example ไปเป็นไฟล์ .envcp .env.example .env
(ไม่จำเป็น) แก้ไขรหัสผ่านในไฟล์ .env (รหัสผ่านนี้สำหรับ การรันบนเครื่อง เท่านั้น)2. รันด้วย Docker Compose (Recommended)คำสั่งนี้จะสร้างและรัน Go service และ MySQL service ให้พร้อมกันในครั้งเดียวdocker-compose up --build
--build ใช้สำหรับการรันครั้งแรก หรือเมื่อมีการแก้ไขโค้ด GoAPI Server จะรันที่: http://localhost:8080MySQL Database จะเชื่อมต่อได้ที่: localhost:3306 (จากโปรแกรมภายนอกเช่น DBeaver/TablePlus)หากต้องการหยุด service:docker-compose down
3. (ทางเลือก) รันบนเครื่อง (Local Development)หากคุณไม่ต้องการใช้ Docker และมี MySQL ติดตั้งบนเครื่องของคุณอยู่แล้ว:Start MySQL: ตรวจสอบให้แน่ใจว่า MySQL service ของคุณกำลังรันอยู่Set Environment: ตรวจสอบไฟล์ .env และแก้ไข DB_DSN ให้ชี้ไปที่ localhost ของคุณDB_DSN=admin:your_admin_secret_password@tcp(localhost:3306)/turtle_db?charset=utf8mb4&parseTime=True&loc=Local
Install Dependencies:go mod tidy
Run:go run ./cmd/server/main.go
🧪 Running Testsโปรเจกต์นี้มี Unit Tests สำหรับ Usecase (Business Logic)รัน Test ทั้งหมด:go test ./...
(ถ้าจำเป็น) สร้าง Mocks ใหม่:หากคุณมีการแก้ไข interfaces.go คุณต้องสร้าง Mocks ใหม่โดยใช้ mockery:# (ติดตั้ง mockery ก่อน ถ้ายังไม่มี)
# go install [github.com/vektra/mockery/v2@latest](https://github.com/vektra/mockery/v2@latest)

# รัน mockery
mockery --all --dir=internal/repository --output=internal/repository/mocks
📂 Project Structure (Clean Architecture)โครงสร้างนี้แบ่งแยกหน้าที่ความรับผิดชอบ (Separation of Concerns) อย่างชัดเจน:/turtle-badminton-api
├── cmd/server/
│   └── main.go           # จุดเริ่มต้น, ประกอบร่าง (Dependency Injection)
├── internal/
│   ├── domain/
│   │   └── models.go       # (Entities) โครงสร้างข้อมูลหลัก
│   ├── repository/
│   │   ├── interfaces.go   # (Interface) "สัญญา" ที่ Usecase จะเรียก
│   │   └── mocks/          # Mocks ที่สร้างโดย Mockery
│   ├── usecase/
│   │   ├── usecases.go       # (Business Logic) ตรรกะของระบบ
│   │   └── usecases_test.go  # Unit Tests
│   ├── delivery/http/
│   │   ├── handlers.go     # (Handlers) ตัวรับ Request จาก Gin
│   │   ├── routes.go       # (Routes) กำหนด API Endpoints
│   │   └── dto.go          # (DTOs) Models สำหรับ Request/Response
│   └── infrastructure/datastore/
│       ├── database.go     # โค้ดเชื่อมต่อ GORM
│       └── gorm_repo.go    # (Implementation) โค้ดจริงที่เชื่อม DB
├── go.mod
├── Dockerfile              # พิมพ์เขียวสำหรับ Go Service
├── docker-compose.yml      # ไฟล์สั่งรัน Go + MySQL
└── .env                    # (Git Ignored) เก็บค่า Secret
📖 API Endpointsคุณสามารถทดสอบ API ด้วยเครื่องมืออย่าง Postman หรือ InsomniaCustomersPOST /api/v1/customersDescription: ลงทะเบียนลูกค้าใหม่Body (JSON):{
    "name": "Somsak Jaidee",
    "phone": "0812345678",
    "notes": "ไม้ประจำ: Yonex Arcsaber 11"
}
GET /api/v1/customers?phone=...Description: ค้นหาลูกค้าด้วยเบอร์โทรExample: GET http://localhost:8080/api/v1/customers?phone=0812345678ProductsGET /api/v1/productsDescription: ดูรายการสินค้า (เอ็น) ทั้งหมดในสต็อกRestringing JobsPOST /api/v1/jobsDescription: บันทึกงานขึ้นเอ็นใหม่ (และ ตัดสต็อก สินค้า 1 ชิ้น)Body (JSON):{
    "customer_id": 1,
    "product_id": 2,
    "tension": "24 lbs",
    "price_charged": 350.00,
    "notes": "ลูกค้าขอ 4 ปม"
}
