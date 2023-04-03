package domain

import "github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"

type PersonService interface {
	GetPersonWithPersonID(personId int, person *model.PersonModel) error
	InsertPerson(newPerson *model.PersonRequest) error
	UpdatePerson(newPerson *model.PersonRequest) error
	DeletePerson(personID int) error
}
