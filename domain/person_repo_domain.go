package domain

import "github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"

type PersonRepo interface {
	GetPersonWithPersonID(personId int, person *model.Person) error
	InsertPerson(newPerson *model.Person) error
	UpdatePerson(newPerson *model.Person) error
	DeletePerson(personID int) error
}
