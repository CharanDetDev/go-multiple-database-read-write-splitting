package service

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/domain"
)

type personService struct {
	PersonRepo     domain.PersonRepo
	AddaddressRepo domain.AddressRepo
}

func NewPersonService(personRepo domain.PersonRepo, addaddressRepo domain.AddressRepo) domain.PersonService {
	return &personService{
		PersonRepo:     personRepo,
		AddaddressRepo: addaddressRepo,
	}
}
