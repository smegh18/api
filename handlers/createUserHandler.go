package handlers

import (
	"book_seller/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) CreateUserHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.UserDomain.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)
}
