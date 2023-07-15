package handlers

import (
	"go-backend/models"
	"log"

	"github.com/go-playground/validator/v10"
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
			foundUser = *user
		}
	}
	return c.JSON(foundUser)
}

// Create a new User.
// @Summary Create a new User.
// @Description Create a new User discarting the value for "id" and return
// @Description such User encoded in JSON.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "invalid user data",
		})
	}

	models.MockedUsersTable = append(models.MockedUsersTable, &user)

	validate := validator.New()

	err = validate.Struct(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	user.Id = models.NextUserId()

	return c.JSON(user)
}
