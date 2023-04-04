package domain

import "github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"

type AddressRepo interface {
	GetAddressByDistrictID(districtID string, language string, addressResponse *model.AddressResponseModel) error
	GetAddressByAmpureID(ampureID int, language string, addressResponse *model.AddressResponseModel) error
	GetAddressByProvinceID(provinceID int, language string, addressResponse *model.AddressResponseModel) error
	GetAddressByGography(geographieID int, addressResponse *model.AddressResponseModel) error
}
