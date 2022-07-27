package main

import (
	"fmt"

	"gorm.io/gorm"
)

func Query(db *gorm.DB) {
	var customer Customer
	db.Take(&customer)
	fmt.Printf("customer: %v", customer)

	fmt.Println("")

	var customer_103 Customer
	db.Take(&customer_103, 103)
	fmt.Printf("customer_103: %v", customer_103)

	fmt.Println("")

	generalCustomer := map[string]interface{}{}
	db.Model(&Customer{}).First(&generalCustomer)
	fmt.Printf("generalCustomer: %v", generalCustomer)

	var employee Employee
	db.Last(&employee)
	fmt.Printf("employee: %v", employee)

	fmt.Println("")

	var office Office
	db.First(&office)
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
