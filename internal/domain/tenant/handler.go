package tenant

import (
	"mms-dbsd/internal/domain/tenant/dto"
	"mms-dbsd/internal/httpresponse"
	"mms-dbsd/internal/validator"
	"net/http"
	"strconv"

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

	tenant, err := h.service.CreateTenant(&ten)
	if err != nil {
		return c.JSON(http.StatusConflict, httpresponse.Conflict(err.Error()))
	}

	response := httpresponse.NewSuccessResponse("Tenant created successfully", tenant)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) FindById(c *echo.Context) error {
	idStr := c.Param("id")

	id64, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.BadRequest("invalid_request_body"))
	}

	tenant, err := h.service.FindById(uint(id64))
	if err != nil {
		return c.JSON(http.StatusNotFound, httpresponse.NotFound(err.Error()))
	}

	response := httpresponse.NewSuccessResponse("Tenant retrieved successfully", tenant)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) FindByEmail(c *echo.Context) error {
	email := c.Param("email")

	tenant, err := h.service.FindByEmail(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, httpresponse.NotFound(err.Error()))
	}

	response := httpresponse.NewSuccessResponse("Tenant retrieved successfully", tenant)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Update(c *echo.Context) error {
	idStr := c.Param("id")

	id64, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.BadRequest("invalid_request_body"))
	}

	var ten dto.UpdateTenantRequest

	if err := c.Bind(&ten); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.BadRequest("invalid_request_body"))
	}

	if err := c.Validate(&ten); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ValidationError(validator.FormatValidationError(err)))
	}

	 update, err := h.service.Update(uint(id64), ten);

	if err != nil {
		return c.JSON(http.StatusConflict, httpresponse.Conflict(err.Error()))
	}

	response := httpresponse.NewSuccessResponse("Tenant updated successfully", update)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) FindAll(c *echo.Context) error {
	tenants, err := h.service.FindAll()

	if err != nil {
		return c.JSON(http.StatusNotFound, httpresponse.NotFound(err.Error()))
	}

	response := httpresponse.NewSuccessResponse("Tenants retrieved successfully", tenants)

	return c.JSON(http.StatusOK, response)
}

func (h *handler) Delete(c *echo.Context) error {
	idStr := c.Param("id")

	id64, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.BadRequest("invalid_request_body"))
	}

	tenant, err := h.service.Delete(uint(id64))

	if err != nil {
		return c.JSON(http.StatusConflict, httpresponse.Conflict(err.Error()))
	}

	response := httpresponse.NewSuccessResponse("Tenant deleted successfully", tenant)

	return c.JSON(http.StatusOK, response)
}