package repositories

import "github.com/daksh-pareek/paperman/src/dtos"

type QuizRepository interface {
	CreateQuiz(quiz *dtos.QuizDTO) (*dtos.QuizDTO, error)
	GetQuiz(quizId uint64) (*dtos.QuizDTO, error)
}
