package handlers

import (
	"go-backend/models"

	"github.com/gofiber/fiber/v2"
)

// Retrieve an User by its `id`.
// @Summary Retrieve an User by its `id`.
// @Description	Retrieve an User by its `ìd`, if there is no User with given `ìd`, then
// @Description	return with status code 404 indicating it, if `id` is not an
// @Description integer, then return with status code 400 incidating that the given value
// @Description has an invalid type and the required type.
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400
// @Failure 404
// @Router /users/{id} [get]
func GetUserByIdHandler(c *fiber.Ctx) error {
	if c.Params("id") == "" {
		return c.JSON(models.MockedUsersTable)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid type: id of User should be an integer",
		})
	}

	var foundUser models.User
	for _, user := range models.MockedUsersTable {
		if id == user.Id {
			foundUser = user
		}
	}
	return c.JSON(foundUser)
}
