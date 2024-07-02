package handlers

import (
	"book_seller/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	user, err := u.UserDomain.GetUser(id, model.User{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}
