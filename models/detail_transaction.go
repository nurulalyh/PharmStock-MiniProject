package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Struct Detail Transaction
type DetailTransactions struct {
	Id            string         `gorm:"primaryKey;type:varchar(10)" json:"id" form:"id"`
	IdTransaction string         `gorm:"type:varchar(10);not null" json:"id_transaction" form:"id_transaction"`
	IdProduct     string         `gorm:"type:varchar(10);not null" json:"id_product" form:"id_product"`
	Quantity      int            `gorm:"type:smallint;not null" json:"quantity" form:"quantity"`
	Price         int            `gorm:"type:int;not null" json:"price" form:"price"`
	CreatedAt     time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
}

// Interface beetween models and controller
type DetailTransactionsModelInterface interface {
	Insert(newDetailTransaction DetailTransactions) (*DetailTransactions, error)
	SelectAll(limit, offset int) ([]DetailTransactions, error)
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

	validate, errValidate := validateDetailTransaction(newDetailTransaction, dtm.db)
	if !validate {
		return nil, errValidate
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

// Searching
func (dtm *DetailTransactionsModel) SearchDetailTransaction(keyword string, limit int, offset int) ([]DetailTransactions, error) {
	var detailTransaction []DetailTransactions
	query := dtm.db.Limit(limit).Offset(offset).Where("id LIKE ? OR id_transaction LIKE ? OR id_product LIKE ? OR quantity LIKE ? OR price LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.Find(&detailTransaction).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return detailTransaction, nil
}

// Hook Before Create to count price
func (dt *DetailTransactions) BeforeCreate(tx *gorm.DB) error {
    var product Products
    if err := tx.Model(&Products{}).Where("id = ?", dt.IdProduct).First(&product).Error; err != nil {
        return err
    }

    dt.Price = product.UnitPrice * dt.Quantity

    return nil
}

// Hook After Create to manage product stock and detail_transaction quantity, to acumulate price and quantity from detail_transaction
func (dt *DetailTransactions) AfterCreate(tx *gorm.DB) (err error) {
	if err := tx.Model(&Products{}).Where("id = ?", dt.IdProduct).UpdateColumn("stock", gorm.Expr("stock - ?", dt.Quantity)).Error; err != nil {
		return err
	}

	var totalQuantity int
    var totalPrice int

    if errQuery := tx.Model(&DetailTransactions{}).Where("id_transaction = ?", dt.IdTransaction).Select("SUM(quantity) as total_quantity, SUM(price) as total_price").Row().Scan(&totalQuantity, &totalPrice); errQuery != nil {
        return errQuery
    }

    if errUpdate := tx.Model(&Transactions{}).Where("id = ?", dt.IdTransaction).UpdateColumns(map[string]interface{}{
        "TotalQuantity": totalQuantity,
        "TotalPrice":    totalPrice,
    }).Error; errUpdate != nil {
        return errUpdate
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
func validateDetailTransaction(detailTransaction DetailTransactions, db *gorm.DB) (bool, error) {
	if detailTransaction.Id == "" || len(detailTransaction.Id) > 10 {
		return false, errors.New("Id is required and must be up to 10 characters")
	}

	if detailTransaction.IdTransaction == "" || len(detailTransaction.IdTransaction) > 10 {
		return false, errors.New("Id Trasantion is required and must be up to 10 characters")
	}

	var transaction Transactions
	if err := db.Where("id = ?", detailTransaction.IdTransaction).First(&transaction).Error; err != nil {
		return false, errors.New("Id transaction is not registered")
	}

	if detailTransaction.IdProduct == "" || len(detailTransaction.IdProduct) > 10 {
		return false, errors.New("Id product is required and must be up to 10 characters")
	}

	var product Products
	if err := db.Where("id = ?", detailTransaction.IdProduct).First(&product).Error; err != nil {
		return false, errors.New("Id product is not registered")
	}

	if detailTransaction.Quantity == 0 {
		return false, errors.New("Quantity is required")
	}

	return true, nil
}
