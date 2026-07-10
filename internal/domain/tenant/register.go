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
	tenantGroup.GET("/tenants", h.FindAll)
	tenantGroup.GET("/tenants/:id", h.FindById)
	tenantGroup.GET("/tenants/email/:email", h.FindByEmail)
	tenantGroup.PUT("/tenants/:id", h.Update)
	tenantGroup.DELETE("/tenants/:id", h.Delete)
}

/*
POST   /api/tenants          -> create tenant
GET    /api/tenants          -> get all tenants
GET    /api/tenants/:id      -> get tenant
PUT    /api/tenants/:id      -> update tenant
DELETE /api/tenants/:id     -> delete tenant
*/