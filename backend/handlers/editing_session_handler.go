package handlers

import (
	"fmt"
	"go-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Retrieve all Editing Sessions.
// @Summary Retrieve all Editing Sessions.
// @Description Retrieve all Editing Sessions stored.
// @Tags EditingSessions
// @Produce json
// @Success 200 {array} models.EditingSession
// @Router /editing-sessions [get]
func GetAllEditingSessions(c *fiber.Ctx) error {
	return c.JSON(models.MockedEditingSessionsTable)
}

// Create a new EditingSession.
// @Summary Create a new EditingSession.
// @Description Create a new EditingSession and return
// @Description such EditingSession encoded in JSON.
// @Tags EditingSessions
// @Accept json
// @Produce json
// @Param editing_session body models.EditingSession true "EditingSession"
// @Success 200 {object} models.EditingSession
// @Failure 400
// @Router /editing-sessions [post]
func CreateEditingSession(c *fiber.Ctx) error {
	var editingSession models.EditingSession

	err := c.BodyParser(&editingSession)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid editing session data",
		})
	}

	details := make([]string, 0)
	_, err = models.GetUserById(editingSession.UserId)
	if err != nil {
		details = append(details, "User not found")
	}
	_, err = models.GetDocumentById(editingSession.DocumentId)
	if err != nil {
		details = append(details, "Document not found")
	}

	if len(details) > 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"details": details,
		})
	}

	models.MockedEditingSessionsTable = append(models.MockedEditingSessionsTable,
		&editingSession)

	return c.JSON(editingSession)
}

// Write in a EditingSession in its current position.
// @Summary Write in new EditingSession.
// @Description Write bytes in a EditingSession in its current position.
// @Tags EditingSessions
// @Accept plain
// @Produce json
// @Param user_id path int true "User Id"
// @Param document_id path int true "Document Id"
// @Success 204
// @Failure 400
// @Router /editing-sessions/{user_id}/{document_id} [post]
func WriteInEditingSession(c *fiber.Ctx) error {
	errorDetails := []string{}

	userId, err := c.ParamsInt("user_id")
	if err != nil {
		errorDetails = append(
			errorDetails,
			"invalid type: id of User should be an integer",
		)
	}
	documentId, err := c.ParamsInt("document_id")
	if err != nil {
		errorDetails = append(
			errorDetails,
			"invalid type: id of Document should be an integer",
		)
	}

	if len(errorDetails) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"details": errorDetails,
		})
	}

	editingSession, err := models.GetEditingSessionByUserIdAndDocumentId(userId,
		documentId)

	contentToBeWritten := fmt.Sprintf("%s", c.Body())

	err = editingSession.WriteToDocument(contentToBeWritten)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Delete a given number characters in a EditingSession in its current position.
// @Summary Delete chars in a EditingSession.
// @Description Delete a given number of chars in a EditingSession in its current
// @Description position.
// @Tags EditingSessions
// @Param user_id path int true "User Id"
// @Param document_id path int true "Document Id"
// @Success 204
// @Failure 400
// @Router /editing-sessions/{user_id}/{document_id}/{number_of_chars} [delete]
func DeleteInEditingSession(c *fiber.Ctx) error {
	errorDetails := []string{}

	userId, err := c.ParamsInt("user_id")
	if err != nil {
		errorDetails = append(
			errorDetails,
			"invalid type: id of User should be an integer",
		)
	}
	documentId, err := c.ParamsInt("document_id")
	if err != nil {
		errorDetails = append(
			errorDetails,
			"invalid type: id of Document should be an integer",
		)
	}

	numberOfChars, err := c.ParamsInt("number_of_chars")
	if err != nil {
		errorDetails = append(
			errorDetails,
			"invalid type: number of chars to be deleted should be an integer",
		)
	}

	if len(errorDetails) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"details": errorDetails,
		})
	}

	editingSession, err := models.GetEditingSessionByUserIdAndDocumentId(userId,
		documentId)

	err = editingSession.DeleteFromDocument(numberOfChars)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Update the activity status or current position of an EditingSession.
// @Summary Update an EditingSession.
// @Description Update the activity status or current position of an EditingSession.
// @Tags EditingSessions
// @Accept json
// @Produce json
// @Param user_id path int true "User Id"
// @Param document_id path int true "Document Id"
// @Param editing_session body models.EditingSession true "EditingSession"
// @Success 200 {object} models.EditingSession
// @Failure 400
// @Router /editing-sessions/{user_id}/{document_id} [put]
func UpdateEditingSession(c *fiber.Ctx) error {
	errorDetails := []string{}

	userId, err := c.ParamsInt("user_id")
	if err != nil {
		errorDetails = append(
			errorDetails,
			"invalid type: id of User should be an integer",
		)
	}
	documentId, err := c.ParamsInt("document_id")
	if err != nil {
		errorDetails = append(
			errorDetails,
			"invalid type: id of Document should be an integer",
		)
	}

	if len(errorDetails) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"details": errorDetails,
		})
	}

	editingSession, err := models.GetEditingSessionByUserIdAndDocumentId(userId,
		documentId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	var newData models.EditingSession
	err = c.BodyParser(&newData)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid body input data",
		})
	}

	editingSession.CurrentPosition = newData.CurrentPosition
	editingSession.IsEditingActive = newData.IsEditingActive

	return c.Status(fiber.StatusOK).JSON(editingSession)
}
