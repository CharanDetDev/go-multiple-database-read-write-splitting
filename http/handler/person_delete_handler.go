package handler

import (
	"strconv"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/caller"

	"github.com/gofiber/fiber/v2"
)

func (personHandler *personHandler) DeletePerson(c *fiber.Ctx) error {

	personId, err := strconv.Atoi(c.Params("personId"))
	if err != nil {
		return caller.Response(c, fiber.StatusBadRequest, err.Error())
	}

	if err := personHandler.PersonService.DeletePerson(personId); err != nil {
		return caller.Response(c, fiber.StatusInternalServerError, err.Error())
	}

	return caller.Response(c, fiber.StatusOK, "Delete completed successfully")
}
