package models

import (
	"time"

	// "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReqProducts struct {
	Id          string         `gorm:"primaryKey;type:varchar(10)"`
	IdEmployee  string         `gorm:"type:varchar(10);not null"`
	ProductName string         `gorm:"type:varchar(100);not null"`
	Quantity    int            `gorm:"type:smallint;not null"`
	Note        string         `gorm:"type:text;not null"`
	StatusReq   string         `gorm:"type:ENUM('rejected','processed', 'accepted');not null"`
	CreatedAt   time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// type ReqProductModelInterface interface {
// 	Insert(newReqProduct ReqProduct) *ReqProduct
// 	SelectAll() []ReqProduct
// 	SelectById(reqProductId int) *ReqProduct
// 	Update(updatedData ReqProduct) *ReqProduct
// 	Delete(reqProductId int) bool
// }

// type ReqProductsModel struct {
// 	db *gorm.DB
// }

// func NewReqProductsModel(db *gorm.DB) ReqProductModelInterface {
// 	return &ReqProductsModel{
// 		db: db,
// 	}
// }

// func (rpm *ReqProductsModel) Insert(newReqProduct ReqProduct) *ReqProduct {
// 	if err := rpm.db.Create(&newReqProduct).Error; err != nil {
// 		logrus.Error("Model : Insert data error, ", err.Error())
// 		return nil
// 	}

// 	return &newReqProduct
// }

// func (rpm *ReqProductsModel) SelectAll() []ReqProduct {
// 	var data = []ReqProduct{}
// 	if err := rpm.db.Find(&data).Error; err != nil {
// 		logrus.Error("Model : Cannot get all request product, ", err.Error())
// 		return nil
// 	}

// 	return data
// }

// func (rpm *ReqProductsModel) SelectById(reqProductId int) *ReqProduct {
// 	var data = ReqProduct{}
// 	if err := rpm.db.Where("id = ?", reqProductId).First(&data).Error; err != nil {
// 		logrus.Error("Model : Data with that ID was not found, ", err.Error())
// 		return nil
// 	}

// 	return &data
// }

// func (rpm *ReqProductsModel) Update(updatedData ReqProduct) *ReqProduct {
// 	var data map[string]interface{} = make(map[string]interface{})

// 	if updatedData.ProductName != "" {
// 		data["product_name"] = updatedData.ProductName
// 	}
// 	if updatedData.IdUser != 0 {
// 		data["id_user"] = updatedData.IdUser
// 	}
// 	if updatedData.Quantity != 0 {
// 		data["quantity"] = updatedData.Quantity
// 	}
// 	if updatedData.Note != "" {
// 		data["note"] = updatedData.Note
// 	}

// 	var qry = rpm.db.Table("req_products").Where("id = ?", updatedData.Id).Updates(data)
// 	if err := qry.Error; err != nil {
// 		logrus.Error("Model : update error, ", err.Error())
// 		return nil
// 	}

// 	if dataCount := qry.RowsAffected; dataCount < 1 {
// 		logrus.Error ("Model : Update error, ", "no data effected")
// 		return nil
// 	}

// 	var updatedReqProduct = ReqProduct{}
// 	if err := rpm.db.Where("id = ?", updatedData.Id).First(&updatedReqProduct).Error; err != nil {
// 		logrus.Error("Model : Error get updated data, ", err.Error())
// 		return nil
// 	}

// 	return &updatedReqProduct
// }

// func (rpm *ReqProductsModel) Delete(reqProductId int) bool {
// 	var data = ReqProduct{}
// 	data.Id = reqProductId

// 	if err := rpm.db.Where("id = ?", reqProductId).First(&data).Error; err != nil {
// 		logrus.Error("Model: Error finding data to delete, ", err.Error())
// 		return false
// 	}

// 	if err := rpm.db.Delete(&data).Error; err != nil {
// 		logrus.Error("Model : Error delete data, ", err.Error())
// 		return false
// 	}

// 	return true
// }
