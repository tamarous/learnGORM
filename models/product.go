package model

type Product struct {
	ProductCode        string  `gorm:"type:varchar(15);column:productCode;primaryKey"`
	ProductName        string  `gorm:"type:varchar(70);column:productName"`
	ProductLine        string  `gorm:"type:varchar(50);column:productLine;index:productLine"`
	ProductScale       string  `gorm:"type:varchar(10);column:productScale"`
	ProductDescription string  `gorm:"type:text;column:productDescription"`
	QuantityInStock    int64   `gorm:"type:smallint;column:quantityInStock"`
	BuyPrice           float64 `gorm:"type:decimal;column:buyPrice"`
	MSRP               float64 `gorm:"type:decimal;column:MSRP"`
}
