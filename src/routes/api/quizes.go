package api

import (
	"github.com/daksh-pareek/paperman/src/controllers"
	"github.com/daksh-pareek/paperman/src/repositories"
	"github.com/gofiber/fiber/v2"
)

var quizController *controllers.QuizController

func QuizzesRoutes(dbInstance repositories.QuizRepository) func(router fiber.Router) {
	// initialize Quiz Controller
	quizController = controllers.CreateQuizController(dbInstance)
	return func(router fiber.Router) {
		router.Get("/:id", GetQuiz)
	}
}

func GetQuiz(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON("Dummy Quiz Data")
}
