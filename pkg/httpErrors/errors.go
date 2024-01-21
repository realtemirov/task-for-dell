package httpErrors

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/realtemirov/task-for-dell/pkg/logger"
)

var (
	BadRequest       string = "BAD_REQUEST"
	NotFound         string = "NOT_FOUND"
	NotRequiredField string = "NOT_REQUIRED_FIELD"
	BadQueryParams   string = "BAD_QUERY_PARAMS"
	RequestTimeOut   string = "REQUEST_TIMEOUT"
	InternalServer   string = "INTERNAL_SERVER_ERROR"
)

type ErrorMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// NewErrorMessage
func NewErrorMessage(message string, code int) (int, *ErrorMessage) {
	return code, &ErrorMessage{
		Message:    message,
		StatusCode: code,
	}
}

func ErrResponse(err error) (int, *ErrorMessage) {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewErrorMessage(NotFound, http.StatusNotFound)
	case errors.Is(err, context.DeadlineExceeded):
		return NewErrorMessage(RequestTimeOut, http.StatusRequestTimeout)
	case strings.Contains(err.Error(), "SQLSTATE"):
		return NewErrorMessage(BadRequest, http.StatusBadRequest)
	case strings.Contains(err.Error(), "Field validation"):
		return NewErrorMessage(BadRequest, http.StatusBadRequest)
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewErrorMessage(BadRequest, http.StatusBadRequest)
	default:
		return NewErrorMessage(InternalServer, http.StatusInternalServerError)
	}
}

func ErrResponseWithLog(ctx echo.Context, logger logger.Logger, err error) error {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
	return ctx.JSON(ErrResponse(err))
}

// Get request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// Get user ip address
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}
