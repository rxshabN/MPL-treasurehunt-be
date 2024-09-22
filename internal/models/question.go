package models

type Questions struct {
	QuestionNumber int    `json:"question_number"`
	Question       string `json:"question"`
	Answer         string `json:"answer"`
}
