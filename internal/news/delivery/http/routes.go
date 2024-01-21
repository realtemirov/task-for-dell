package http

import (
	"github.com/labstack/echo/v4"
	"github.com/realtemirov/task-for-dell/internal/news"
)

// MapNewsRoutes maps routes for newss
func MapNewsRoutes(newsGroup *echo.Group, h news.Handlers) {
	newsGroup.POST("", h.Create())
	newsGroup.PUT("/:id", h.Update())
	newsGroup.DELETE("/:id", h.Delete())
	newsGroup.GET("/:id", h.GetByID())
	newsGroup.GET("", h.GetAll())
}
