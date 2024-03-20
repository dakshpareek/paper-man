package validators

import (
	"github.com/daksh-pareek/paperman/src/dtos"
	"github.com/gofiber/fiber/v2"
)

func ValidateCreateQuiz(quiz *dtos.QuizDTO) *ValidationError {
	if quiz.Title == "" {
		return &ValidationError{
			status:  fiber.StatusUnprocessableEntity,
			message: "Title cannot be empty",
		}
	}
	if quiz.CreatorId == 0 {
		return &ValidationError{
			status:  fiber.StatusUnprocessableEntity,
			message: "Creator id cannot be empty",
		}
	}
	return nil
}
