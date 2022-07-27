package model

type Office struct {
	OfficeCode   string `gorm:"type:varchar(10);column:officeCode;primaryKey"`
	City         string `gorm:"type:varchar(50);column:city"`
	Phone        string `gorm:"type:varchar(50):column:phone"`
	AddressLine1 string `gorm:"type:varchar(50);column:addressLine1"`
	AddressLine2 string `gorm:"type:varchar(50);column:addressLine2"`
	State        string `gorm:"type:varchar(50);column:state"`
	Country      string `gorm:"type:varchar(50);column:country"`
	PostalCode   string `gorm:"type:varchar(15);column:postalCode"`
	Territory    string `gorm:"type:varchar(10);column:territory"`
}
