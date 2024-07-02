package handlers

import (
	"book_seller/middleware"
	"book_seller/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) LoginHandler(c *gin.Context) {
	var credentials model.User
	// name := c.Param("name")
	// credentials.Password = c.Param("password")
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	storedPassword, email, err := u.UserDomain.Login(credentials.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"pwd":   storedPassword,
			"email": email,
			"error": err,
		})
		return
	}
	if credentials.Password != storedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := middleware.GenerateJWT(credentials.Name, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
