package models

import (
	"errors"
	"fmt"
	"time"

	// "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Struct Request Product
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

// Interface beetween models and controller
type ReqProductsModelInterface interface {
	Insert(newReqProduct ReqProducts) (*ReqProducts, error)
	SelectAll(limit, offset int) ([]ReqProducts, error)
	Update(updatedData ReqProducts) (*ReqProducts, error)
	Delete(reqProductId string) (bool, error)
	SearchReqProduct(keyword string, limit int, offset int) ([]ReqProducts, error)
}

// connect into db
type ReqProductsModel struct {
	db *gorm.DB
}

// New Instance from ReqProductsModel
func NewReqProductsModel(db *gorm.DB) ReqProductsModelInterface {
	return &ReqProductsModel{
		db: db,
	}
}

// Insert Request Product
func (rpm *ReqProductsModel) Insert(newReqProduct ReqProducts) (*ReqProducts, error) {
	var latestReqProduct ReqProducts
	if errSort := rpm.db.Unscoped().Order("id DESC").First(&latestReqProduct).Error; errSort != nil {
		latestReqProduct.Id = "RQP-0000"
	}

	newID := generateReqProductId(latestReqProduct.Id)
	if newID == "" {
		return nil, errors.New("Failed generate Id")
	}

	newReqProduct.Id = newID
	newReqProduct.StatusReq = "processed"

	validate := validateReqProduct(newReqProduct, rpm.db)
	if !validate {
		return nil, errors.New("Data not valid")
	}

	if err := rpm.db.Create(&newReqProduct).Error; err != nil {
		return nil, errors.New("Error insert request product, " + err.Error())
	}

	return &newReqProduct, nil
}

// Select All Request Product 
func (rpm *ReqProductsModel) SelectAll(limit, offset int) ([]ReqProducts, error) {
	var data []ReqProducts
	if err := rpm.db.Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get all request product, " + err.Error())
	}

	return data, nil
}

// Update Request Product
func (rpm *ReqProductsModel) Update(updatedData ReqProducts) (*ReqProducts, error) {
	var data map[string]interface{} = make(map[string]interface{})

	if updatedData.ProductName != "" {
		data["product_name"] = updatedData.ProductName
	}
	if updatedData.Quantity != 0 {
		data["quantity"] = updatedData.Quantity
	}
	if updatedData.Note != "" {
		data["note"] = updatedData.Note
	}
	if updatedData.StatusReq != "" {
		data["status_req"] = updatedData.StatusReq
	}

	var qry = rpm.db.Table("req_products").Where("id = ?", updatedData.Id).Updates(data)
	if err := qry.Error; err != nil {
		return nil, errors.New("Error update data" + err.Error())
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return nil, errors.New("Error update data, no data affected")
	}

	var updatedReqProduct = ReqProducts{}
	if err := rpm.db.Where("id = ?", updatedData.Id).First(&updatedReqProduct).Error; err != nil {
		return nil, errors.New("Cannot get updated data" + err.Error())
	}

	return &updatedReqProduct, nil
}

// Delete Request Product
func (rpm *ReqProductsModel) Delete(reqProductId string) (bool, error) {
	var data = ReqProducts{}
	data.Id = reqProductId

	if err := rpm.db.Where("id = ?", reqProductId).First(&data).Error; err != nil {
		return false, errors.New("Error finding data to delete, " + err.Error())
	}

	if err := rpm.db.Delete(&data).Error; err != nil {
		return false, errors.New("Error delete data, " + err.Error())
	}

	return true, nil
}

// Searching
func (rpm *ReqProductsModel) SearchReqProduct(keyword string, limit int, offset int) ([]ReqProducts, error) {
	var reqProduct []ReqProducts
	query := rpm.db.Where("id LIKE ? OR id_employee LIKE ? OR product_name LIKE ? OR quantity LIKE ? OR note LIKE ? OR status_req LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Limit(limit).Offset(offset)

	if err := query.Find(&reqProduct).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return reqProduct, nil
}

// Generate Id
func generateReqProductId(latestID string) string {
	var numID int
	if _, err := fmt.Sscanf(latestID, "RQP-%04d", &numID); err != nil {
		return ""
	}
	numID++
	return fmt.Sprintf("RQP-%04d", numID)
}

// Validate
func validateReqProduct(reqProduct ReqProducts, db *gorm.DB) bool {
	if reqProduct.Id == "" || len(reqProduct.Id) > 10 {
		logrus.Error("Model: Id is required and must be up to 10 characters")
		return false
	}

	if reqProduct.IdEmployee == "" || len(reqProduct.IdEmployee) > 10 {
		logrus.Error("Model: Id Employee is required and must be up to 10 characters")
		return false
	}

	var user Users
	if err := db.Where("id = ?", reqProduct.IdEmployee).First(&user).Error; err != nil {
		logrus.Error("Model: Employee with the specified ID does not exist")
		return false
	}

	if reqProduct.ProductName == "" || len(reqProduct.ProductName) > 100 {
		logrus.Error("Model: Product name is required and must be up to 100 characters")
		return false
	}

	if reqProduct.Quantity == 0 {
		logrus.Error("Model: Quantity is required")
		return false
	}

	if reqProduct.Note == "" {
		logrus.Error("Model: Note is required")
		return false
	}

	if reqProduct.StatusReq == "" || (reqProduct.StatusReq != "rejected" && reqProduct.StatusReq != "processed" && reqProduct.StatusReq != "accepted") {
		logrus.Error("Model: Status request is required")
		return false
	}

	return true
}

