package handlers

import (
	"go-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Retrieve all Documents.
// @Summary Retrieve all Documents.
// @Description Retrieve all Documents stored.
// @Tags Documents
// @Produce json
// @Success 200 {array} models.Document
// @Router /documents [get]
func GetAllDocuments(c *fiber.Ctx) error {
	return c.JSON(models.MockedDocumentsTable)
}

// Retrieve a Document by its `id`.
// @Summary Retrieve a Document by its `id`.
// @Description	Retrieve an User by its `ìd`, if there is no User with given `ìd`, then
// @Description	return with status code 404 indicating it, if `id` is not an
// @Description integer, then return with status code 400 incidating that the given value
// @Description has an invalid type and the required type.
// @Tags Documents
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400
// @Failure 404
// @Router /documents/{id} [get]
func GetDocumentByIdHandler(c *fiber.Ctx) error {
	if c.Params("id") == "" {
		return c.JSON(models.MockedDocumentsTable)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid type: id of user should be an integer",
		})
	}

	var foundDocument models.Document
	for _, document := range models.MockedDocumentsTable {
		if id == document.Id {
			foundDocument = *document
		}
	}
	return c.JSON(foundDocument)
}

// Create a new Document.
// @Summary Create a new Document.
// @Description Create a new Document with an optional initial textual content and return
// @Description such Document encoded in JSON.
// @Tags Documents
// @Accept json
// @Produce json
// @Param document body models.Document true "Document"
// @Success 200 {object} models.Document
// @Failure 400
// @Router /documents [post]
func CreateDocument(c *fiber.Ctx) error {
	var document models.Document

	err := c.BodyParser(&document)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid document data",
		})
	}

	models.MockedDocumentsTable = append(models.MockedDocumentsTable, &document)

	return c.JSON(document)
}
