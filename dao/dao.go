package dao

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLSetting struct {
	Username string
	Password string
	Port     int
}

func Initialize() (*gorm.DB, error) {
	viper.SetConfigName("db")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("err is %v", err)
		panic(err)
	}

	var mysqlSetting MySQLSetting
	if err = viper.UnmarshalKey("MySQL", &mysqlSetting); err != nil {
		panic(err)
	}

	username := mysqlSetting.Username
	password := mysqlSetting.Password
	port := mysqlSetting.Port
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/classicmodels?charset=utf8mb4&parseTime=True&loc=Local", username, password, port)

	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return engine, nil
}
