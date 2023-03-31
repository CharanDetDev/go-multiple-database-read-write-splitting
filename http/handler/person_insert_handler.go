package handler

import (
	"encoding/json"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/caller"
	"github.com/gofiber/fiber/v2"
)

func (personHandler *personHandler) InsertPerson(c *fiber.Ctx) error {

	var newPerson model.PersonRequest
	err := json.Unmarshal(c.Body(), &newPerson)
	if err != nil {
		return caller.Response(c, fiber.StatusBadRequest, err.Error())
	}

	if err := personHandler.PersonService.InsertPerson(&newPerson); err != nil {
		return caller.Response(c, fiber.StatusInternalServerError, err.Error())
	}

	return caller.Response(c, fiber.StatusOK, "Insert completed successfully")
}
