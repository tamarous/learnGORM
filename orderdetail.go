package main

type OrderDetail struct {
	OrderNumber     int64   `gorm:"type:int;column:orderNumber;primaryKey"`
	ProductCode     string  `gorm:"type:varchar(15);column:productCode;index:productCode;primaryKey"`
	QuantityOrdered int64   `gorm:"type:int;column:quantityOrdered"`
	PriceEach       float64 `gorm:"type:decimal;column:priceEach"`
	OrderLineNumber int8    `gorm:"type:smallint;column:orderLineNumber"`
}
