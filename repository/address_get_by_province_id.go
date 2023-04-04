package repository

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
)

// GetAddressByProvinceID implements domain.AddressRepo
func (repo *addressRepo) GetAddressByProvinceID(provinceID int, language string, addressResponse *model.AddressResponseModel) error {

	var province model.ProvinceModel
	if language == "th" {
		// WHERE แบบที่ 1 => .First(&district, provinceID)
		if result := repo.DatabaseConn.Select("id, name_th, geography_id").First(&province, provinceID); result.Error != nil {
			return result.Error
		}
		addressResponse.ProvinceName = province.NameTh
	} else {
		// WHERE แบบที่ 2 => .Where("id=?", provinceID)
		if result := repo.DatabaseConn.Select("id, name_en, geography_id").Where("id=?", provinceID).First(&province); result.Error != nil {
			return result.Error
		}
		addressResponse.ProvinceName = province.NameEn
	}

	addressResponse.ProvinceID = province.ID
	addressResponse.GeographieID = province.GeographyID

	return nil
}
