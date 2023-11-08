package controllers

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/models"
	"gorm.io/gorm"
	"net/http"
)

func read(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var blog []models.Blog
	db.Find(&blog)
	return c.JSON(http.StatusOK, blog)
}

func find(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var blog models.Blog

	if err := db.First(&blog, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Record Not Found"})
	}

	return c.JSON(http.StatusOK, blog)
}

func create(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var blog models.Blog
	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	db.Create(&blog)
	return c.JSON(http.StatusCreated, blog)
}

func update(c echo.Context) error {
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

func delete(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var blog models.Blog

	if err := db.First(&blog, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Record Not Found"})
	}

	db.Delete(&blog)
	return c.JSON(http.StatusOK, blog)
}
