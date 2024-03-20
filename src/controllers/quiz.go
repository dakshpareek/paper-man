package controllers

import (
	"github.com/daksh-pareek/paperman/src/dtos"
	"github.com/daksh-pareek/paperman/src/repositories"
)

type QuizController struct {
	db repositories.QuizRepository
}

func CreateQuizController(db repositories.QuizRepository) *QuizController {
	return &QuizController{db: db}
}

func (qc *QuizController) BuildQuiz(quiz *dtos.QuizDTO) (*dtos.QuizDTO, error) {
	// Store quiz in the database using repository
	createdQuiz, err := qc.db.CreateQuiz(quiz)
	return createdQuiz, err
}

func (qc *QuizController) GetQuiz(quizId uint64) (*dtos.QuizDTO, error) {
	quiz, err := qc.db.GetQuiz(quizId)
	return quiz, err
}

// List all the methods I should be having for quizzes
// CreateQuiz
// GetQuiz
// UpdateQuiz
// DeleteQuiz
// ListQuizzes
// ListQuizzesByCreator
