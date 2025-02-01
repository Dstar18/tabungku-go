package main

import (
	"net/http"
	"tabungku-go/controllers"
	"tabungku-go/database"
	"tabungku-go/models"
	"tabungku-go/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize logger
	utils.InitLogger()

	// connect database
	database.Connect()

	// Migrate schema
	database.DB.AutoMigrate(&models.Nasabah{})
	database.DB.AutoMigrate(&models.Tabungan{})

	// initialize echo
	e := echo.New()

	// Logging start apps
	utils.Logger.Info("Starting server on port 8080")

	// public routes
	e.GET("hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Golang echo!")
	})

	e.POST("/daftar", controllers.StoreNasabah)
	e.POST("/tabung", controllers.Tabung)

	e.Logger.Fatal(e.Start(":8080"))
}
