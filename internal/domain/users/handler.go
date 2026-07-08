package users

import (
	"mms-dbsd/internal/domain/users/dto"
	"mms-dbsd/internal/httpresponse"
	"mms-dbsd/internal/validator"
	"net/http"

	"github.com/labstack/echo/v5"
)

type userHandler struct {
	service IRegisterService
}

func NewUserHandler(service IRegisterService) *userHandler {
	return &userHandler{
		service: service,
	}
}

func (h *userHandler) Register(c *echo.Context) error {
	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.BadRequest("invalid_request_body"))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ValidationError(validator.FormatValidationError(err)))
	}

	user, err := h.service.Register(&req)
	if err != nil {
		return c.JSON(http.StatusConflict, httpresponse.Conflict(err.Error()))
	}

	response := httpresponse.NewSuccessResponse("User registered successfully", user)
	return c.JSON(http.StatusOK, response)
}