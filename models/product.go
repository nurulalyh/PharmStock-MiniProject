package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Struct Product
type Products struct {
	Id                string               `gorm:"primaryKey;type:varchar(10)"`
	Name              string               `gorm:"type:varchar(100);not null"`
	Photo             string               `gorm:"type:text;not null"`
	IdCatProduct      string               `gorm:"type:varchar(10);not null"`
	MfDate            time.Time            `gorm:"type:date;not null"`
	ExpDate           time.Time            `gorm:"type:date;not null"`
	BatchNumber       int                  `gorm:"type:smallint;not null"`
	UnitPrice         int                  `gorm:"type:smallint;not null"`
	Stock             int                  `gorm:"type:smallint;not null"`
	Description       string               `gorm:"type:text;not null"`
	IdDistributor     string               `gorm:"type:varchar(10);not null"`
	CreatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt         gorm.DeletedAt       `gorm:"index"`
	DetailTransaction []DetailTransactions `gorm:"foreignKey:id_product;references:Id"`
}

// Interface beetween models and controller
type ProductsModelInterface interface {
	Insert(newProduct Products) (*Products, error)
	SelectAll(limit, offset int) ([]Products, error)
	Update(updatedData Products) (*Products, error)
	Delete(productId string) (bool, error)
	SearchProduct(keyword string, limit int, offset int) ([]Products, error)
}

// connect into db
type ProductsModel struct {
	db *gorm.DB
}

// New Instance from ProductsModel
func NewProductsModel(db *gorm.DB) ProductsModelInterface {
	return &ProductsModel{
		db: db,
	}
}

// Insert Product
func (pm *ProductsModel) Insert(newProduct Products) (*Products, error) {
	var latestProduct Products
	if errSort := pm.db.Unscoped().Order("id DESC").First(&latestProduct).Error; errSort != nil {
		latestProduct.Id = "PDT-0000"
	}

	newID := generateProductId(latestProduct.Id)
	if newID == "" {
		return nil, errors.New("Failed generate Id")
	}

	newProduct.Id = newID
	
	validate := validateProduct(newProduct, pm.db)
	if !validate {
		return nil, errors.New("Data not valid")
	}

	if err := pm.db.Create(&newProduct).Error; err != nil {
		return nil, errors.New("Error insert product, " + err.Error())
	}

	return &newProduct, nil
}

// Select All Product
func (pm *ProductsModel) SelectAll(limit, offset int) ([]Products, error) {
	var data []Products
	if err := pm.db.Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get all product, " + err.Error())
	}

	return data, nil
}

// Update Product
func (pm *ProductsModel) Update(updatedData Products) (*Products, error) {
	var data map[string]interface{} = make(map[string]interface{})

	if updatedData.Name != "" {
		data["name"] = updatedData.Name
	}
	if updatedData.Photo != "" {
		data["photo"] = updatedData.Photo
	}
	if updatedData.IdCatProduct != "" {
		data["id_cat_product"] = updatedData.IdCatProduct
	}
	if !updatedData.MfDate.IsZero() {
		data["mf_date"] = updatedData.MfDate
	}
	if !updatedData.ExpDate.IsZero() {
		data["exp_date"] = updatedData.ExpDate
	}
	if updatedData.BatchNumber != 0 {
		data["batch_number"] = updatedData.BatchNumber
	}
	if updatedData.UnitPrice != 0 {
		data["unit_price"] = updatedData.UnitPrice
	}
	if updatedData.Stock != 0 {
		data["stock"] = updatedData.Stock
	}
	if updatedData.Description != "" {
		data["description"] = updatedData.Description
	}
	if updatedData.IdDistributor != "" {
		data["id_distributor"] = updatedData.IdDistributor
	}

	var qry = pm.db.Table("products").Where("id = ?", updatedData.Id).Updates(data)
	if err := qry.Error; err != nil {
		return nil, errors.New("Error update data" + err.Error())
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return nil, errors.New("Error update data, no data affected")
	}

	var updatedProduct = Products{}
	if err := pm.db.Where("id = ?", updatedData.Id).First(&updatedProduct).Error; err != nil {
		return nil, errors.New("Cannot get updated data" + err.Error())
	}

	return &updatedProduct, nil
}

// Delete Product
func (pm *ProductsModel) Delete(productId string) (bool, error) {
	var data = Products{}
	data.Id = productId

	if err := pm.db.Where("id = ?", productId).First(&data).Error; err != nil {
		return false, errors.New("Error finding data to delete, " + err.Error())
	}

	if err := pm.db.Delete(&data).Error; err != nil {
		return false, errors.New("Error delete data, " + err.Error())
	}

	return true, nil
}

// Searching
func (pm *ProductsModel) SearchProduct(keyword string, limit int, offset int) ([]Products, error) {
	var Product []Products
	query := pm.db.Where("id LIKE ? OR name LIKE ? OR photo LIKE ? OR id_cat_product LIKE ? OR mf_Date LIKE ? OR exp_date LIKE ? OR batch_number LIKE ? OR unit_price LIKE ? OR stock LIKE ? OR description LIKE ? OR id_distributor LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Limit(limit).Offset(offset)

	if err := query.Find(&Product).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return Product, nil
}

// Generate Id
func generateProductId(latestID string) string {
	var numID int
	if _, err := fmt.Sscanf(latestID, "PDT-%04d", &numID); err != nil {
		return ""
	}
	numID++
	return fmt.Sprintf("PDT-%04d", numID)
}

// Validate
func validateProduct(product Products, db *gorm.DB) bool {
	if product.Id == "" || len(product.Id) > 10 {
		logrus.Error("Model: Id is required and must be up to 10 characters")
		return false
	}

	if product.Name == "" || len(product.Name) > 100 {
		logrus.Error("Model: Id Employee is required and must be up to 10 characters")
		return false
	}

	if product.Photo == "" {
		logrus.Error("Model: Id Employee is required and must be up to 10 characters")
		return false
	}

	if product.IdCatProduct == "" || len(product.IdCatProduct) > 10 {
		logrus.Error("Model: Id Employee is required and must be up to 10 characters")
		return false
	}

	var catProduct CatProducts
	if err := db.Where("id = ?", product.IdCatProduct).First(&catProduct).Error; err != nil {
		logrus.Error("Model: Employee with the specified ID does not exist")
		return false
	}

	if product.MfDate.IsZero() {
		logrus.Error("Model: Product name is required and must be up to 100 characters")
		return false
	}

	if product.MfDate.IsZero() {
		logrus.Error("Model: Quantity is required")
		return false
	}

	if product.BatchNumber == 0 {
		logrus.Error("Model: Note is required")
		return false
	}

	if product.UnitPrice == 0 {
		logrus.Error("Model: Note is required")
		return false
	}

	if product.Stock == 0 {
		logrus.Error("Model: Note is required")
		return false
	}

	if product.Description == "" {
		logrus.Error("Model: Status is required")
		return false
	}

	if product.IdDistributor == "" || len(product.IdDistributor) > 10 {
		logrus.Error("Model: Id Employee is required and must be up to 10 characters")
		return false
	}

	var distributor Distributors
	if err := db.Where("id = ?", product.IdDistributor).First(&distributor).Error; err != nil {
		logrus.Error("Model: Employee with the specified ID does not exist")
		return false
	}

	return true
}
