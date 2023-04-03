package domain

import "github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"

type PersonRepo interface {
	GetPersonWithPersonID(personId int, person *model.PersonModel) error
	InsertPerson(newPerson *model.PersonModel) error
	UpdatePerson(newPerson *model.PersonModel) error
	DeletePerson(personID int) error
}
