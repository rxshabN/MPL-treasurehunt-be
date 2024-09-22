package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Oik17/MPL-treasurehunt-be/internal/models"
	"github.com/Oik17/MPL-treasurehunt-be/internal/services"
)

func CheckAnswer(c *fiber.Ctx) error {
	var req struct {
		QuestionNumber int    `json:"question_number"`
		Answer         string `json:"answer"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	isCorrect, err := services.CheckAnswer(req.QuestionNumber, req.Answer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to check answer",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"correct": isCorrect,
	})
}

func CreateQuestion(c *fiber.Ctx) error {
	var req models.Questions

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	err := services.CreateQuestion(req.QuestionNumber, req.Question, req.Answer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create question",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Question created successfully",
	})
}
