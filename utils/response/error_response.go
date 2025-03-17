package response

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hudaputrasantosacommon-package-microapp/utils/validation"
)

func ErrorMessage(c *fiber.Ctx, statusCode int, message string, err error, data ...interface{}) error {
	response := &Response{
		Error:   true,
		Message: message,
	}

	if err != nil {
		response.ErrorDetails = fmt.Sprintf("%v", err)
	}

	if data != nil {
		response.Data = data
	}

	return c.Status(statusCode).JSON(response)
}

func ErrorValidationMessage(c *fiber.Ctx, statusCode int, dataValidation []*validation.ErrorResponse) error {
	response := &Response{
		Error:   true,
		Message: "Validation Error",
		Data:    dataValidation,
	}

	return c.Status(statusCode).JSON(response)
}

func ErrorMessageDetail(c *fiber.Ctx, statusCode int, errorCode int, errorType, message, details string, err error) error {
	response := &Response{
		Error:        true,
		Message:      message,
		ErrorCode:    errorCode,
		ErrorType:    errorType,
		ErrorDetails: fmt.Sprintf("%s : %v", details, err),
	}
	return c.Status(statusCode).JSON(response)
}
