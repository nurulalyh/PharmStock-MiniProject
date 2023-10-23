package models

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Distributor struct {
	Id        int            `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	CreatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DetailProducts []DetailProduct `gorm:"foreignKey:IdDistributor;references:Id"`
}

type DistributorModelInterface interface {
	Insert(newDistributor Distributor) *Distributor	
	SelectAll() []Distributor
	SelectById(distributorId int) *Distributor
}

type DistributorsModel struct {
	db *gorm.DB
}

func NewDistributorsModel(db *gorm.DB) DistributorModelInterface {
	return &DistributorsModel {
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