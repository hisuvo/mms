package users

import (
	"mms-dbsd/internal/config"
	"mms-dbsd/internal/domain/tenant"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func UserRegister(e *echo.Echo, db *gorm.DB, cfg config.Config) {

	user := e.Group("/api/v1/auth")

	tRepo := tenant.NewRepository(db)
	repo := NewRegisterRepository(db)
	svc := NewRegisterService(repo,tRepo)
	h := NewUserHandler(svc)

	user.POST("/register", h.Register)

}