package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Struct Distributor
type Distributors struct {
	Id        string         `gorm:"primaryKey;type:varchar(10)" json:"id" form:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	CreatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Products  []Products     `gorm:"foreignKey:id_distributor;references:Id" json:"product" form:"product"`
}

// Interface beetween models and controller
type DistributorModelInterface interface {
	Insert(newDistributor Distributors) (*Distributors, error)
	SelectAll(limit int, offset int) ([]Distributors, error)
	Update(updatedData Distributors) (*Distributors, error)
	Delete(DistributorId string) (bool, error)
	SearchDistributor(keyword string, limit int, offset int) ([]Distributors, error)
}

// Connect into db
type DistributorsModel struct {
	db *gorm.DB
}

// New Instance from DistributorsModel
func NewDistributorsModel(db *gorm.DB) DistributorModelInterface {
	return &DistributorsModel{
		db: db,
	}
}

// Insert Distributor
func (dm *DistributorsModel) Insert(newDistributor Distributors) (*Distributors, error) {
	var latestDistributor Distributors
	if errSort := dm.db.Unscoped().Order("id DESC").First(&latestDistributor).Error; errSort != nil {
		latestDistributor.Id = "DST-0000"
	}

	newID := generateDistributorId(latestDistributor.Id)
	if newID == "" {
		return nil, errors.New("Failed generate Id")
	}

	newDistributor.Id = newID

	validate, errValidate := validateDistributor(newDistributor)
	if !validate {
		return nil, errValidate
	}

	if checkName := dm.db.Where("name = ?", newDistributor.Name).First(&newDistributor).Error; checkName != nil {
		if checkName == gorm.ErrRecordNotFound {
			if err := dm.db.Create(&newDistributor).Error; err != nil {
				return nil, errors.New("Error insert distributor, " + err.Error())
			}
		} else {
			return nil, errors.New("Error checking distributor name availability")
		}
	} else {
		return nil, errors.New("Distributor already exists")
	}

	return &newDistributor, nil
}

// Select All Distributor
func (dm *DistributorsModel) SelectAll(limit int, offset int) ([]Distributors, error) {
	var data = []Distributors{}

	if err := dm.db.
		Limit(limit).
		Offset(offset).
		Preload("Products").
		Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get distributor with products, " + err.Error())
	}

	return data, nil
}

// Update Distributor
func (dm *DistributorsModel) Update(updatedData Distributors) (*Distributors, error) {
	var qry = dm.db.Table("distributors").Where("id = ?", updatedData.Id).Update("name", updatedData.Name)
	if err := qry.Error; err != nil {
		return nil, errors.New("Update data error, " + err.Error())
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return nil, errors.New("Update data error, no data effected")
	}

	var updatedDistributor = Distributors{}
	if err := dm.db.Where("id = ?", updatedData.Id).First(&updatedDistributor).Error; err != nil {
		return nil, errors.New("Error get updated data, " + err.Error())
	}

	return &updatedDistributor, nil
}

// Delete Distributor
func (dm *DistributorsModel) Delete(DistributorId string) (bool, error) {
	var data = Distributors{}
	data.Id = DistributorId

	if err := dm.db.Where("id = ?", DistributorId).First(&data).Error; err != nil {
		return false, errors.New("Error finding data to delete, " + err.Error())
	}

	if err := dm.db.Delete(&data).Error; err != nil {
		return false, errors.New("Error delete data, " + err.Error())
	}

	return true, nil
}

// Searching
func (dm *DistributorsModel) SearchDistributor(keyword string, limit int, offset int) ([]Distributors, error) {
	var distributor []Distributors
	query := dm.db.Limit(limit).Offset(offset).Where("id LIKE ? OR name LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.
		Limit(limit).
		Offset(offset).
		Preload("Products").
		Find(&distributor).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return distributor, nil
}

// Generate Id
func generateDistributorId(latestID string) string {
	var numID int
	if _, err := fmt.Sscanf(latestID, "DST-%04d", &numID); err != nil {
		return ""
	}
	numID++
	return fmt.Sprintf("DST-%04d", numID)
}

// Validate
func validateDistributor(distributor Distributors) (bool, error) {
	if distributor.Id == "" || len(distributor.Id) > 10 {
		return false, errors.New("Id is required and must be up to 10 characters")
	}
	if distributor.Name == "" || len(distributor.Name) > 100 {
		return false, errors.New("Distributor name is required and must be up to 100 characters")
	}

	return true, nil
}
