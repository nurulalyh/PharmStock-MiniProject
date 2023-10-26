package models

import (
	"time"

	// "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatProducts struct {
	Id        string         `gorm:"primaryKey;type:varchar(10)"`
	Name      string         `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Product   []Products     `gorm:"foreignKey:id_cat_product;references:id"`
}

// type CatProductModelInterface interface {
// 	Insert(newCatProduct CatProduct) *CatProduct
// 	SelectAll() []CatProduct
// 	SelectById(catProductId int) *CatProduct
// 	Update(updatedData CatProduct) *CatProduct
// 	Delete(catProductId int) bool
// }

// type CatProductsModel struct {
// 	db *gorm.DB
// }

// func NewCatProductsModel(db *gorm.DB) CatProductModelInterface {
// 	return &CatProductsModel{
// 		db: db,
// 	}
// }

// func (cpm *CatProductsModel) Insert(newCatProduct CatProduct) *CatProduct {
// 	if err := cpm.db.Create(&newCatProduct).Error; err != nil {
// 		logrus.Error("Model : Insert data error, ", err.Error())
// 		return nil
// 	}

// 	return &newCatProduct
// }

// func (cpm *CatProductsModel) SelectAll() []CatProduct {
// 	var data = []CatProduct{}
// 	if err := cpm.db.Find(&data).Error; err != nil {
// 		logrus.Error("Model : Cannot get all category product, ", err.Error())
// 		return nil
// 	}

// 	return data
// }

// func (cpm *CatProductsModel) SelectById(catProductId int) *CatProduct {
// 	var data = CatProduct{}
// 	if err := cpm.db.Where("id = ?", catProductId).First(&data).Error; err != nil {
// 		logrus.Error("Model : Data with that ID was not found, ", err.Error())
// 		return nil
// 	}

// 	return &data
// }

// func (cpm *CatProductsModel) Update(updatedData CatProduct) *CatProduct {
// 	var qry = cpm.db.Table("cat_products").Where("id = ?", updatedData.Id).Update("name", updatedData.Name)
// 	if err := qry.Error; err != nil {
// 		logrus.Error("Model : update error, ", err.Error())
// 		return nil
// 	}

// 	if dataCount := qry.RowsAffected; dataCount < 1 {
// 		logrus.Error("Model : Update error, ", "no data effected")
// 		return nil
// 	}

// 	var updatedCatProduct = CatProduct{}
// 	if err := cpm.db.Where("id = ?", updatedData.Id).First(&updatedCatProduct).Error; err != nil {
// 		logrus.Error("Model : Error get updated data, ", err.Error())
// 		return nil
// 	}

// 	return &updatedCatProduct
// }

// func (cpm *CatProductsModel) Delete(catProductId int) bool {
// 	var data = CatProduct{}
// 	data.Id = catProductId

// 	if err := cpm.db.Where("id = ?", catProductId).First(&data).Error; err != nil {
// 		logrus.Error("Model: Error finding data to delete, ", err.Error())
// 		return false
// 	}

// 	if err := cpm.db.Delete(&data).Error; err != nil {
// 		logrus.Error("Model : Error delete data, ", err.Error())
// 		return false
// 	}

// 	return true
// }
