package services

import (
	"database/sql"
	"log"

	"github.com/Oik17/MPL-treasurehunt-be/internal/database"
	"github.com/Oik17/MPL-treasurehunt-be/internal/models"
)

func CheckAnswer(questionNumber int, userAnswer string) (bool, error) {
	db := database.DB.Db
	var question models.Questions

	query := `SELECT question_number, question, answer FROM questions WHERE question_number = $1`
	err := db.Get(&question, query, questionNumber)

	if err == sql.ErrNoRows {
		log.Printf("Question %d not found", questionNumber)
		return false, nil
	} else if err != nil {
		log.Printf("Database error: %v", err)
		return false, err
	}

	isCorrect := question.Answer == userAnswer
	return isCorrect, nil
}
