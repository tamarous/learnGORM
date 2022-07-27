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

	customers := make([]model.Customer, 10)
	db.Find(&customers, []int{103, 112, 114, 119, 121, 124, 125, 128, 129, 131})

	fmt.Println("customers match ids:")
	for i, c := range customers {
		fmt.Printf("customers[%d]: %v\n", i, c)
	}

	var parisCustomers []model.Customer
	db.Where("city = ?", "Paris").Find(&parisCustomers)
	fmt.Println("customers in Paris:")

	for i, c := range parisCustomers {
		fmt.Printf("parisCustomers[%d]:%v\n", i, c)
	}

	var employee model.Employee
	db.Last(&employee)
	fmt.Printf("employee: %v", employee)

	fmt.Println("")

	var office model.Office
	db.First(&office)
	fmt.Printf("office: %v", office)

	fmt.Println("")

	var orderDetail model.OrderDetail
	db.Table("orderdetails").Take(&orderDetail)
	fmt.Printf("orderDetail: %v", orderDetail)

	fmt.Println("")

	var order model.Order
	db.Take(&order)
	fmt.Printf("order: %v", order)

	fmt.Println("")

	var payment model.Payment
	db.Take(&payment)
	fmt.Printf("payment: %v", payment)

	fmt.Println("")

	var productLine model.ProductLine
	db.Take(&productLine)
	fmt.Printf("productLine: %v", productLine)

	fmt.Println("")

	var product model.Product
	db.Take(&product)
	fmt.Printf("product: %v", product)
}
