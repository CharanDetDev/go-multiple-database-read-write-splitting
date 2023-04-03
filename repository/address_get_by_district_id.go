package repository

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
)

// GetAddressByDistrictID implements AddressRepo
func (repo *addressRepo) GetAddressByDistrictID(districtID int, language string, addressResponse *model.AddressResponseModel) error {

	err := repo.DatabaseConn.Raw(repo.preparationConditions(language)).Scan(&addressResponse).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *addressRepo) preparationConditions(language string) string {

	var temp string
	if language == "th" {
		temp = `
				SELECT 
					dstc.id AS DistrictID,
					dstc.name_th AS DistrictName,
					amph.id AS AmphureID   ,
					amph.name_th AS AmphureName ,
					dstc.zip_code AS ZipCode     ,
					prv.id AS ProvinceID  ,
					prv.name_th AS ProvinceName ,
					gogph.id AS GeographieID ,
					gogph.name AS GeographieName 
				FROM districts AS dstc
				LEFT JOIN amphures AS amph ON amph.id = dstc.amphure_id
				LEFT JOIN provinces prv ON prv.id = amph.province_id
				LEFT JOIN geographies gogph ON gogph.id = prv.geography_id
				WHERE dstc.id = 302202;
			`
	} else {
		temp = `
				SELECT 
					dstc.id AS DistrictID,
					dstc.name_en AS DistrictName,
					amph.id AS AmphureID   ,
					amph.name_en AS AmphureName ,
					dstc.zip_code AS ZipCode     ,
					prv.id AS ProvinceID  ,
					prv.name_en AS ProvinceName ,
					gogph.id AS GeographieID ,
					gogph.name AS GeographieName 
				FROM districts AS dstc
				LEFT JOIN amphures AS amph ON amph.id = dstc.amphure_id
				LEFT JOIN provinces prv ON prv.id = amph.province_id
				LEFT JOIN geographies gogph ON gogph.id = prv.geography_id
				WHERE dstc.id = 302202;
			`
	}

	return temp
}
