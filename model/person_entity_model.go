package model

type PersonModel struct {
	PersonID       int    `gorm:"column:PersonID;primary_key;AUTO_INCREMENT"` // Primary Key
	LastName       string `gorm:"column:LastName"`
	FirstName      string `gorm:"column:FirstName"`
	DistrictID     int    `gorm:"column:DistrictID"`
	DistrictName   string `gorm:"column:DistrictName;NOT NULL"`
	AmphureID      int    `gorm:"column:AmphureID"`
	AmphureName    string `gorm:"column:AmphureName"`
	ZipCode        int    `gorm:"column:ZipCode"`
	ProvinceID     int    `gorm:"column:ProvinceID"`
	ProvinceName   string `gorm:"column:ProvinceName"`
	GeographieID   int    `gorm:"column:GeographieID"`
	GeographieName string `gorm:"column:GeographieName"`
}

func (PersonModel) TableName() string {
	return "Persons"
}
