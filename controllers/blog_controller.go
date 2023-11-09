package controllers

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/helper"
	"go-rest-api/models"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
)

func Read(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var blog []models.Blog
	db.Find(&blog)
	return c.JSON(http.StatusOK, blog)
}

func Find(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var blog models.Blog

	if err := db.First(&blog, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Record Not Found"})
	}

	return c.JSON(http.StatusOK, blog)
}

func Create(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	// parse the form data to get the uploaded file
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}

	// get file from the form data
	file, err := form.File["path"][0].Open()
	if err != nil {
		return c.JSON(http.StatusBadGateway, "Failed to retrieve the file")
	}
	defer file.Close()

	// access the filename from the form data
	generateNameFile := helper.GenerateUniqueFileName(form.File["path"][0].Filename)
	dst, err := os.Create("public/" + generateNameFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the content of the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	var data struct {
		blog     models.Blog
		category models.Category
		user     models.User
		image    models.Image
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&data.blog)
	db.Create(&data.category)
	db.Create(&data.user)
	db.Create(&data.image)

	return c.JSON(http.StatusCreated, data)
}

func Update(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var blog models.Blog

	if err := db.First(&blog, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Record Not Found"})
	}

	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Save(&blog)
	return c.JSON(http.StatusOK, blog)
}

func Delete(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var blog models.Blog

	if err := db.First(&blog, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Record Not Found"})
	}

	db.Delete(&blog)
	return c.JSON(http.StatusOK, blog)
}
