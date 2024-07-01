package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *pgxpool.Pool

func initDB() {
	dsn := "postgresql://postgres:lolpro123@localhost:5432/books"
	var err error
	db, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	fmt.Println("Database connection successfully established")
}

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(context.Background(), sql, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not founc"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func getUsers(c *gin.Context) {
	rows, err := db.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	err := db.QueryRow(context.Background(), "SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	_, err := db.Exec(context.Background(), sql, user.Name, user.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	sql := `DELETE FROM users WHERE id=$1`
	_, err := db.Exec(context.Background(), sql, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// GetUser fetches the user having the specific userId
// func (u *UserService) GetUser(userID *uuid.UUID) (*model.UserResponse, error) {
// 	user := &model.UserResponse{}
// 	contactInfo := &model.ContactInfo{}
// 	authority := &model.Authority{}

// 	sqlStatement := `
// 	SELECT
// 		"name",
// 		"email",
// 		"mobileNumber",
// 		"countryCode",
// 		"userType",
// 		"IDNumber",

// 		CASE
// 		WHEN "IDNumber" IS NULL THEN $1
// 		ELSE "designation" END AS "designation",

// 		CASE WHEN "errorMap" IS NULL THEN '{}'::jsonb
// 		ELSE "errorMap" END AS "errorMap",

// 		"createdAtUTC",
// 		"updatedAtUTC"
// 	FROM
// 		"users"
// 	WHERE
// 		"userID" = $2;
// 	`

// 	if err := u.db.QueryRow(sqlStatement, model.NOT_APPLICABLE_DESIGNATION, userID).Scan(&user.Name, &user.Email, &contactInfo.MobileNumber, &contactInfo.TelCountryCode, &user.Type, &authority.IdNumber, &authority.Designation, &user.ErrorMap, &user.CreatedAtUTC, &user.UpdatedAtUTC); err != nil {
// 		if errors.Is(sql.ErrNoRows, err) {
// 			return nil, nil
// 		}
// 		log.Error().Msgf("GetUser() Error: %v", err)
// 		return nil, ErrGetUserFailed
// 	}

// 	user.Authority = authority
// 	user.ContactInfo = contactInfo

// 	return user, nil
// }
