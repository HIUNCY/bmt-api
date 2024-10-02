package main

import (
	"net/http"

	"github.com/HIUNCY/bmt-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/bmt?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{}, &models.Transfer{})

	r := gin.Default()
	dummy := models.User{
		ID:           1,
		FullName:     "Muhamad Zainul Kamal",
		Email:        "nigga@hitam.com",
		PasswordHash: "098u7y6tghuj87654esdrft543ew",
		PhoneNumber:  "0895414726787",
	}
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, dummy)
	})
	r.Run()
}
