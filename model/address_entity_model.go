package model

type DistrictModel struct {
	ID        string `gorm:"column:id;NOT NULL"`
	ZipCode   int    `gorm:"column:zip_code;NOT NULL"`
	NameTh    string `gorm:"column:name_th;NOT NULL"`
	NameEn    string `gorm:"column:name_en;NOT NULL"`
	AmphureID int    `gorm:"column:amphure_id;default:0;NOT NULL"`
}

type AmphureModel struct {
	ID         int    `gorm:"column:id;NOT NULL"`
	Code       string `gorm:"column:code;NOT NULL"`
	NameTh     string `gorm:"column:name_th;NOT NULL"`
	NameEn     string `gorm:"column:name_en;NOT NULL"`
	ProvinceID int    `gorm:"column:province_id;default:0;NOT NULL"`
}

type ProvinceModel struct {
	ID          int    `gorm:"column:id;NOT NULL"`
	Code        string `gorm:"column:code;NOT NULL"`
	NameTh      string `gorm:"column:name_th;NOT NULL"`
	NameEn      string `gorm:"column:name_en;NOT NULL"`
	GeographyID int    `gorm:"column:geography_id;default:0;NOT NULL"`
}

type GeographieModel struct {
	ID   int    `gorm:"column:id;NOT NULL"`
	Name string `gorm:"column:name;NOT NULL"`
}

func (DistrictModel) TableName() string {
	return "districts"
}

func (AmphureModel) TableName() string {
	return "amphures"
}

func (ProvinceModel) TableName() string {
	return "provinces"
}

func (GeographieModel) TableName() string {
	return "geographies"
}
