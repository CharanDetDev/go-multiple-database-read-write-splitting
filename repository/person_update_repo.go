package repository

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/database"
)

func (repo *personRepo) UpdatePerson(newPerson *model.Person) error {

	result := database.Conn.Model(&newPerson).Where("PersonID = ?", newPerson.PersonID).Updates(newPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
