package dtos

type Option struct {
	ID   uint64
	Text string `json:"text"`
	// Add additional properties as needed (e.g., value, isCorrect)
}

// QuestionType defines the types of questions
type QuestionType string

const (
	MultipleChoice QuestionType = "multipleChoice"
	// Add other question types as needed
)

// QuestionDTO represents a question data transfer object
type QuestionDTO struct {
	ID            uint64
	Text          string
	Type          QuestionType
	Options       []Option
	CorrectAnswer *Option
}
