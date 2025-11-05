package http

import "github.com/gin-gonic/gin"

// SetupRoutes กำหนด endpoints ทั้งหมดของ API
func SetupRoutes(
	r *gin.Engine,
	customerHandler *CustomerHandler,
	productHandler *ProductHandler,
	restringingHandler *RestringingHandler,
) {
	// จัดกลุ่ม API v1
	api := r.Group("/api/v1")
	{
		// Customer routes
		custGroup := api.Group("/customers")
		{
			custGroup.POST("/", customerHandler.RegisterCustomer)
			custGroup.GET("/", customerHandler.GetCustomer) // GET /api/v1/customers?phone=...
		}

		// Product routes
		prodGroup := api.Group("/products")
		{
			prodGroup.GET("/", productHandler.ListProducts)
			// prodGroup.POST("/", productHandler.CreateProduct) // (Endpoint สำหรับเพิ่มสินค้าใหม่)
		}

		// Job routes (การตัดสต็อก)
		jobGroup := api.Group("/jobs")
		{
			jobGroup.POST("/", restringingHandler.LogNewJob)
			// jobGroup.GET("/customer/:id", ...) // (Endpoint สำหรับดูประวัติการขึ้นเอ็นของลูกค้า)
		}
	}
}
