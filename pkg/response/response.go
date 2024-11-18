package response

import (
	"bytes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/raxmet/pkg/logger"
)

type defaultResponse struct {
	Success  bool        `json:"success"`
	Status   int         `json:"status"`
	Code     string      `json:"code,omitempty"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	Paginate *Pagination `json:"paginate,omitempty"`
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"description"`
	Message          string `json:"message"`
}

func JsonSuccess(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(defaultResponse{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "OK",
		Data:    data,
	})
}

func JsonPagination(ctx *fiber.Ctx, pagination *Pagination) error {
	return ctx.Status(fiber.StatusOK).JSON(defaultResponse{
		Success:  true,
		Status:   fiber.StatusOK,
		Message:  "OK",
		Data:     pagination.Rows,
		Paginate: pagination,
	})
}

func JsonError(ctx *fiber.Ctx, errs []error, code string) error {
	errorMessage := logErrorFormat(errs, code)
	logger.Logger.Error(errorMessage)
	return ctx.Status(fiber.StatusBadRequest).JSON(defaultResponse{
		Success: false,
		Status:  fiber.StatusBadRequest,
		Code:    code,
		Message: errs[0].Error(),
		Data:    errorsToStrings(errs),
	})
}

func JsonErrorInternal(ctx *fiber.Ctx, errs []error, code string) error {
	errorMessage := logErrorFormat(errs, code)
	logger.Logger.Error(errorMessage)
	return ctx.Status(fiber.StatusInternalServerError).JSON(defaultResponse{
		Success: false,
		Status:  fiber.StatusInternalServerError,
		Code:    code,
		Message: errs[0].Error(),
		Data:    errorsToStrings(errs),
	})
}

func JsonErrorValidation(ctx *fiber.Ctx, errs []error) error {
	errorMessage := logErrorFormat(errs, "E_VALIDATION")
	logger.Logger.Error(errorMessage)
	return ctx.Status(fiber.StatusBadRequest).JSON(defaultResponse{
		Success: false,
		Status:  fiber.StatusBadRequest,
		Code:    "E_VALIDATION",
		Message: errs[0].Error(),
		Data:    errorsToStrings(errs),
	})
}

func JsonErrorNotFound(ctx *fiber.Ctx, errs []error) error {
	errorMessage := logErrorFormat(errs, "E_NOT_FOUND")
	logger.Logger.Error(errorMessage)
	return ctx.Status(fiber.StatusNotFound).JSON(defaultResponse{
		Success: false,
		Status:  fiber.StatusNotFound,
		Code:    "E_NOT_FOUND",
		Message: errs[0].Error(),
		Data:    errorsToStrings(errs),
	})
}

func JsonErrorUnauthorized(ctx *fiber.Ctx, errs []error) error {
	errorMessage := logErrorFormat(errs, "E_UNAUTHORIZED")
	logger.Logger.Error(errorMessage)
	return ctx.Status(fiber.StatusUnauthorized).JSON(defaultResponse{
		Success: false,
		Status:  fiber.StatusUnauthorized,
		Code:    "E_UNAUTHORIZED",
		Message: errs[0].Error(),
		Data:    errorsToStrings(errs),
	})
}

func JsonErrorEnvironment(ctx *fiber.Ctx, env string) error {
	err := errors.New("missing env " + env + " variable")
	errorMessage := logErrorFormat([]error{err}, "E_ENV")
	logger.Logger.Error(errorMessage)
	return ctx.Status(fiber.StatusInternalServerError).JSON(defaultResponse{
		Success: false,
		Status:  fiber.StatusInternalServerError,
		Code:    "E_ENV",
		Message: err.Error(),
		Data:    []string{err.Error()},
	})
}

func JsonErrorForbidden(ctx *fiber.Ctx, errs []error) error {
	errorMessage := logErrorFormat(errs, "E_FORBIDDEN")
	logger.Logger.Error(errorMessage)
	return ctx.Status(fiber.StatusForbidden).JSON(defaultResponse{
		Success: false,
		Status:  fiber.StatusForbidden,
		Code:    "E_FORBIDDEN",
		Message: errs[0].Error(),
		Data:    errorsToStrings(errs),
	})
}

func logErrorFormat(errs []error, code string) string {
	var buffer bytes.Buffer
	for _, err := range errs {
		buffer.WriteString(err.Error())
	}
	return "❌ " + "[" + code + "] " + buffer.String()
}

func errorsToStrings(errs []error) interface{} {
	var errsString []string
	for _, err := range errs {
		errsString = append(errsString, err.Error())
	}
	return fiber.Map{
		"errors": errsString,
	}
}
