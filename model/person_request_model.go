package model

type PersonRequest struct {
	PersonID       int    `json:"person_id"`
	LastName       string `json:"last_name" validate:"required"`
	FirstName      string `json:"first_name" validate:"required"`
	DistrictID     int    `json:"district_id" validate:"required"`
	DistrictName   string `json:"district_name"`
	AmphureName    string `json:"amphure_name"`
	ZipCode        int    `json:"zip_code"`
	ProvinceName   string `json:"province_name"`
	GeographieName string `json:"geographie_name"`
	Language       string `json:"language"`
}
