package server

import (
	"net/http"

	md "github.com/LeilaBeken/golang_ass_1/models"
	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	db, err := md.GetDB()
	if err != nil {panic(err)}
	products := []md.Product{}
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	db, err := md.GetDB()
	if err != nil {panic(err)}
	productID := c.Param("id")
	product := md.Product{}
	db.First(&product, productID)
	if product.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	db, err := md.GetDB()
	if err != nil {panic(err)}
	var productData md.Product
	if err := c.ShouldBindJSON(&productData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&productData)
	c.JSON(http.StatusOK, productData)
}

func deleteProduct(c *gin.Context) {
	db, err := md.GetDB()
	if err != nil {panic(err)}
	productID := c.Param("id")
	existingProduct := md.Product{}
	db.First(&existingProduct, productID)
	if existingProduct.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	db.Delete(&existingProduct)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
