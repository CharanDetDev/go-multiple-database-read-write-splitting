package service

import (
	"fmt"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/logg"
	"gorm.io/gorm"
)

func (service *personService) GetPersonWithPersonID(personId int, person *model.Person) error {

	err := service.PersonRepo.GetPersonWithPersonID(personId, person)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logg.PrintloggerJsonMarshalIndentHasHeader("********** GET failed ********** | ", "Method GetPersonWithPersonID() ", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}
