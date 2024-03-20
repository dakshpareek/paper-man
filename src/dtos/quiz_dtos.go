package dtos

type QuizDTO struct {
	ID        uint64
	Title     string
	CreatorId uint64
	Questions []QuestionDTO
}
