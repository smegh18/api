package main

import (
	"book_seller/domain"
	"book_seller/handlers"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	postgresConnectionString := "host=%s port=%s user=%s password=%s dbname='%s' sslmode=disable"
	postgresConnectionString = fmt.Sprintf(postgresConnectionString, "localhost", 5432, "postgres", "lolpro123", "books")
	db, err := sql.Open("postgres", postgresConnectionString)
	if err != nil {
		panic("Databse not found")
	}
	defer db.Close()
	userService := domain.NewDomainDB(db)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/users", userHandler.CreateUserHandler)
	r.GET("/users", userHandler.GetUsersHandler)
	r.GET("/users/:id", userHandler.GetUserHandler)
	r.PUT("/users/:id", userHandler.UpdateUserHandler)
	r.DELETE("/users/:id", userHandler.DeleteUserHandler)

	r.Run(":8080")
}
