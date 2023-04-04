package repository

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
)

// GetAddressByAmpureID implements domain.AddressRepo
func (repo *addressRepo) GetAddressByAmpureID(ampureID int, language string, addressResponse *model.AddressResponseModel) error {

	var ampure model.AmphureModel
	if language == "th" {
		// WHERE แบบที่ 1 => .First(&district, ampureID)
		if result := repo.DatabaseConn.Select("id, name_th, province_id").First(&ampure, ampureID); result.Error != nil {
			return result.Error
		}
		addressResponse.AmphureName = ampure.NameTh
	} else {
		// WHERE แบบที่ 2 => .Where("id=?", ampureID)
		if result := repo.DatabaseConn.Select("id, name_en, province_id").Where("id=?", ampureID).First(&ampure); result.Error != nil {
			return result.Error
		}
		addressResponse.AmphureName = ampure.NameEn
	}

	addressResponse.AmphureID = ampure.ID
	addressResponse.ProvinceID = ampure.ProvinceID

	return nil
}
