package model

type AddressResponseModel struct {
	DistrictID     int    `gorm:"column:DistrictID;NOT NULL"`
	DistrictName   string `gorm:"column:DistrictName;NOT NULL"`
	AmphureID      int    `gorm:"column:AmphureID;NOT NULL"`
	AmphureName    string `gorm:"column:AmphureName;NOT NULL"`
	ZipCode        int    `gorm:"column:ZipCode;NOT NULL"`
	ProvinceID     int    `gorm:"column:ProvinceID;NOT NULL"`
	ProvinceName   string `gorm:"column:ProvinceName;NOT NULL"`
	GeographieID   int    `gorm:"column:GeographieID;NOT NULL"`
	GeographieName string `gorm:"column:GeographieName;NOT NULL"`
}
