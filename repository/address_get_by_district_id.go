package repository

import (
	"strconv"

	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/model"
)

// GetAddressByDistrictID implements AddressRepo
func (repo *addressRepo) GetAddressByDistrictID(districtID string, language string, addressResponse *model.AddressResponseModel) error {

	var district model.DistrictModel
	if language == "th" {
		// WHERE แบบที่ 1 => .First(&district, districtID)
		if result := repo.DatabaseConn.Select("id, name_th, zip_code, amphure_id").First(&district, districtID); result.Error != nil {
			return result.Error
		}
		addressResponse.DistrictName = district.NameTh
	} else {
		// WHERE แบบที่ 2 => .Where("id=?", districtID)
		if result := repo.DatabaseConn.Select("id, name_en, zip_code, amphure_id").Where("id=?", districtID).First(&district); result.Error != nil {
			return result.Error
		}
		addressResponse.DistrictName = district.NameEn
	}

	addressResponse.DistrictID, _ = strconv.Atoi(district.ID)
	addressResponse.AmphureID = district.AmphureID
	addressResponse.ZipCode = district.ZipCode

	return nil
}

// func (repo *addressRepo) GetAddressByDistrictID(districtID string, language string, addressResponse *model.AddressResponseModel) error {

// err := repo.DatabaseConn.Raw(repo.preparationRowStatement(districtID, language)).Scan(&addressResponse).Error
// if err != nil {
// 	return err
// }

// result := repo.DatabaseConn.
// 	Table(fmt.Sprintf("%s AS dstc", model.DistrictModel{}.TableName())).
// 	Select(repo.preparationSelectStatement(language)).
// 	Joins(fmt.Sprintf("LEFT JOIN %s AS amph ON amph.id = dstc.amphure_id", model.AmphureModel{}.TableName())).
// 	Joins(fmt.Sprintf("LEFT JOIN %s AS prv ON prv.id = amph.province_id", model.ProvinceModel{}.TableName())).
// 	Joins(fmt.Sprintf("LEFT JOIN %s AS gogph ON gogph.id = prv.geography_id", model.GeographieModel{}.TableName())).
// 	Where("dstc.id = ?", districtID).
// 	Scan(&addressResponse)
// if result.Error != nil {
// 	return result.Error
// }

// 	return nil
// }

// func (repo *addressRepo) preparationRowStatement(districtID string, language string) string {

// 	var temp string
// 	if language == "th" {
// 		temp = fmt.Sprintf(`
// 				SELECT
// 					dstc.id AS DistrictID,
// 					dstc.name_th AS DistrictName,
// 					amph.id AS AmphureID   ,
// 					amph.name_th AS AmphureName ,
// 					dstc.zip_code AS ZipCode     ,
// 					prv.id AS ProvinceID  ,
// 					prv.name_th AS ProvinceName ,
// 					gogph.id AS GeographieID ,
// 					gogph.name AS GeographieName
// 				FROM districts AS dstc
// 				LEFT JOIN amphures AS amph ON amph.id = dstc.amphure_id
// 				LEFT JOIN provinces prv ON prv.id = amph.province_id
// 				LEFT JOIN geographies gogph ON gogph.id = prv.geography_id
// 				WHERE dstc.id = %v;
// 			`, districtID)
// 	} else {
// 		temp = fmt.Sprintf(`
// 				SELECT
// 					dstc.id AS DistrictID,
// 					dstc.name_en AS DistrictName,
// 					amph.id AS AmphureID   ,
// 					amph.name_en AS AmphureName ,
// 					dstc.zip_code AS ZipCode     ,
// 					prv.id AS ProvinceID  ,
// 					prv.name_en AS ProvinceName ,
// 					gogph.id AS GeographieID ,
// 					gogph.name AS GeographieName
// 				FROM districts AS dstc
// 				LEFT JOIN amphures AS amph ON amph.id = dstc.amphure_id
// 				LEFT JOIN provinces prv ON prv.id = amph.province_id
// 				LEFT JOIN geographies gogph ON gogph.id = prv.geography_id
// 				WHERE dstc.id = %v;
// 			`, districtID)
// 	}

// 	return temp
// }

// func (repo *addressRepo) preparationSelectStatement(language string) string {

// 	var temp string
// 	if language == "th" {
// 		temp = `
// 					dstc.id AS DistrictID,
// 					dstc.name_th AS DistrictName,
// 					amph.id AS AmphureID   ,
// 					amph.name_th AS AmphureName ,
// 					dstc.zip_code AS ZipCode     ,
// 					prv.id AS ProvinceID  ,
// 					prv.name_th AS ProvinceName ,
// 					gogph.id AS GeographieID ,
// 					gogph.name AS GeographieName
// 				`
// 	} else {
// 		temp = `
// 					dstc.id AS DistrictID,
// 					dstc.name_en AS DistrictName,
// 					amph.id AS AmphureID   ,
// 					amph.name_en AS AmphureName ,
// 					dstc.zip_code AS ZipCode     ,
// 					prv.id AS ProvinceID  ,
// 					prv.name_en AS ProvinceName ,
// 					gogph.id AS GeographieID ,
// 					gogph.name AS GeographieName
// 				`
// 	}

// 	return temp
// }
