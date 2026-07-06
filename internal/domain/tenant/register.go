package tenant

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func TenantRouter(e *echo.Echo, db *gorm.DB) {
	tenantGroup := e.Group("/api/v1")

	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	tenantGroup.POST("/tenants", h.CreateTenant)
	tenantGroup.GET("/tenants/:id", h.FindById)
	tenantGroup.GET("/tenants/email/:email", h.FindByEmail)
	tenantGroup.PUT("/tenants/:id", h.Update)
}