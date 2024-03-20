package parsers

import (
	"github.com/daksh-pareek/paperman/src/dtos"
	"github.com/gofiber/fiber/v2"
)

// Parsing error is a custom error type that can be used to send parsing errors to the client
type ParsingError struct {
	status  int
	message string
}

// Error returns the error message
func (e *ParsingError) Error() string {
	return e.message
}

func (e *ParsingError) ErrorDetails() (int, string) {
	return e.status, e.message
}

// NewParsingError creates a new ParsingError with the given status code and message
func NewParsingError(status int, message string) *ParsingError {
	return &ParsingError{status: status, message: message}
}

func SendParsingError(ctx *fiber.Ctx, err *ParsingError) error {
	return ctx.Status(err.status).JSON(dtos.CreateErrorResponse(err.ErrorDetails()))
}

func ParseBody[T any](ctx *fiber.Ctx) (*T, *ParsingError) {
	var body = new(T)
	if err := ctx.BodyParser(body); err != nil {
		return nil, NewParsingError(fiber.StatusBadRequest, "Invalid request body")
	}
	return body, nil
}
