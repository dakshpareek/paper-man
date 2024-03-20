package api

import (
	"strconv"

	"github.com/daksh-pareek/paperman/src/controllers"
	"github.com/daksh-pareek/paperman/src/dtos"
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
	quizId := ctx.Params("id")
	// convert string to uint64
	quizIdUint, err := strconv.ParseUint(quizId, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.CreateErrorResponse(fiber.StatusBadRequest, "Error parsing quiz id"))
	}

	quiz, err := quizController.GetQuiz(quizIdUint)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.CreateErrorResponse(fiber.StatusBadRequest, "Error getting quiz"))
	}
	return ctx.Status(fiber.StatusOK).JSON(quiz)
}
