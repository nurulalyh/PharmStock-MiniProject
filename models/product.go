package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	openai "github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

// Struct Product
type Products struct {
	Id                string               `gorm:"primaryKey;type:varchar(10)" json:"id" form:"id"`
	Name              string               `gorm:"type:varchar(255);not null" json:"name" form:"name"`
	Image             string               `gorm:"type:text;not null" json:"image" form:"image"`
	IdCatProduct      string               `gorm:"type:varchar(10);not null" json:"id_cat_product" form:"id_cat_product"`
	MfDate            time.Time            `gorm:"type:date;not null" json:"mf_date" form:"mf_date"`
	ExpDate           time.Time            `gorm:"type:date;not null" json:"exp_date" form:"exp_date"`
	BatchNumber       int                  `gorm:"type:smallint;not null" json:"batch_number" form:"batch_number"`
	UnitPrice         int                  `gorm:"type:smallint;not null" json:"unit_price" form:"unit_price"`
	Stock             int                  `gorm:"type:smallint;not null" json:"stock" form:"stock"`
	Description       string               `gorm:"type:text;not null" json:"description" form:"description"`
	IdDistributor     string               `gorm:"type:varchar(10);not null" json:"id_distributor" form:"id_distributor"`
	CreatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt         gorm.DeletedAt       `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DetailTransaction []DetailTransactions `gorm:"foreignKey:id_product;references:Id" json:"transaction" form:"transaction"`
}

// Interface beetween models and controller
type ProductsModelInterface interface {
	Insert(newProduct Products) (*Products, error)
	SelectAll(limit, offset int) ([]Products, error)
	Update(updatedData Products) (*Products, error)
	SearchProduct(keyword string, limit int, offset int) ([]Products, error)
	AIGenerateDescription(userInput string, openAIKey string) (string, error)
}

// connect into db
type ProductsModel struct {
	db *gorm.DB
}

// New Instance from ProductsModel
func NewProductsModel(db *gorm.DB) ProductsModelInterface {
	return &ProductsModel{
		db: db,
	}
}

// Insert Product
func (pm *ProductsModel) Insert(newProduct Products) (*Products, error) {
	var latestProduct Products
	if errSort := pm.db.Unscoped().Order("id DESC").First(&latestProduct).Error; errSort != nil {
		latestProduct.Id = "PDT-0000"
	}

	newID := generateProductId(latestProduct.Id)
	if newID == "" {
		return nil, errors.New("Failed generate Id")
	}

	newProduct.Id = newID

	validate, errValidate := validateProduct(newProduct, pm.db)
	if !validate {
		return nil, errValidate
	}

	if err := pm.db.Create(&newProduct).Error; err != nil {
		return nil, errors.New("Error insert product, " + err.Error())
	}

	return &newProduct, nil
}

// Select All Product
func (pm *ProductsModel) SelectAll(limit, offset int) ([]Products, error) {
	var data []Products

	if err := pm.db.
		Limit(limit).
		Offset(offset).
		Preload("DetailTransactions").
		Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get product with detail_transactions, " + err.Error())
	}

	return data, nil
}

// Update Product
func (pm *ProductsModel) Update(updatedData Products) (*Products, error) {
	var data map[string]interface{} = make(map[string]interface{})

	if updatedData.Name != "" {
		data["name"] = updatedData.Name
	}
	if updatedData.Image != "" {
		data["image"] = updatedData.Image
	}
	if updatedData.IdCatProduct != "" {
		var catProduct CatProducts
		if err := pm.db.Where("id = ?", updatedData.IdCatProduct).First(&catProduct).Error; err != nil {
			return nil, errors.New("Id category product is not registered")
		}
		data["id_cat_product"] = updatedData.IdCatProduct
	}
	if !updatedData.MfDate.IsZero() {
		data["mf_date"] = updatedData.MfDate
	}
	if !updatedData.ExpDate.IsZero() {
		data["exp_date"] = updatedData.ExpDate
	}
	if updatedData.BatchNumber != 0 {
		data["batch_number"] = updatedData.BatchNumber
	}
	if updatedData.UnitPrice != 0 {
		data["unit_price"] = updatedData.UnitPrice
	}
	if updatedData.Stock != 0 {
		data["stock"] = updatedData.Stock
	}
	if updatedData.Description != "" {
		data["description"] = updatedData.Description
	}
	if updatedData.IdDistributor != "" {
		var distributor Distributors
		if err := pm.db.Where("id = ?", updatedData.IdDistributor).First(&distributor).Error; err != nil {
			return nil, errors.New("Id distributor is not registered")
		}
		data["id_distributor"] = updatedData.IdDistributor
	}

	var qry = pm.db.Table("products").Where("id = ?", updatedData.Id).Updates(data)
	if err := qry.Error; err != nil {
		return nil, errors.New("Error update data" + err.Error())
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return nil, errors.New("Error update data, no data affected")
	}

	var updatedProduct = Products{}
	if err := pm.db.Where("id = ?", updatedData.Id).First(&updatedProduct).Error; err != nil {
		return nil, errors.New("Cannot get updated data" + err.Error())
	}

	return &updatedProduct, nil
}

// Searching
func (pm *ProductsModel) SearchProduct(keyword string, limit int, offset int) ([]Products, error) {
	var product []Products
	query := pm.db.Limit(limit).Offset(offset).Where("id LIKE ? OR name LIKE ? OR photo LIKE ? OR id_cat_product LIKE ? OR mf_Date LIKE ? OR exp_date LIKE ? OR batch_number LIKE ? OR unit_price LIKE ? OR stock LIKE ? OR description LIKE ? OR id_distributor LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if err := query.
		Preload("DetailTransactions").
		Find(&product).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return product, nil
}

// AI Generate Description
func (pm *ProductsModel) AIGenerateDescription(userInput string, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Halo, perkenalkan saya sistem untuk generate deskripsi dari produk farmasi",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		},
	}

	resp, err := pm.getCompletionFromMessages(ctx, client, message, model)
	if err != nil {
		return "", err
	}

	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func (pm *ProductsModel) getCompletionFromMessages(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage, model string) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

// Generate Id
func generateProductId(latestID string) string {
	var numID int
	if _, err := fmt.Sscanf(latestID, "PDT-%04d", &numID); err != nil {
		return ""
	}
	numID++
	return fmt.Sprintf("PDT-%04d", numID)
}

// Validate
func validateProduct(product Products, db *gorm.DB) (bool, error) {
	if product.Id == "" || len(product.Id) > 10 {
		return false, errors.New("Id is required and must be up to 10 characters")
	}

	if product.Name == "" || len(product.Name) > 255 {
		return false, errors.New("Product name is required and must be up to 255 characters")
	}

	if product.Image == "" {
		return false, errors.New("Image is required")
	}

	if product.IdCatProduct == "" || len(product.IdCatProduct) > 10 {
		return false, errors.New("Id category product is required and must be up to 10 characters")
	}

	var catProduct CatProducts
	if err := db.Where("id = ?", product.IdCatProduct).First(&catProduct).Error; err != nil {
		return false, errors.New("Id category product is not registered")
	}

	if product.MfDate.IsZero() {
		return false, errors.New("Manufactoring date is required")
	}

	if product.MfDate.IsZero() {
		return false, errors.New("Expire date is required")
	}

	if product.BatchNumber == 0 {
		return false, errors.New("Batch number is required")
	}

	if product.UnitPrice == 0 {
		return false, errors.New("Unit Price is required")
	}

	if product.Stock == 0 {
		return false, errors.New("Stock is required")
	}

	if product.Description == "" {
		return false, errors.New("Decription product is required")
	}

	if product.IdDistributor == "" || len(product.IdDistributor) > 10 {
		return false, errors.New("Id distributor is required and must be up to 10 characters")
	}

	var distributor Distributors
	if err := db.Where("id = ?", product.IdDistributor).First(&distributor).Error; err != nil {
		return false, errors.New("Id distributor is not registered")
	}

	return true, nil
}
