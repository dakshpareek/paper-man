package repositories

import "github.com/daksh-pareek/paperman/src/dtos"

type QuizRepository interface {
	CreateQuiz(quiz *dtos.QuizDTO) error
}
