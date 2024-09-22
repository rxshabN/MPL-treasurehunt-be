package models

type Questions struct {
	QuestionNumber int    `db:"question_number" json:"question_number"`
	Question       string `db:"question" json:"question"`
	Answer         string `db:"answer" json:"answer"`
}
