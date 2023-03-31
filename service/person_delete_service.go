package service

import (
	"fmt"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/logg"
)

func (service *personService) DeletePerson(personID int) error {

	err := service.PersonRepo.DeletePerson(personID)
	if err != nil {
		logg.PrintloggerJsonMarshalIndentHasHeader("***** DELETE failed ***** | ", "Delete failed ", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm delete failed")
	}

	return nil

}
