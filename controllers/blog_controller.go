package controllers

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/models"
	"gorm.io/gorm"
	"net/http"
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
	var blog models.Blog
	var category models.Category
	var image models.Image

	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else if err := c.Bind(&image); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Create(&blog)
	db.Create(&category)
	db.Create(&image)

	return c.JSON(http.StatusCreated, blog)
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
