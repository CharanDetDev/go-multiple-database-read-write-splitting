package repository

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
)

func (repo *personRepo) UpdatePerson(newPerson *model.PersonModel) error {

	tx := repo.DatabaseConn.Begin()
	defer func() {
		if recovery := recover(); recovery != nil {
			tx.Rollback()
		}
	}()

	if result := tx.Model(&newPerson).Where("PersonID = ?", newPerson.PersonID).Updates(newPerson); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}
