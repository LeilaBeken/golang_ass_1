package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	md "github.com/LeilaBeken/golang_ass_1/models"
)

func HashPassword(password string) ([]byte, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return bytes, err
}

func Register(c *gin.Context) {
    // Parse the request body
    var input struct {
        Username string `json:"username" binding:"required"`
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the password
    hashedPassword, err := HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Create the user
    user := md.User{
    	ID:       0,
    	Username: input.Username,
    	Email:    input.Email,
    	Password: string(hashedPassword),
    }
    if err := user.Create(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return a success response
    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
