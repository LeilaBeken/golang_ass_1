package main


router := gin.Default()

// Define the routes
router.GET("/products", listProducts)
router.GET("/products/:id", getProduct)
router.POST("/products", createProduct)
router.PUT("/products/:id", updateProduct)
router.DELETE("/products/:id", deleteProduct)