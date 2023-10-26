package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Struct Category Products
type CatProducts struct {
	Id        string         `gorm:"primaryKey;type:varchar(10)"`
	Name      string         `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Product   []Products     `gorm:"foreignKey:id_cat_product;references:id"`
}

// Interface beetween models and controller
type CatProductsModelInterface interface {
	Insert(newCatProduct CatProducts) (*CatProducts, error)
	SelectAll(limit, offset int) ([]CatProducts, error)
	Update(updatedData CatProducts) (*CatProducts, error)
	Delete(catProductId string) (bool, error)
	SearchCatProduct(keyword string, limit int, offset int) ([]CatProducts, error)
}

// connect into db
type CatProductsModel struct {
	db *gorm.DB
}

// New Instance from UsersModel
func NewCatProductsModel(db *gorm.DB) CatProductsModelInterface {
	return &CatProductsModel{
		db: db,
	}
}

// Insert Category Product
func (cpm *CatProductsModel) Insert(newCatProduct CatProducts) (*CatProducts, error) {
	var latestCatProduct CatProducts
	if errSort := cpm.db.Order("id DESC").First(&latestCatProduct).Error; errSort != nil {
    	latestCatProduct.Id = "CPT-0000"
	}

	newID := generateCatProductId(latestCatProduct.Id)

	if newID == "" {
    	return nil, errors.New("Failed generate Id")
	}

	newCatProduct.Id = newID

	validate := validateCatProduct(newCatProduct)
	if !validate {
		return nil, errors.New("Data not valid")
	}

	if checkName := cpm.db.Where("name = ?", newCatProduct.Name).First(&newCatProduct).Error; checkName != nil {
		if err := cpm.db.Create(&newCatProduct).Error; err != nil {
			return nil, errors.New("Error insert category product, " + err.Error())
		}
	} else {
		return nil, errors.New("Username already exists")
	}

	return &newCatProduct, nil
}

// Select All Category Product
func (cpm *CatProductsModel) SelectAll(limit, offset int) ([]CatProducts, error) {
	var data = []CatProducts{}
	if err := cpm.db.Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get all category product, " + err.Error())
	}

	return data, nil
}

// Update Category Product
func (cpm *CatProductsModel) Update(updatedData CatProducts) (*CatProducts, error) {
	var qry = cpm.db.Table("cat_products").Where("id = ?", updatedData.Id).Update("name", updatedData.Name)
	if err := qry.Error; err != nil {
		return nil, errors.New("Update error, " + err.Error())
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return nil, errors.New("Update error, no data effected")
	}

	var updatedCatProduct = CatProducts{}
	if err := cpm.db.Where("id = ?", updatedData.Id).First(&updatedCatProduct).Error; err != nil {
		logrus.Error("Model : Error get updated data, ", err.Error())
		return nil, errors.New("Error get updated data, " + err.Error())
	}

	return &updatedCatProduct, nil
}

// Delete Category Product
func (cpm *CatProductsModel) Delete(catProductId string) (bool, error) {
	var data = CatProducts{}
	data.Id = catProductId

	if err := cpm.db.Where("id = ?", catProductId).First(&data).Error; err != nil {
		return false, errors.New("Error finding data to delete, " + err.Error())
	}

	if err := cpm.db.Delete(&data).Error; err != nil {
		return false, errors.New("Error delete data, " + err.Error())
	}

	return true, nil
}

// Searching
func (cpm *CatProductsModel) SearchCatProduct(keyword string, limit int, offset int) ([]CatProducts, error) {
	var catProduct []CatProducts
	query := cpm.db.Where("id LIKE ? OR name LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Limit(limit).Offset(offset)

	if err := query.Find(&catProduct).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return catProduct, nil
}

// Generate Id
func generateCatProductId(latestID string) string {
	var numID int
	if _, err := fmt.Sscanf(latestID, "CPT-%04d", &numID); err != nil {
		return ""
	}
	numID++
	return fmt.Sprintf("CPT-%04d", numID)
}

// Validate
func validateCatProduct(catProduct CatProducts) bool {
	if catProduct.Id == "" || len(catProduct.Id) > 10 {
		logrus.Error("Model: Id is required and must be up to 10 characters")
		return false
	}
	if catProduct.Name == "" || len(catProduct.Name) > 100 {
		logrus.Error("Model: Product name is required and must be up to 100 characters")
		return false
	}

	return true
}