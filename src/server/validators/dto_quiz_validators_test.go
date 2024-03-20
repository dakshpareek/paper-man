package validators

// Write test to validate ValidateCreateQuiz function

import (
	"testing"

	"github.com/daksh-pareek/paperman/src/dtos"
)

func TestValidateCreateQuiz(t *testing.T) {
	t.Run("Should throw if title missing", func(t *testing.T) {
		quiz := &dtos.QuizDTO{
			CreatorId: 101,
		}
		err := ValidateCreateQuiz(quiz)
		if err == nil {
			t.Errorf("Expected error")
		}
		if err.status != 422 {
			t.Errorf("Expected status 422")
		}
		if err.message != "Title cannot be empty" {
			t.Errorf("Expected message Title cannot be empty")
		}
	})
	t.Run("Should throw if creator id missing", func(t *testing.T) {
		quiz := &dtos.QuizDTO{
			Title: "Test Quiz",
		}
		err := ValidateCreateQuiz(quiz)
		if err == nil {
			t.Errorf("Expected error")
		}
		if err.status != 422 {
			t.Errorf("Expected status 422")
		}
		if err.message != "Creator id cannot be empty" {
			t.Errorf("Expected message Creator id cannot be empty")
		}
	})
	t.Run("Should not throw if title and creator id present", func(t *testing.T) {
		quiz := &dtos.QuizDTO{
			Title:     "Test Quiz",
			CreatorId: 101,
		}
		err := ValidateCreateQuiz(quiz)
		if err != nil {
			t.Errorf("Expected no error")
		}
	})
}
