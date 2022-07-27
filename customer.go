package main

type Customer struct {
	Number                 int64   `gorm:"column:customerNumber;primaryKey;type:int"`
	Name                   string  `gorm:"column:customerName;type:varchar(50)"`
	LastName               string  `gorm:"column:contactLastName;type:varchar(50)"`
	FirstName              string  `gorm:"column:contactFirstName;type:varchar(50)"`
	Phone                  string  `gorm:"column:phone;type:varchar(50)"`
	AddressLine1           string  `gorm:"column:addressLine1;type:varchar(50)"`
	AddressLine2           string  `gorm:"column:addressLine2;type:varchar(50)"`
	City                   string  `gorm:"column:city;type:varchar(50)"`
	State                  string  `gorm:"column:state;type:varchar(50)"`
	PostalCode             string  `gorm:"column:postalCode;type:varchar(15)"`
	Country                string  `gorm:"column:country;type:varchar(50)"`
	SalesRepEmployeeNumber int64   `gorm:"column:salesRepEmployeeNumber;type:int"`
	CreditLimit            float64 `gorm:"column:creditLimit"`
}
