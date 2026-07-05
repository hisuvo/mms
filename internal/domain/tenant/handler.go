package tenant

import (
	"mms-dbsd/internal/domain/tenant/dto"
	"mms-dbsd/internal/httpresponse"
	"mms-dbsd/internal/validator"
	"net/http"

	"github.com/labstack/echo/v5"
)



type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateTenant(c *echo.Context) error {
	var ten dto.CreateTenantRequest

	if err := c.Bind(&ten); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.BadRequest("invalid_request_body"))
	}
	
	if err := c.Validate(&ten); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ValidationError(validator.FormatValidationError(err)))
	}

	res, err := h.service.CreateTenant(&ten)
	if err != nil {
		return c.JSON(http.StatusConflict, httpresponse.Conflict(err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}