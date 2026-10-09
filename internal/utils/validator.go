package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type CustomValidator struct {
	validate *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		if name != "" {
			return name
		}
		return field.Name
	})

	return &CustomValidator{
		validate: v,
	}
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validate.Struct(i)
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors"`
}

func FormatValidationErrors(err error) []FieldError {
	var errors []FieldError

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return []FieldError{
			{
				Field:   "body",
				Message: "Invalid request",
			},
		}
	}

	for _, fieldErr := range validationErrors {
		errors = append(errors, FieldError{
			Field:   fieldErr.Field(),
			Message: validationMessage(fieldErr),
		})
	}

	return errors
}

func validationMessage(fe validator.FieldError) string {
	field := strings.ReplaceAll(fe.Field(), "_", " ")

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s harus diisi", field)
	case "min":
		return fmt.Sprintf("%s tidak boleh kurang dari %s karakter", field, fe.Param())
	case "max":
		return fmt.Sprintf("%s tidak boleh lebih dari %s karakter", field, fe.Param())
	case "gte":
		return fmt.Sprintf("%s harus lebih atau sama dengan %s", field, fe.Param())
	case "lte":
		return fmt.Sprintf("%s harus kurang atau sama dengan %s", field, fe.Param())
	case "e164":
		return "nomor tidak sesuai standard e164 atau tidak valid"
	default:
		return fmt.Sprintf("%s tidak valid", field)
	}
}

func ValidationErrorHandler(c *echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, ValidationErrorResponse{
		Message: "invalid",
		Errors:  FormatValidationErrors(err),
	})
}
