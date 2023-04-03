package repository

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
	"gorm.io/plugin/dbresolver"
)

func (repo *personRepo) GetPersonWithPersonID(personId int, person *model.PersonModel) error {

	err := repo.DatabaseConn.Clauses(dbresolver.Write).First(&person, personId).Error
	if err != nil {
		return err
	}

	return nil
}
