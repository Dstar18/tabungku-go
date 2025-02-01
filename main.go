package main

import (
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
	e.POST("daftar", controllers.StoreNasabah)
	e.POST("tabung", controllers.Deposit)
	e.POST("tarik", controllers.Withdraw)
	e.GET("saldo/:no_rekening", controllers.Balance)

	e.Logger.Fatal(e.Start(":8080"))
}
