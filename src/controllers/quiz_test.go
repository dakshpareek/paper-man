package controllers

import (
	"testing"

	"github.com/daksh-pareek/paperman/src/dtos"
)

// Implement repository interface for testing
type QuizRepositoryMock struct {
}

func (qrm *QuizRepositoryMock) CreateQuiz(quiz *dtos.QuizDTO) (*dtos.QuizDTO, error) {
	return quiz, nil
}

func (qrm *QuizRepositoryMock) GetQuiz(quizId uint64) (*dtos.QuizDTO, error) {
	if quizId != 1 {
		return nil, nil
	}
	quiz := dtos.QuizDTO{
		ID:        1,
		Title:     "Test Quiz",
		CreatorId: 101,
	}
	return &quiz, nil
}

func TestBuildQuiz(t *testing.T) {
	// Create instance of repository mock
	QuizRepositoryMock := QuizRepositoryMock{}
	// Create a quiz controller
	// quizController := QuizController{db: &QuizRepositoryMock}
	quizController := CreateQuizController(&QuizRepositoryMock)
	// Create first TestQuiz function
	t.Run("Create a quiz without questions", func(t *testing.T) {
		quizDTO := dtos.QuizDTO{
			Title:     "Test Quiz",
			CreatorId: 101,
		}
		quiz, _ := quizController.BuildQuiz(&quizDTO)
		if quiz == nil {
			t.Errorf("Quiz not created")
		}
		if quiz.ID == 1 {
			t.Errorf("Quiz id is not correct")
		}
		if quiz.Title != "Test Quiz" {
			t.Errorf("Quiz title is not correct")
		}
		if quiz.CreatorId != 101 {
			t.Errorf("Quiz creator id is not correct")
		}
		if len(quiz.Questions) != 0 {
			t.Errorf("Quiz questions are not correct")
		}
	})

	t.Run("Create a quiz with one question and correct answer", func(t *testing.T) {
		quizDTO := dtos.QuizDTO{
			Title:     "Test Quiz",
			CreatorId: 101,
		}
		question1 := dtos.QuestionDTO{
			ID:   1,
			Text: "What is 2+2?",
			Type: dtos.MultipleChoice,
			Options: []dtos.Option{
				{ID: 1, Text: "1"},
				{ID: 2, Text: "2"},
				{ID: 3, Text: "3"},
				{ID: 4, Text: "4"},
			},
			CorrectAnswer: &dtos.Option{ID: 4, Text: "4"},
		}
		quizDTO.Questions = append(quizDTO.Questions, question1)
		quiz, _ := quizController.BuildQuiz(&quizDTO)
		if quiz == nil {
			t.Errorf("Quiz not created")
		}
		if quiz.ID == 1 {
			t.Errorf("Quiz id is not correct")
		}
		if quiz.Title != "Test Quiz" {
			t.Errorf("Quiz title is not correct")
		}
		if quiz.CreatorId != 101 {
			t.Errorf("Quiz creator id is not correct")
		}
		if len(quiz.Questions) != 1 {
			t.Errorf("Quiz questions are not correct")
		}
		if quiz.Questions[0].ID != 1 {
			t.Errorf("Question id is not correct")
		}
		if quiz.Questions[0].Text != "What is 2+2?" {
			t.Errorf("Question text is not correct")
		}
		if quiz.Questions[0].Type != dtos.MultipleChoice {
			t.Errorf("Question type is not correct")
		}
		if len(quiz.Questions[0].Options) != 4 {
			t.Errorf("Question options are not correct")
		}
		if quiz.Questions[0].CorrectAnswer == nil {
			t.Errorf("Question correct answer is not correct")
		}
		if quiz.Questions[0].CorrectAnswer.ID != 4 {
			t.Errorf("Question correct answer id is not correct")
		}
		if quiz.Questions[0].CorrectAnswer.Text != "4" {
			t.Errorf("Question correct answer text is not correct")
		}
	})
}

func TestGetQuiz(t *testing.T) {
	// Create instance of repository mock
	QuizRepositoryMock := QuizRepositoryMock{}
	// Create a quiz controller
	quizController := QuizController{db: &QuizRepositoryMock}

	t.Run("Quiz not found", func(t *testing.T) {
		quiz, _ := quizController.GetQuiz(0)
		if quiz != nil {
			t.Errorf("Quiz found")
		}
	})

	// Create first TestGetQuiz function
	t.Run("Get a quiz", func(t *testing.T) {
		quiz, _ := quizController.GetQuiz(1)
		if quiz == nil {
			t.Errorf("Quiz not found")
		}
		if quiz.ID != 1 {
			t.Errorf("Quiz id is not correct")
		}
		if quiz.Title != "Test Quiz" {
			t.Errorf("Quiz title is not correct")
		}
		if quiz.CreatorId != 101 {
			t.Errorf("Quiz creator id is not correct")
		}
	})
}
