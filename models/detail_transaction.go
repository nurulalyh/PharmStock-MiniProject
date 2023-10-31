package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Struct Detail Transaction
type DetailTransactions struct {
	Id            string         `gorm:"primaryKey;type:varchar(10)"`
	IdTransaction string         `gorm:"type:varchar(10);not null"`
	IdProduct     string         `gorm:"type:varchar(10);not null"`
	Quantity      int            `gorm:"type:smallint;not null"`
	Price         int            `gorm:"type:smallint;not null"`
	CreatedAt     time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Interface beetween models and controller
type DetailTransactionsModelInterface interface {
	Insert(newDetailTransaction DetailTransactions) (*DetailTransactions, error)
	SelectAll(limit, offset int) ([]DetailTransactions, error)
	Update(updatedData DetailTransactions) (*DetailTransactions, error)
	Delete(DetailTransactionId string) (bool, error)
	SearchDetailTransaction(keyword string, limit int, offset int) ([]DetailTransactions, error)
}

// connect into db
type DetailTransactionsModel struct {
	db *gorm.DB
}

// New Instance from DetailTransactionsModel
func NewDetailTransactionsModel(db *gorm.DB) DetailTransactionsModelInterface {
	return &DetailTransactionsModel{
		db: db,
	}
}

// Insert Detail Transactions
func (dtm *DetailTransactionsModel) Insert(newDetailTransaction DetailTransactions) (*DetailTransactions, error) {
	var latestDetailTransaction DetailTransactions
	if errSort := dtm.db.Unscoped().Order("id DESC").First(&latestDetailTransaction).Error; errSort != nil {
		latestDetailTransaction.Id = "DTL-0000"
	}

	newID := generateDetailTransactionId(latestDetailTransaction.Id)
	if newID == "" {
		return nil, errors.New("Failed generate Id")
	} 

	newDetailTransaction.Id = newID
	newDetailTransaction.Price = 0

	validate := validateDetailTransaction(newDetailTransaction, dtm.db)
	if !validate {
		return nil, errors.New("Data not valid")
	}

	if err := dtm.db.Create(&newDetailTransaction).Error; err != nil {
		return nil, errors.New("Error insert Detail Transactions, " + err.Error())
	}

	return &newDetailTransaction, nil
}

// Select All Detail Transactions
func (dtm *DetailTransactionsModel) SelectAll(limit, offset int) ([]DetailTransactions, error) {
	var data []DetailTransactions
	if err := dtm.db.Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get all Detail Transactions, " + err.Error())
	}

	return data, nil
}

// Update Detail Transactions
func (dtm *DetailTransactionsModel) Update(updatedData DetailTransactions) (*DetailTransactions, error) {
	var data map[string]interface{} = make(map[string]interface{})

	if updatedData.IdTransaction != "" {
		data["id_transaction"] = updatedData.IdTransaction
	}
	if updatedData.IdProduct != "" {
		data["id_product"] = updatedData.IdProduct
	}
	if updatedData.Quantity != 0 {
		data["quantity"] = updatedData.Quantity
	}

	var qry = dtm.db.Table("detail_transactions").Where("id = ?", updatedData.Id).Updates(data)
	if err := qry.Error; err != nil {
		return nil, errors.New("Error update data" + err.Error())
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return nil, errors.New("Error update data, no data affected")
	}

	var updatedDetailTransactions = DetailTransactions{}
	if err := dtm.db.Where("id = ?", updatedData.Id).First(&updatedDetailTransactions).Error; err != nil {
		return nil, errors.New("Cannot get updated data" + err.Error())
	}

	return &updatedDetailTransactions, nil
}

// Delete Detail Transactions
func (dtm *DetailTransactionsModel) Delete(detailTransactionId string) (bool, error) {
	var data = DetailTransactions{}
	data.Id = detailTransactionId

	if err := dtm.db.Where("id = ?", detailTransactionId).First(&data).Error; err != nil {
		return false, errors.New("Error finding data to delete, " + err.Error())
	}

	if err := dtm.db.Delete(&data).Error; err != nil {
		return false, errors.New("Error delete data, " + err.Error())
	}

	return true, nil
}

// Searching
func (dtm *DetailTransactionsModel) SearchDetailTransaction(keyword string, limit int, offset int) ([]DetailTransactions, error) {
	var detailTransaction []DetailTransactions
	query := dtm.db.Where("id LIKE ? OR id_transaction LIKE ? OR id_product LIKE ? OR quantity LIKE ? OR price LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Limit(limit).Offset(offset)

	if err := query.Find(&detailTransaction).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return detailTransaction, nil
}

// Hook After Create to manage product stock and detail_transaction quantity
func (dtm *DetailTransactionsModel) AfterCreate(dt *DetailTransactions) (err error) {
    if err := dtm.db.Table("Products").Where("id = ?", dt.IdProduct).UpdateColumn("stock", gorm.Expr("stock - ?", dt.Quantity)).Error; err != nil {
		return err
	}
    return nil
}

func (dt *DetailTransactions) AfterCreate(tx *gorm.DB) (err error) {
    if err := tx.Model(&Products{}).Where("id = ?", dt.IdProduct).UpdateColumn("stock", gorm.Expr("stock - ?", dt.Quantity)).Error; err != nil {
		return err
	}
    return nil
}


// Generate Id
func generateDetailTransactionId(latestID string) string {
	var numID int
	if _, err := fmt.Sscanf(latestID, "DTL-%04d", &numID); err != nil {
		return ""
	}
	numID++
	return fmt.Sprintf("DTL-%04d", numID)
}

// Validate
func validateDetailTransaction(detailTransaction DetailTransactions, db *gorm.DB) bool {
	if detailTransaction.Id == "" || len(detailTransaction.Id) > 10 {
		logrus.Error("Model: Id is required and must be up to 10 characters")
		return false
	}

	if detailTransaction.IdTransaction == "" || len(detailTransaction.IdTransaction) > 10 {
		logrus.Error("Model: Id is required and must be up to 10 characters")
		return false
	}

	var transaction Transactions
	if err := db.Where("id = ?", detailTransaction.IdTransaction).First(&transaction).Error; err != nil {
		logrus.Error("Model: Employee with the specified ID does not exist")
		return false
	}

	if detailTransaction.IdProduct == "" || len(detailTransaction.IdProduct) > 10 {
		logrus.Error("Model: Id is required and must be up to 10 characters")
		return false
	}

	var product Products
	if err := db.Where("id = ?", detailTransaction.IdProduct).First(&product).Error; err != nil {
		logrus.Error("Model: Employee with the specified ID does not exist")
		return false
	}

	if detailTransaction.Quantity == 0 {
		logrus.Error("Model: Quantity is required")
		return false
	}

	if detailTransaction.Price == 0 {
		logrus.Error("Model: Note is required")
		return false
	}
	
	return true
}
