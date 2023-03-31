package handler

import (
	"strconv"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/caller"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (personHandler *personHandler) GetPersonWithPersonID(c *fiber.Ctx) error {

	personId, err := strconv.Atoi(c.Params("personId"))
	if err != nil {
		return caller.Response(c, fiber.StatusBadRequest, "Invalid param")
	}

	var resPerson model.Person
	err = personHandler.PersonService.GetPersonWithPersonID(personId, &resPerson)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return caller.Response(c, fiber.StatusOK, "Not found")
		} else {
			return caller.Response(c, fiber.StatusInternalServerError, "Internal server error")
		}
	}

	return caller.Response(c, fiber.StatusOK, resPerson)
}
