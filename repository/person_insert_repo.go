package repository

import "github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"

func (repo *personRepo) InsertPerson(newPerson *model.Person) error {

	result := repo.DatabaseConn.Create(newPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
