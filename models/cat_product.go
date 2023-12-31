package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Struct Category Products
type CatProducts struct {
	Id        string         `gorm:"primaryKey;type:varchar(10)" json:"id" form:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	CreatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Products  []Products     `gorm:"foreignKey:id_cat_product;references:id" json:"product" form:"product"`
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

// New Instance from CatProductsModel
func NewCatProductsModel(db *gorm.DB) CatProductsModelInterface {
	return &CatProductsModel{
		db: db,
	}
}

// Insert Category Product
func (cpm *CatProductsModel) Insert(newCatProduct CatProducts) (*CatProducts, error) {
	var existingCatProduct CatProducts
	if err := cpm.db.Where("name = ?", newCatProduct.Name).First(&existingCatProduct).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.New("Error checking category product availability")
		}
	}

	if existingCatProduct.Id == "" {
		tx := cpm.db.Begin()
		var latestCatProduct CatProducts
		if errSort := tx.Unscoped().Order("id DESC").First(&latestCatProduct).Error; errSort != nil {
			latestCatProduct.Id = "CPT-0000"
		}

		newID := generateCatProductId(latestCatProduct.Id)
		if newID == "" {
			tx.Rollback()
			return nil, errors.New("Failed to generate ID")
		}

		newCatProduct.Id = newID
		
		validate, errValidate := validateCatProduct(newCatProduct)
		if !validate {
			tx.Rollback()
			return nil, errValidate
		}

		if err := tx.Create(&newCatProduct).Error; err != nil {
			tx.Rollback()
			return nil, errors.New("Error inserting category product, " + err.Error())
		}

		tx.Commit()
		return &newCatProduct, nil
	}

	return nil, errors.New("Category product already exists")
}

// Select All Category Product
func (cpm *CatProductsModel) SelectAll(limit, offset int) ([]CatProducts, error) {
	var data = []CatProducts{}

	if err := cpm.db.
		Limit(limit).
		Offset(offset).
		Preload("Products").
		Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get category product with products, " + err.Error())
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
	query := cpm.db.Limit(limit).Offset(offset).Where("id LIKE ? OR name LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.
		Preload("Products").
		Find(&catProduct).Error; err != nil {
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
func validateCatProduct(catProduct CatProducts) (bool, error) {
	if catProduct.Id == "" || len(catProduct.Id) > 10 {
		return false, errors.New("Id is required and must be up to 10 characters")
	}
	if catProduct.Name == "" || len(catProduct.Name) > 100 {
		return false, errors.New("Name is required and must be up to 100 characters")
	}

	return true, nil
}
