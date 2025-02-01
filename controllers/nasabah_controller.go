package controllers

import (
	"fmt"
	"net/http"
	"tabungku-go/database"
	"tabungku-go/models"
	"tabungku-go/utils"
	"time"

	"math/rand"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// genereate random number (no rekening)
func generateRandNumber(length int) string {
	min := 1
	max := 9
	var result string

	// sett seed for random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		if i == 0 {
			// Digit pertama tidak boleh 0
			result += fmt.Sprintf("%d", rand.Intn(max)+min)
		} else {
			result += fmt.Sprintf("%d", rand.Intn(10)) // Digit berikutnya boleh 0
		}
	}
	return result
}

// store validaiton
type NasabahValStore struct {
	Nik  string `json:"nik" validate:"required,numeric,max=16"`
	Name string `json:"name" validate:"required,min=2,max=16"`
	NoHp string `json:"no_hp" validate:"required,numeric,max=15"`
}

func StoreNasabah(c echo.Context) error {

	// request struct validation
	var nasabah NasabahValStore

	// Request Post Parameter, and check body
	if err := c.Bind(&nasabah); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Invalid request body",
		})
	}

	// Validation struct
	validate := validator.New()
	if err := validate.Struct(nasabah); err != nil {
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
	var nasabahM models.Nasabah

	// check NIK is already
	checkNIK := database.DB.Where("nik = ?", nasabah.Nik).First(&nasabahM)
	if checkNIK.Error == nil {
		utils.Logger.Warn("NIK " + nasabah.Nik + " is already registered")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "NIK " + nasabah.Nik + " is already registered.",
		})
	}

	// check No HP is ready
	checkNoHP := database.DB.Where("no_hp = ?", nasabah.NoHp).First(&nasabahM)
	if checkNoHP.Error == nil {
		utils.Logger.Warn("No HP " + nasabah.NoHp + " is already registered")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "No HP " + nasabah.NoHp + " is already registered.",
		})
	}

	param := models.Nasabah{
		Nik:       nasabah.Nik,
		Name:      nasabah.Name,
		NoHp:      nasabah.NoHp,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// create to db (nasabah)
	if err := database.DB.Create(&param).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		})
	}

	resultRekening := generateRandNumber(10)

	paramTabungan := models.Tabungan{
		IdNasabah:  int(param.ID),
		NoRekening: resultRekening,
		Saldo:      0,
	}
	// create to db (tabungan by id nasabah)
	if err := database.DB.Create(&paramTabungan).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		})
	}

	// return success
	utils.Logger.Info("Created successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Created successfully",
		"data": map[string]interface{}{
			"nama":           param.Name,
			"Nomor Rekening": paramTabungan.NoRekening,
			"Saldo":          paramTabungan.Saldo,
		},
	})
}
