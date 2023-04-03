package repository

import "github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"

func (repo *personRepo) DeletePerson(personID int) error {

	tx := repo.DatabaseConn.Begin()
	defer func() {
		if recovery := recover(); recovery != nil {
			tx.Rollback()
		}
	}()

	if result := tx.Where("PersonID = ?", personID).Delete(&model.PersonModel{}); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}
