# --- Stage 1: Build ---
# ใช้ Go 1.21 (หรือเวอร์ชันที่คุณใช้) เป็น base image สำหรับ build
FROM golang:1.25.3-alpine AS builder

# ตั้งค่า Working Directory
WORKDIR /app

# Copy go.mod และ go.sum
# แยก layer นี้เพื่อ cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code ที่เหลือ
COPY . .

# Build Go app
# CGO_ENABLED=0 สร้าง static binary (สำคัญมากสำหรับ alpine)
# GOOS=linux build สำหรับ linux (Docker)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./main ./cmd/server/main.go

# --- Stage 2: Final Image ---
# ใช้ Alpine image (เล็กมาก) เป็น base สุดท้าย
FROM alpine:latest

# Alpine ไม่มี root CA certificates, ต้องติดตั้งเผื่อ service ต้องเรียก API ภายนอก (HTTPS)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary ที่ build เสร็จแล้วจาก stage 'builder'
COPY --from=builder /app/main .

# (ถ้ามี) Copy .env.production หรือไฟล์ config อื่นๆ
# COPY .env.production .env

# Expose port ที่ Gin server รันอยู่ (จาก main.go)
EXPOSE 8080

# Command ที่จะรันเมื่อ container เริ่มทำงาน
CMD ["./main"]