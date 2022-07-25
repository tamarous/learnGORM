package main

import (
	"time"
)

type Payment struct {
	CustomerNumber int64     `gorm:"type:int;column:customerNumber;primaryKey"`
	CheckNumber    string    `gorm:"type:varchar(50);column:checkNumber;primaryKey"`
	PaymentDate    time.Time `gorm:"type:date;column:paymentDate"`
	Amount         float64   `gorm:"type:decimal;column:amount"`
}
