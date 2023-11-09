package routers

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/controllers"
	"go-rest-api/middlewares"
	"gorm.io/gorm"
	"net/http"
)

func SetupRouter(e *echo.Echo, db *gorm.DB) {
	e.Use(middlewares.DBMiddleware(db))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello Echo")
	})

	e.GET("/blog", controllers.Read)
	e.GET("/blog/:id", controllers.Find)
	e.POST("/blog", controllers.Create)
	e.PUT("/blog/:id", controllers.Update)
	e.DELETE("/blog/:id", controllers.Delete)
}
