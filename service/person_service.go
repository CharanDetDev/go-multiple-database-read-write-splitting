package service

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/domain"
)

type personService struct {
	PersonRepo domain.PersonRepo
}

func NewPersonService(personRepo domain.PersonRepo) domain.PersonService {
	return &personService{
		PersonRepo: personRepo,
	}
}
