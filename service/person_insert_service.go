package service

import (
	"fmt"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/logg"
)

func (service *personService) InsertPerson(newPerson *model.PersonRequest) error {

	var addressResponse model.AddressResponseModel
	districtID := fmt.Sprintf("%v", newPerson.DistrictID)
	err := service.AddaddressRepo.GetAddressByDistrictID(districtID, newPerson.Language, &addressResponse)
	if err != nil {
		logg.PrintloggerJsonMarshalIndentHasHeader("********** GET Address failed ********** | ", "Insert failed, Method InsertPerson() at GetAddressByDistrictID() | ", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm get address failed")
	}

	err = service.AddaddressRepo.GetAddressByAmpureID(addressResponse.AmphureID, newPerson.Language, &addressResponse)
	if err != nil {
		logg.PrintloggerJsonMarshalIndentHasHeader("********** GET Address failed ********** | ", "Insert failed, Method InsertPerson() at GetAddressByAmpureID() | ", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm get address failed")
	}

	err = service.AddaddressRepo.GetAddressByProvinceID(addressResponse.ProvinceID, newPerson.Language, &addressResponse)
	if err != nil {
		logg.PrintloggerJsonMarshalIndentHasHeader("********** GET Address failed ********** | ", "Insert failed, Method InsertPerson() at GetAddressByProvinceID() | ", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm get address failed")
	}

	err = service.AddaddressRepo.GetAddressByGography(addressResponse.GeographieID, &addressResponse)
	if err != nil {
		logg.PrintloggerJsonMarshalIndentHasHeader("********** GET Address failed ********** | ", "Insert failed, Method InsertPerson() at GetAddressByGography() | ", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm get address failed")
	}

	newPersonTemp := model.PersonModel{
		PersonID:       newPerson.PersonID,
		LastName:       newPerson.LastName,
		FirstName:      newPerson.FirstName,
		DistrictID:     addressResponse.DistrictID,
		DistrictName:   addressResponse.DistrictName,
		AmphureID:      addressResponse.AmphureID,
		AmphureName:    addressResponse.AmphureName,
		ZipCode:        addressResponse.ZipCode,
		ProvinceID:     addressResponse.ProvinceID,
		ProvinceName:   addressResponse.ProvinceName,
		GeographieID:   addressResponse.GeographieID,
		GeographieName: addressResponse.GeographieName,
	}

	err = service.PersonRepo.InsertPerson(&newPersonTemp)
	if err != nil {
		logg.PrintloggerJsonMarshalIndentHasHeader("********** INSERT failed ********** | ", "Insert failed ", fmt.Sprintf("%v | %v", err.Error(), logg.GetCallerPathNameFileNameLineNumber()))
		return fmt.Errorf("gorm insert failed")
	}

	return nil
}
