package api

import (
	"strconv"

	"github.com/daksh-pareek/paperman/src/controllers"
	_ "github.com/daksh-pareek/paperman/src/docs"
	"github.com/daksh-pareek/paperman/src/dtos"
	"github.com/daksh-pareek/paperman/src/repositories"
	"github.com/daksh-pareek/paperman/src/server/parsers"
	"github.com/daksh-pareek/paperman/src/server/validators"
	"github.com/gofiber/fiber/v2"
)

var quizController *controllers.QuizController

func QuizzesRoutes(dbInstance repositories.QuizRepository) func(router fiber.Router) {
	// initialize Quiz Controller
	quizController = controllers.CreateQuizController(dbInstance)
	return func(router fiber.Router) {
		router.Get("/:id", GetQuiz)
		router.Post("/", CreateQuiz)
	}
}

// GetQuiz
//
//	@Summary	Get a quiz by id
//	@Tags		Quizzes
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"Quiz ID"
//	@Success	200	{object}	dtos.QuizDTO
//	@Failure	400	{object}	dtos.ErrorResponse
//	@Router		/quiz/{id} [get]
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

// CreateQuiz
//
// @Summary	Create a quiz
// @Tags		Quizzes
// @Accept		json
// @Produce	json
// @Param		quiz	body		dtos.QuizDTO	true	"Quiz object"
// @Success	200	{object}	dtos.QuizDTO
// @Failure	400	{object}	dtos.ErrorResponse
// @Router		/quiz [post]
func CreateQuiz(ctx *fiber.Ctx) error {
	// pares request body to QuizDTO
	parsedBody, parseError := parsers.ParseBody[dtos.QuizDTO](ctx)
	if parseError != nil {
		return parsers.SendParsingError(ctx, parseError)
	}
	// validate parsed body
	validationError := validators.ValidateCreateQuiz(parsedBody)
	if validationError != nil {
		return validators.SendValidationError(ctx, validationError)
	}
	// create quiz
	createdQuiz, err := quizController.BuildQuiz(parsedBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dtos.CreateErrorResponse(fiber.StatusBadRequest, "Error creating quiz"))
	}
	return ctx.Status(fiber.StatusOK).JSON(createdQuiz)
}
