package models

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Distributor struct {
	Id             int             `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Name           string          `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	CreatedAt      time.Time       `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt      time.Time       `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt      gorm.DeletedAt  `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DetailProducts []DetailProduct `gorm:"foreignKey:IdDistributor;references:Id"`
}

type DistributorModelInterface interface {
	Insert(newDistributor Distributor) *Distributor
	SelectAll() []Distributor
	SelectById(distributorId int) *Distributor
	Update(updatedData Distributor) *Distributor
	Delete(DistributorId int) bool
}

type DistributorsModel struct {
	db *gorm.DB
}

func NewDistributorsModel(db *gorm.DB) DistributorModelInterface {
	return &DistributorsModel{
		db: db,
	}
}

func (dm *DistributorsModel) Insert(newDistributor Distributor) *Distributor {
	if err := dm.db.Create(&newDistributor).Error; err != nil {
		logrus.Error("Model : Insert data error, ", err.Error())
		return nil
	}

	return &newDistributor
}

func (dm *DistributorsModel) SelectAll() []Distributor {
	var data = []Distributor{}
	if err := dm.db.Find(&data).Error; err != nil {
		logrus.Error("Model : Cannot get all distributor, ", err.Error())
		return nil
	}

	return data
}

func (dm *DistributorsModel) SelectById(distributorId int) *Distributor {
	var data = Distributor{}
	if err := dm.db.Where("id = ?", distributorId).First(&data).Error; err != nil {
		logrus.Error("Model : Data with that ID was not found, ", err.Error())
		return nil
	}

	return &data
}

func (dm *DistributorsModel) Update(updatedData Distributor) *Distributor {
	var qry = dm.db.Table("distributors").Where("id = ?", updatedData.Id).Update("name", updatedData.Name)
	if err := qry.Error; err != nil {
		logrus.Error("Model : update error, ", err.Error())
		return nil
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update error, ", "no data effected")
		return nil
	}

	var updatedDistributor = Distributor{}
	if err := dm.db.Where("id = ?", updatedData.Id).First(&updatedDistributor).Error; err != nil {
		logrus.Error("Model : Error get updated data, ", err.Error())
		return nil
	}

	return &updatedDistributor
}

func (dm *DistributorsModel) Delete(DistributorId int) bool {
	var data = Distributor{}
	data.Id = DistributorId

	if err := dm.db.Where("id = ?", DistributorId).First(&data).Error; err != nil {
		logrus.Error("Model: Error finding data to delete, ", err.Error())
		return false
	}

	if err := dm.db.Delete(&data).Error; err != nil {
		logrus.Error("Model : Error delete data, ", err.Error())
		return false
	}

	return true
}
