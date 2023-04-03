package repository

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/domain"
	"gorm.io/gorm"
)

type addressRepo struct {
	DatabaseConn *gorm.DB
}

func NewAddressRepo(db *gorm.DB) domain.AddressRepo {
	return &addressRepo{
		DatabaseConn: db,
	}
}
