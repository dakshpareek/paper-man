package controllers

type QuizController struct {
}

type Quiz struct {
	userId uint64
}

func (qc *QuizController) BuildQuiz(userId uint64) (*Quiz, error) {
	return &Quiz{userId: userId}, nil
}
