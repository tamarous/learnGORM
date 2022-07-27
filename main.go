package main

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	action "main/actions"
)

func main() {

	viper.SetConfigName("db")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("err is %v", err)
	}

	username := viper.GetString("MySQL.Username")
	password := viper.GetString("MySQL.Password")
	port := viper.GetInt("MySQL.Port")

	fmt.Printf("username: %s, password: %s, port: %d\n", username, password, port)

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/classicmodels?charset=utf8mb4&parseTime=True&loc=Local", username, password, port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	action.Query(db)
}
