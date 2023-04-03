package domain

import "github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"

type AddressRepo interface {
	GetAddressByDistrictID(districtID int, language string, addressResponse *model.AddressResponseModel) error
}
