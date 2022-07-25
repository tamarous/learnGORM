package main

type ProductLine struct {
	ProductLine     string `gorm:"type:varchar(50);primaryKey;column:productLine"`
	TextDescription string `gorm:"type:varchar(4000);column:textDescription"`
	HtmlDescription string `gorm:"type:mediumtext;column:htmlDescription"`
	Image           []byte `gorm:"type:mediumblob;column:image;column:image"`
}

func (p ProductLine) TableName() string {
	return "productlines"
}
