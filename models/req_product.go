package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Struct Request Product
type ReqProducts struct {
	Id          string         `gorm:"primaryKey;type:varchar(10)" json:"id" form:"id"`
	IdEmployee  string         `gorm:"type:varchar(10);not null" json:"id_employee" form:"id_employee"`
	ProductName string         `gorm:"type:varchar(100);not null" json:"product_name" form:"product_name"`
	Quantity    int            `gorm:"type:smallint;not null" json:"quantity" form:"quantity"`
	Note        string         `gorm:"type:text;not null" json:"note" form:"note"`
	StatusReq   string         `gorm:"type:ENUM('rejected','processed', 'accepted');not null" json:"status_request" form:"status_request"`
	CreatedAt   time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
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

	validate, errValidate := validateReqProduct(newReqProduct, rpm.db)
	if !validate {
		return nil, errValidate
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
	if updatedData.StatusReq != "rejected" && updatedData.StatusReq != "processed" && updatedData.StatusReq != "accepted" {
		return nil, errors.New("input does not match the format")
	}

	var qry = rpm.db.Table("req_products").Where("id = ?", updatedData.Id).Update("status_req", updatedData.StatusReq)
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
	query := rpm.db.Limit(limit).Offset(offset).Where("id LIKE ? OR id_employee LIKE ? OR product_name LIKE ? OR quantity LIKE ? OR note LIKE ? OR status_req LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

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
func validateReqProduct(reqProduct ReqProducts, db *gorm.DB) (bool, error) {
	if reqProduct.Id == "" || len(reqProduct.Id) > 10 {
		return false, errors.New("Id is required and must be up to 10 characters")
	}

	if reqProduct.IdEmployee == "" || len(reqProduct.IdEmployee) > 10 {
		return false, errors.New("Id employee is required and must be up to 10 characters")
	}

	var user Users
	if err := db.Where("id = ?", reqProduct.IdEmployee).First(&user).Error; err != nil {
		return false, errors.New("Id employee is not registered")
	}

	if reqProduct.ProductName == "" || len(reqProduct.ProductName) > 100 {
		return false, errors.New("Product name is required and must be up to 100 characters")
	}

	if reqProduct.Quantity == 0 {
		return false, errors.New("Quantity is required")
	}

	if reqProduct.Note == "" {
		return false, errors.New("Note is required")
	}

	if reqProduct.StatusReq == "" || (reqProduct.StatusReq != "rejected" && reqProduct.StatusReq != "processed" && reqProduct.StatusReq != "accepted") {
		return false, errors.New("Status request is required and must be processed or rejected or accepted")
	}

	return true, nil
}
