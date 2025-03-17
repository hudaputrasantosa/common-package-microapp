package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Error        bool        `json:"error,omitempty"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    int         `json:"error_code,omitempty"`
	ErrorType    string      `json:"error_type,omitempty"`
	ErrorDetails string      `json:"error_details,omitempty"`
}

func SuccessMessage(c *fiber.Ctx, statusCode int, message string) error {
	response := &Response{
		Error:   false,
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

func SuccessMessageWithData(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := &Response{
		Error:   false,
		Message: message,
		Data:    data,
	}
	return c.Status(statusCode).JSON(response)
}

// func RespondWithPagination(c *fiber.Ctx, code int, message string, total int, page int, perPage int, dataName string, data interface{}) error {
// 	return c.Status(code).JSON(fiber.Map{
// 		"error":   false,
// 		"message": message,
// 		"data": fiber.Map{
// 			dataName:   data,
// 			"total":    total,
// 			"page":     page,
// 			"per_page": perPage,
// 		},
// 	})
// }
