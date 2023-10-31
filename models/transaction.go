package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Struct Transactions
type Transactions struct {
	Id                 string               `gorm:"primaryKey;type:varchar(10)" json:"id" form:"id"`
	IdEmployee         string               `gorm:"type:varchar(10);not null" json:"id_employee" form:"id_employee"`
	TotalQuantity      int                  `gorm:"type:smallint;not null" json:"total_quantity" form:"total_quantity"`
	TotalPrice         int                  `gorm:"type:smallint;not null" json:"total_price" form:"total_price"`
	Type               string               `gorm:"type:ENUM('inbound','outbound');not null" json:"type" form:"type"`
	CreatedAt          time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt          time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt          gorm.DeletedAt       `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DetailTransactions []DetailTransactions `gorm:"foreignKey:id_transaction;references:id" json:"detail" form:"detail"`
}

// Interface beetween models and controller
type TransactionsModelInterface interface {
	Insert(newTransaction Transactions) (*Transactions, error)
	SelectAll(limit int, offset int) ([]Transactions, error)
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

	validate, errValidate := validateTransaction(newTransaction, tm.db)
	if !validate {
		return nil, errValidate
	}

	if err := tm.db.Create(&newTransaction).Error; err != nil {
		return nil, errors.New("Error insert Transaction, " + err.Error())
	}

	return &newTransaction, nil
}

// Select All transaction
func (tm *TransactionsModel) SelectAll(limit int, offset int) ([]Transactions, error) {
	var data = []Transactions{}

	if err := tm.db.
		Limit(limit).
		Offset(offset).
		Preload("DetailTransactions").
		Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get transaction with detail_transactions, " + err.Error())
	}

	return data, nil
}

// Searching
func (tm *TransactionsModel) SearchTransaction(keyword string, limit int, offset int) ([]Transactions, error) {
	var transaction []Transactions
	query := tm.db.Limit(limit).Offset(offset).Where("id LIKE ? OR id_employee LIKE ? OR total_quantity LIKE ? OR total_price LIKE ? OR type LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.
		Preload("DetailTransactions").
		Find(&transaction).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return transaction, nil
}

// After Create to acumulate price and quantity from detail_transaction
func (tm *TransactionsModel) AfterCreate(t *Transactions) error {
	var totalQuantity int
	var totalPrice int
	
	if errQuery := tm.db.Table("detail_transactions").Where("id_transaction = ?", t.Id).Select("SUM(quantity) as total_quantity, SUM(price) as total_price").Row().Scan(&totalQuantity, &totalPrice); errQuery != nil {
		return errQuery
	}
	
	t.TotalQuantity = totalQuantity
	t.TotalPrice = totalPrice

	if err := tm.db.Table("transactions").Where("id = ?", t.Id).UpdateColumns(map[string]interface{}{
		"total_quantity": t.TotalQuantity,
		"total_price":    t.TotalPrice,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (t *Transactions) AfterCreate(tx *gorm.DB) (err error) {
    var totalQuantity int
    var totalPrice int

    if errQuery := tx.Model(&DetailTransactions{}).Where("id_transaction = ?", t.Id).Select("SUM(quantity) as total_quantity, SUM(price) as total_price").Row().Scan(&totalQuantity, &totalPrice); errQuery == nil {
		return errQuery
	}

	t.TotalQuantity = totalQuantity
	t.TotalPrice = totalPrice

	if err := tx.Model(&Transactions{}).Where("id = ?", t.Id).UpdateColumns(map[string]interface{}{
	"TotalQuantity": t.TotalQuantity,
	"TotalPrice":    t.TotalPrice,
	}).Error; err != nil {
		return err
	}

	return nil
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
func validateTransaction(transaction Transactions, db *gorm.DB) (bool, error) {
	if transaction.Id == "" || len(transaction.Id) > 10 {
		return false, errors.New("Id is required and must be up to 10 characters")
	}

	if transaction.IdEmployee == "" || len(transaction.IdEmployee) > 10 {
		return false, errors.New("Id employee is required and must be up to 10 characters")
	}

	var user Users
	if err := db.Where("id = ?", transaction.IdEmployee).First(&user).Error; err != nil {
		return false, errors.New("Id employee is not registered")
	}

	if transaction.Type == "" || (transaction.Type != "inbound" && transaction.Type != "outbound") {
		return false, errors.New("Transaction type is required and must be inbound or outbound")
	}

	return true, nil
}
