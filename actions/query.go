package action

import (
	"fmt"
	model "main/models"

	"gorm.io/gorm"
)

func Query(db *gorm.DB) {
	fmt.Println("------------------------------------------")
	fmt.Println("Query Action")

	var customer model.Customer
	db.Take(&customer)
	fmt.Printf("customer: %v", customer)

	fmt.Println("")

	var customer_103 model.Customer
	db.Take(&customer_103, 103)
	fmt.Printf("customer_103: %v", customer_103)
	fmt.Println("")

	customersMatchId := make([]model.Customer, 10)
	db.Find(&customersMatchId, []int{103, 112, 114, 119, 121, 124, 125, 128, 129, 131})

	for i, c := range customersMatchId {
		fmt.Printf("customersMatchId[%d]: %v\n", i, c)
	}

	var parisCustomers []model.Customer
	db.Where("city = ?", "Paris").Find(&parisCustomers)
	fmt.Println("customers in Paris:")

	for i, c := range parisCustomers {
		fmt.Printf("parisCustomers[%d]:%v\n", i, c)
	}

	var partialCustomes []model.Customer
	db.Select("customerName", "contactLastName", "contactFirstName").Where("city = ?", "Paris").Find(&partialCustomes)
	for i, c := range partialCustomes {
		fmt.Printf("partialCustomes[%d]:%v\n", i, c)
	}

	var sortedCustomes []model.Customer
	db.Where("city = ?", "Paris").Order("customerName asc, customerNumber asc").Find(&sortedCustomes)
	for i, c := range sortedCustomes {
		fmt.Printf("sortedCustomes[%d]:%v\n", i, c)
	}

	limit3Customers := make([]model.Customer, 3)
	db.Limit(3).Find(&limit3Customers)
	for i, c := range limit3Customers {
		fmt.Printf("limit3Customers[%d]:%v\n", i, c)
	}

}
