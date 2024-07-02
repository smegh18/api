package main

import (
	"book_seller/domain"
	"book_seller/handlers"
	"book_seller/middleware"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	var db *sql.DB
	postgresConnectionString := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	postgresConnectionString = fmt.Sprintf(postgresConnectionString, "localhost", "5432", "postgres", "lolpro123", "books")
	db, err := sql.Open("postgres", postgresConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userService := domain.NewDomainDB(db)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/login", userHandler.LoginHandler)
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.POST("/users", userHandler.CreateUserHandler)
	auth.GET("/users", userHandler.GetUsersHandler)
	auth.GET("/users/:id", userHandler.GetUserHandler)
	auth.PUT("/users/:id", userHandler.UpdateUserHandler)
	auth.DELETE("/users/:id", userHandler.DeleteUserHandler)

	r.Run(":8080")
}
