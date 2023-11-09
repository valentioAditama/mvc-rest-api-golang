package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-rest-api/models"
	"go-rest-api/routers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// define the mysql database connecting
	dsn := "root:@tcp(127.0.0.1:3306)/simple_blog_golang?charset=utf8mb4&parseTime=True&loc=Local"

	// establish a mysql database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database" + err.Error())
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Blog{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Image{})

	routers.SetupRouter(e, db)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.Logger.Fatal(e.Start(":8081"))
}
