package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type CustomValidator struct {
	validate *validator.Validate
}

func NewValidator() *CustomValidator{
	return &CustomValidator{
		validate: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i any) error{
	if err :=cv.validate.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}