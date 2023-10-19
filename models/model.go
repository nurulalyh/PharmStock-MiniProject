package models

import (
	"fmt"
	"pharm-stock/configs"
	"pharm-stock/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitModel(config configs.Config) *gorm.DB {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		helper.ErrorResponse("Model Error: cannot connect to database", err.Error())
		return nil

	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Distributor{})
	db.AutoMigrate(&CatProduct{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&ReqProduct{})
	db.AutoMigrate(&DetailProduct{})
}