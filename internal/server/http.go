package server

import (
	"mms-dbsd/internal/config"
	"mms-dbsd/internal/utils"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

func Start(db *gorm.DB, cfg *config.Config){
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.Validator = utils.NewValidator()

	// GET /welcome api
	e.GET("/",func(c *echo.Context) error {
		return c.String(http.StatusOK,"Welcome to the Messify beackend!")
	})

	// All Routes:
	// users.RegisterRoute(e, db)
	

	port := cfg.PORT
	if err := e.Start(":"+port); err != nil {
    e.Logger.Error("failed to start server", "error", err)
  }
}