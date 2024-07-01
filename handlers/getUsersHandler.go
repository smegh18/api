package handlers

import (
	"book_seller/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) GetUsersHandler(c *gin.Context) {
	rows, error := u.UserDomain.GetUsers()
	if error != nil {
		panic(error)
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}
