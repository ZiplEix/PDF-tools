package routes

import (
	"github.com/ZiplEix/PDF-tools/controllers"
	"github.com/labstack/echo/v4"
)

func setupPDFRoutes(e *echo.Echo) {
	g := e.Group("/pdf")

	g.POST("/merge", controllers.MergePDF)
}
