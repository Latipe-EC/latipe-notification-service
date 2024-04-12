package responses

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Error struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func New(msg string) error {
	return errors.New(msg)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	msg := DefaultError

	var (
		customErr *Error
		fiberErr  *fiber.Error
	)

	switch {
	// trieve the custom status code if it's an fiber.*Error
	case errors.As(err, &fiberErr):
		msg.Code = fiberErr.Code
		msg.Message = fiberErr.Message
		// TODO: handle fiber errors
	case errors.As(err, &customErr):
		msg.Code = customErr.Code
		msg.Message = customErr.Message

	default:
		msg.Code = http.StatusInternalServerError
		msg.Message = "Internal server error"
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return msg.JSON(ctx)
}
