package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/classicmodels?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var customer Customer
	db.Take(&customer)
	fmt.Printf("customer: %v", customer)

	fmt.Println("")

	var employee Employee
	db.Take(&employee)
	fmt.Printf("employee: %v", employee)

	fmt.Println("")

	var office Office
	db.Take(&office)
	fmt.Printf("office: %v", office)

	fmt.Println("")

	var orderDetail OrderDetail
	db.Table("orderdetails").Take(&orderDetail)
	fmt.Printf("orderDetail: %v", orderDetail)

	fmt.Println("")

	var order Order
	db.Take(&order)
	fmt.Printf("order: %v", order)

	fmt.Println("")

	var payment Payment
	db.Take(&payment)
	fmt.Printf("payment: %v", payment)

	fmt.Println("")

	var productLine ProductLine
	db.Take(&productLine)
	fmt.Printf("productLine: %v", productLine)

	fmt.Println("")

	var product Product
	db.Take(&product)
	fmt.Printf("product: %v", product)

}
