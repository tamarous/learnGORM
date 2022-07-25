package main

import (
	"time"
)

type Order struct {
	OrderNumber    int64     `gorm:"type:int;column:orderNumber;primaryKey;index:orderNumber"`
	OrderDate      time.Time `gorm:"type:date;column:orderDate"`
	RequiredDate   time.Time `gorm:"type:date;column:requiredDate"`
	ShippedDate    time.Time `gorm:"type:date;column:shippedDate"`
	Status         string    `gorm:"type:varchar(15);column:status"`
	Comments       string    `gorm:"type:text;column:comments"`
	CustomerNumber int64     `gorm:"type:int;column:customerNumber"`
}
