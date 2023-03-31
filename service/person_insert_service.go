package service

import (
	"fmt"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/logg"
)

func (service *personService) InsertPerson(newPerson *model.PersonRequest) error {

	newPersonTemp := model.Person{
		PersonID:  newPerson.PersonID,
		LastName:  newPerson.LastName,
		FirstName: newPerson.FirstName,
		Address:   newPerson.Address,
		City:      newPerson.City,
	}

	err := service.PersonRepo.InsertPerson(&newPersonTemp)
	if err != nil {
		logg.PrintloggerJsonMarshalIndentHasHeader("********** INSERT failed ********** | ", "Insert failed ", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm insert failed")
	}

	return nil

}
