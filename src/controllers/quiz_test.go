package controllers

import "testing"

func TestQuiz(t *testing.T) {
	// Create first TestQuiz function
	t.Run("TestQuiz", func(t *testing.T) {
		// create a instance of quiz controller
		quizController := QuizController{}
		quiz, _ := quizController.BuildQuiz(1)
		// assert if the quiz is not nil and has userId
		if quiz == nil || quiz.userId != 1 {
			t.Errorf("Quiz is not created")
		}
	})
}
