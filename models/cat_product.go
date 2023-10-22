package models

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatProduct struct {
	Id             int             `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Name           string          `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	CreatedAt      time.Time       `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt      time.Time       `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt      gorm.DeletedAt  `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DetailProducts []DetailProduct `gorm:"foreignKey:IdDistributor;references:Id"`
}

type CatProductModelInterface interface {
	Insert(newCatProduct CatProduct) *CatProduct
	SelectAll() []CatProduct
	SelectById(carProductId int) *CatProduct
	Update(updatedData CatProduct) *CatProduct
	Delete(userId int) bool
}

type CatProductsModel struct {
	db *gorm.DB
}

func NewCatProductsModel(db *gorm.DB) CatProductModelInterface {
	return &CatProductsModel{
		db: db,
	}
}

func (cpm *CatProductsModel) Insert(newCatProduct CatProduct) *CatProduct {
	if err := cpm.db.Create(&newCatProduct).Error; err != nil {
		logrus.Error("Model : Insert data error, ", err.Error())
		return nil
	}

	return &newCatProduct
}

func (cpm *CatProductsModel) SelectAll() []CatProduct {
	var data = []CatProduct{}
	if err := cpm.db.Find(&data).Error; err != nil {
		logrus.Error("Model : Cannot get all category product, ", err.Error())
		return nil
	}

	return data
}

func (cpm *CatProductsModel) SelectById(catProductId int) *CatProduct {
	var data = CatProduct{}
	if err := cpm.db.Where("id = ?", catProductId).First(&data).Error; err != nil {
		logrus.Error("Model : Data with that ID was not found, ", err.Error())
		return nil
	}

	return &data
}

func (cpm *CatProductsModel) Update(updatedData CatProduct) *CatProduct {
	var qry = cpm.db.Table("users").Where("id = ?", updatedData.Id).Update("name", updatedData.Name)
	if err := qry.Error; err != nil {
		logrus.Error("Model : update error, ", err.Error())
		return nil
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update error, ", "no data effected")
		return nil
	}

	var updatedCatProduct = CatProduct{}
	if err := cpm.db.Where("id = ?", updatedData.Id).First(&updatedCatProduct).Error; err != nil {
		logrus.Error("Model : Error get updated data, ", err.Error())
		return nil
	}

	return &updatedCatProduct
}

func (cpm *CatProductsModel) Delete(userId int) bool {
	var data = User{}
	data.Id = userId

	if err := cpm.db.Where("id = ?", userId).First(&data).Error; err != nil {
		logrus.Error("Model: Error finding data to delete, ", err.Error())
		return false
	}

	if err := cpm.db.Delete(&data).Error; err != nil {
		logrus.Error("Model : Error delete data, ", err.Error())
		return false
	}

	return true
}
