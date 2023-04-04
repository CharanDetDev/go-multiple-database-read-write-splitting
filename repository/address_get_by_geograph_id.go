package repository

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
)

// GetAddressByGography implements domain.AddressRepo
func (repo *addressRepo) GetAddressByGography(geographieID int, addressResponse *model.AddressResponseModel) error {

	var geographie model.GeographieModel
	if result := repo.DatabaseConn.Select("id, name").First(&geographie, geographieID); result.Error != nil {
		return result.Error
	}

	addressResponse.GeographieID = geographie.ID
	addressResponse.GeographieName = geographie.Name

	return nil
}
