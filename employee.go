package main

type Employee struct {
	Number     int64  `gorm:"column:employeeNumber;type:int;primaryKey"`
	LastName   string `gorm:"column:lastName;type:varchar(50)"`
	FirstName  string `gorm:"column:firstName;type:varchar(50)"`
	Extension  string `gorm:"column:extension;type:varchar(10)"`
	Email      string `gorm:"column:email;type:varchar(100)"`
	OfficeCode string `gorm:"column:officeCode;type:varchar(10);index:officeCode"`
	ReportsTo  int64  `gorm:"column:reportsTo;type:int;index:reportsTo"`
	JobTitle   string `gorm:"column:jobTitle;type:varchar(50)"`
}
