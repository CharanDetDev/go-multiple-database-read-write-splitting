package repository

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/domain"
	"gorm.io/gorm"
)

type personRepo struct {
	DatabaseConn *gorm.DB
}

func NewPersonRepo(dbConn *gorm.DB) domain.PersonRepo {
	return &personRepo{
		DatabaseConn: dbConn,
	}
}
