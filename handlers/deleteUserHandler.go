package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	if err := u.UserDomain.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
