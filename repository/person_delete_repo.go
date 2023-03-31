package repository

import "github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"

func (repo *personRepo) DeletePerson(personID int) error {

	result := repo.DatabaseConn.Where("PersonID = ?", personID).Delete(&model.Person{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
