package server

import (
	"net/http"
	"fmt"
	md "github.com/LeilaBeken/golang_ass_1/models"
	"github.com/gin-gonic/gin"
)

func searchProducts(c *gin.Context) {
    // Retrieve the search query from the request
    searchQuery := c.Query("q")

    // Retrieve the products from the database
    products := []md.Product{}
	db, err := md.GetDB()
	if err != nil {panic(err)}
    db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", searchQuery)).Find(&products)

    // Return the products in the response
    c.JSON(http.StatusOK, products)
}
