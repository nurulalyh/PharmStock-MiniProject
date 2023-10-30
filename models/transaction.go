package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Struct Transactions
type Transactions struct {
	Id                string               `gorm:"primaryKey;type:varchar(10)" json:"id" form:"id"`
	IdEmployee        string               `gorm:"type:varchar(10);not null" json:"id_employee" form:"id_employee"`
	TotalQuantity     int                  `gorm:"type:smallint;not null" json:"total_quantity" form:"total_quantity"`
	TotalPrice        int                  `gorm:"type:smallint;not null" json:"total_price" form:"total_price"`
	Type              string               `gorm:"type:ENUM('inbound','outbound');not null" json:"type" form:"type"`
	CreatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt         gorm.DeletedAt       `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DetailTransaction []DetailTransactions `gorm:"foreignKey:id_transaction;references:id" json:"detail" form:"detail"`
}

// Interface beetween models and controller
type TransactionsModelInterface interface {
	Insert(newTransaction Transactions) (*Transactions, error)
	SelectAll(limit int, offset int) ([]Transactions, error)
	Update(updatedData Transactions) (*Transactions, error)
	Delete(TransactionId string) (bool, error)
	SearchTransaction(keyword string, limit int, offset int) ([]Transactions, error)
}

// Connect into db
type TransactionsModel struct {
	db *gorm.DB
}

// New Instance from TransactionsModel
func NewTransactionsModel(db *gorm.DB) TransactionsModelInterface {
	return &TransactionsModel{
		db: db,
	}
}

// Insert Transaction
func (tm *TransactionsModel) Insert(newTransaction Transactions) (*Transactions, error) {
	var latestTransaction Transactions
	if errSort := tm.db.Unscoped().Order("id DESC").First(&latestTransaction).Error; errSort != nil {
		latestTransaction.Id = "TRN-0000"
	}

	newID := generateTransactionId(latestTransaction.Id)

	if newID == "" {
		return nil, errors.New("Failed generate Id")
	}

	newTransaction.Id = newID
	newTransaction.TotalQuantity = 0
	newTransaction.TotalPrice = 0

	validate := validateTransaction(newTransaction, tm.db)
	if !validate {
		return nil, errors.New("Data not valid")
	}

	if err := tm.db.Create(&newTransaction).Error; err != nil {
		return nil, errors.New("Error insert Transaction, " + err.Error())
	}

	return &newTransaction, nil
}

// Select All transaction
func (tm *TransactionsModel) SelectAll(limit int, offset int) ([]Transactions, error) {
	var data = []Transactions{}
	if err := tm.db.Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get all transaction, " + err.Error())
	}

	return data, nil
}

// Update transaction
func (tm *TransactionsModel) Update(updatedData Transactions) (*Transactions, error) {
	var data map[string]interface{} = make(map[string]interface{})

	if updatedData.IdEmployee != "" {
		data["id_employee"] = updatedData.IdEmployee
	}
	if updatedData.Type != "" {
		data["type"] = updatedData.Type
	}

	var qry = tm.db.Table("transactions").Where("id = ?", updatedData.Id).Updates(data)
	if err := qry.Error; err != nil {
		return nil, errors.New("Error update data" + err.Error())
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return nil, errors.New("Update error, no data effected")
	}

	var updatedTransaction = Transactions{}
	if err := tm.db.Where("id = ?", updatedData.Id).First(&updatedTransaction).Error; err != nil {
		return nil, errors.New("Error get updated data, " + err.Error())
	}

	return &updatedTransaction, nil
}

// Delete Distributor
func (tm *TransactionsModel) Delete(transactionId string) (bool, error) {
	var data = Transactions{}
	data.Id = transactionId

	if err := tm.db.Where("id = ?", transactionId).First(&data).Error; err != nil {
		return false, errors.New("Error finding data to delete, " + err.Error())
	}

	if err := tm.db.Delete(&data).Error; err != nil {
		return false, errors.New("Error delete data, " + err.Error())
	}

	return true, nil
}

// Searching
func (tm *TransactionsModel) SearchTransaction(keyword string, limit int, offset int) ([]Transactions, error) {
	var transaction []Transactions
	query := tm.db.Where("id LIKE ? OR id_employee LIKE ? OR total_quantity LIKE ? OR total_price LIKE ? OR type LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Limit(limit).Offset(offset)

	if err := query.Find(&transaction).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return transaction, nil
}

// Generate Id
func generateTransactionId(latestID string) string {
	var numID int
	if _, err := fmt.Sscanf(latestID, "TRN-%04d", &numID); err != nil {
		return ""
	}
	numID++
	return fmt.Sprintf("TRN-%04d", numID)
}

// Validate
func validateTransaction(transaction Transactions, db *gorm.DB) bool {
	if transaction.Id == "" || len(transaction.Id) > 10 {
		logrus.Error("Model: Id is required and must be up to 10 characters")
		return false
	}

	if transaction.IdEmployee == "" || len(transaction.IdEmployee) > 10 {
		logrus.Error("Model: Id Employee is required and must be up to 10 characters")
		return false
	}

	var user Users
	if err := db.Where("id = ?", transaction.IdEmployee).First(&user).Error; err != nil {
		logrus.Error("Model: Employee with the specified ID does not exist")
		return false
	}

	if transaction.Type == "" || (transaction.Type != "inbound" && transaction.Type != "outbound") {
		logrus.Error("Model: Status request is required")
		return false
	}

	return true
}
