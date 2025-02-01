package controllers

import (
	"net/http"
	"tabungku-go/database"
	"tabungku-go/models"
	"tabungku-go/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// store validaiton
type ValidateTabung struct {
	NoRekening string `json:"no_rekening" validate:"required,numeric,max=10"`
	Nominal    int    `json:"nominal" validate:"required,numeric"`
}

func Deposit(c echo.Context) error {
	// request struct validation
	var tabungan ValidateTabung

	// Request Post Parameter, and check body
	if err := c.Bind(&tabungan); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Invalid request body",
		})
	}

	// Validation struct
	validate := validator.New()
	if err := validate.Struct(tabungan); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		utils.Logger.Error(errors)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": errors,
		})
	}

	// request struct model
	var tabunganM models.Tabungan

	// check No Rekening is not already
	checkNoRek := database.DB.Where("no_rekening = ?", tabungan.NoRekening).First(&tabunganM)
	if checkNoRek.Error != nil {
		utils.Logger.Warn("Nomor Rekening " + tabungan.NoRekening + " Not Found")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Nomor Rekening " + tabungan.NoRekening + " Not Found",
		})
	}

	tabunganM.NoRekening = tabungan.NoRekening
	tabunganM.Saldo = tabunganM.Saldo + tabungan.Nominal

	// save to db
	if err := database.DB.Save(&tabunganM).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		})
	}

	// return success
	utils.Logger.Info("Deposit successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Deposit successfully",
		"data": map[string]interface{}{
			"Saldo": tabunganM.Saldo,
		},
	})
}

func Withdraw(c echo.Context) error {
	// request struct validation
	var tabungan ValidateTabung

	// Request Post Parameter, and check body
	if err := c.Bind(&tabungan); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Invalid request body",
		})
	}

	// Validation struct
	validate := validator.New()
	if err := validate.Struct(tabungan); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		utils.Logger.Error(errors)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": errors,
		})
	}

	// request struct model
	var tabunganM models.Tabungan

	// check No Rekening is not already
	checkNoRek := database.DB.Where("no_rekening = ?", tabungan.NoRekening).First(&tabunganM)
	if checkNoRek.Error != nil {
		utils.Logger.Warn("Nomor Rekening " + tabungan.NoRekening + " Not Found")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Nomor Rekening " + tabungan.NoRekening + " Not Found",
		})
	}

	// check saldo, if saldo < nominal = Error
	if tabunganM.Saldo < tabungan.Nominal {
		utils.Logger.Warn("Saldo anda tidak cukup")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Saldo anda tidak cukup",
		})
	}

	tabunganM.NoRekening = tabungan.NoRekening
	tabunganM.Saldo = tabunganM.Saldo - tabungan.Nominal

	// save to db
	if err := database.DB.Save(&tabunganM).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		})
	}

	// return success
	utils.Logger.Info("Withdraw successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Withdraw successfully",
		"data": map[string]interface{}{
			"Saldo": tabunganM.Saldo,
		},
	})
}

func Balance(c echo.Context) error {
	// request param id
	param_no_rek := c.Param("no_rekening")

	// request struct model
	var tabunganM models.Tabungan

	// check No Rekening is not already
	checkNoRek := database.DB.Where("no_rekening = ?", param_no_rek).First(&tabunganM)
	if checkNoRek.Error != nil {
		utils.Logger.Warn("Nomor Rekening " + param_no_rek + " Not Found")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Nomor Rekening " + param_no_rek + " Not Found",
		})
	}

	// return success
	utils.Logger.Info("Check saldo successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Check saldo successfully",
		"data": map[string]interface{}{
			"Saldo": tabunganM.Saldo,
		},
	})
}
