package models

import (
	"errors"
	"fmt"
	"pharm-stock/helper/authentication"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Struct Users
type Users struct {
	Id           string         `gorm:"primaryKey;type:varchar(10)"`
	Name         string         `gorm:"type:varchar(100);not null"`
	Username     string         `gorm:"type:varchar(50);not null" valid:"required~your user name is required"`
	Password     string         `gorm:"type:text;not null" valid:"required~your password is required,minstringlength(8)~Password has to have a minimum length of 6 characters"`
	Email        string         `gorm:"type:varchar(50);not null" valid:"required~your email is required, email~invalid email format"`
	Phone        string         `gorm:"type:varchar(15);not null" valid:"required~your phone is required"`
	Address      string         `gorm:"type:varchar(255);not null"`
	Role         string         `gorm:"type:ENUM('administrator','apoteker');not null"`
	CreatedAt    time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Transactions []Transactions `gorm:"foreignKey:id_employee;references:id"`
	ReqProducts  []ReqProducts  `gorm:"foreignKey:id_employee;references:id"`
}

// Interface beetween models and controller
type UsersModelInterface interface {
	Login(username string, password string) (*Users, error)
	Insert(newItem Users) (*Users, error)
	SelectAll(limit, offset int) ([]Users, error)
	Update(updatedData Users) (*Users, error)
	Delete(userId string) (bool, error)
	SearchUsers(keyword string, limit int, offset int) ([]Users, error)
	InsertAdmin(newItem Users) (*Users, error)
}

// connect into db
type UsersModel struct {
	db *gorm.DB
}

// New Instance from UsersModel
func NewUsersModel(db *gorm.DB) UsersModelInterface {
	return &UsersModel{
		db: db,
	}
}

// Login
func (um *UsersModel) Login(username string, password string) (*Users, error) {
	var data = Users{}
	if err := um.db.Where("username = ?", username).First(&data).Error; err != nil {
		return nil, errors.New("Error user login, " + err.Error())
	}

	if data.Id == "" {
		return nil, errors.New("User not found")
	}

	return &data, nil
}

// Insert User
func (um *UsersModel) Insert(newUser Users) (*Users, error) {
	var latestUser Users
	if errSort := um.db.Order("id DESC").First(&latestUser).Error; errSort != nil {
		return nil, errors.New("Error filter data, " + errSort.Error())
	}

	newID := generateUserId(latestUser.Id)
	if newID == "" {
		return nil, errors.New("Failed generate Id")
	}

	newUser.Id = newID
	newUser.Role = "apoteker"

	validate := validateUser(newUser)
	if !validate {
		return nil, errors.New("Data not valid")
	}

	if checkUsername := um.db.Where("username = ?", newUser.Username).First(&newUser).Error; checkUsername != nil {
		if checkUsername == gorm.ErrRecordNotFound {
			if err := um.db.Create(&newUser).Error; err != nil {
				return nil, errors.New("Error insert user, " + err.Error())
			}
		} else {
			return nil, errors.New("Error checking username availability")
		}
	} else {
		return nil, errors.New("Username already exists")
	}

	return &newUser, nil
}

// Select All User
func (um *UsersModel) SelectAll(limit, offset int) ([]Users, error) {
	var data []Users
	if err := um.db.Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, errors.New("Cannot get all users, " + err.Error())
	}
	return data, nil
}

// Update User
func (um *UsersModel) Update(updatedData Users) (*Users, error) {
	var data map[string]interface{} = make(map[string]interface{})

	if updatedData.Name != "" {
		data["name"] = updatedData.Name
	}
	if updatedData.Username != "" {
		data["username"] = updatedData.Username
	}
	if updatedData.Password != "" {
		hashedPassword, err := authentication.HashPassword(updatedData.Password)
		if err != nil {
			return nil, errors.New("Cannot process data, something happened. " + err.Error())
		}
		data["password"] = hashedPassword
	}
	if updatedData.Email != "" {
		data["email"] = updatedData.Email
	}
	if updatedData.Phone != "" {
		data["phone"] = updatedData.Phone
	}
	if updatedData.Address != "" {
		data["address"] = updatedData.Address
	}
	if updatedData.Role != "" {
		data["role"] = updatedData.Role
	}

	var qry = um.db.Table("users").Where("id = ?", updatedData.Id).Updates(data)
	if err := qry.Error; err != nil {
		return nil, errors.New("Error update data" + err.Error())
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return nil, errors.New("Error update data, no data affected")
	}

	var updatedUser = Users{}
	if err := um.db.Where("id = ?", updatedData.Id).First(&updatedUser).Error; err != nil {
		return nil, errors.New("Cannot get updated data" + err.Error())
	}

	return &updatedUser, nil
}

// Delete User
func (um *UsersModel) Delete(userId string) (bool, error) {
	var data = Users{}
	data.Id = userId

	if err := um.db.Where("id = ?", userId).First(&data).Error; err != nil {
		return false, errors.New("Error finding data to delete, " + err.Error())
	}

	if err := um.db.Delete(&data).Error; err != nil {
		return false, errors.New("Error delete data, " + err.Error())
	}

	return true, nil
}

// Searching
func (um *UsersModel) SearchUsers(keyword string, limit int, offset int) ([]Users, error) {
	var users []Users
	query := um.db.Where("id LIKE ? OR name LIKE ? OR username LIKE ? OR password LIKE ? OR email LIKE ? OR phone LIKE ? OR address LIKE ? OR role LIKE ? OR created_at LIKE ? OR updated_at LIKE ? OR deleted_at LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Limit(limit).Offset(offset)

	if err := query.Find(&users).Error; err != nil {
		return nil, errors.New("Error search data, " + err.Error())
	}

	return users, nil
}

// Insert Admin
func (um *UsersModel) InsertAdmin(newUser Users) (*Users, error) {
	newUser.Id = "USR-0001"
	newUser.Role = "administrator"

	validate := validateUser(newUser)
	if !validate {
		return nil, errors.New("The data must not be empty and should adhere to the specified format")
	}

	if checkUsername := um.db.Where("username = ?", newUser.Username).First(&newUser).Error; checkUsername != nil {
		if checkUsername == gorm.ErrRecordNotFound {
			if err := um.db.Create(&newUser).Error; err != nil {
				return nil, errors.New("Error insert user, " + err.Error())
			}
		} else {
			return nil, errors.New("Error checking username availability")
		}
	} else {
		return nil, errors.New("Username already exists")
	}

	return &newUser, nil
}

// Generate Id
func generateUserId(latestID string) string {
	var numID int
	if _, err := fmt.Sscanf(latestID, "USR-%04d", &numID); err != nil {
		return ""
	}
	numID++
	return fmt.Sprintf("USR-%04d", numID)
}

// Validate
func validateUser(user Users) bool {
	if user.Id == "" || len(user.Id) > 10 {
		logrus.Error("Model: Id is required and must be up to 10 characters")
		return false
	}
	if user.Name == "" || len(user.Name) > 100 {
		logrus.Error("Model: Name is required and must be up to 100 characters")
		return false
	}
	if user.Username == "" || len(user.Username) > 50 {
		logrus.Error("Model: Username is required and must be up to 50 characters")
		return false
	}
	if user.Password == "" {
		logrus.Error("Model: Password is required and must be up to 25 characters")
		return false
	}
	if user.Email == "" || len(user.Email) > 50 {
		logrus.Error("Model: Email is required and must be up to 50 characters")
		return false
	}
	if user.Phone == "" || len(user.Phone) > 15 {
		logrus.Error("Model: Phone is required and must be up to 15 characters")
		return false
	}
	if user.Address == "" || len(user.Address) > 255 {
		logrus.Error("Model: Address is required and must be up to 255 characters")
		return false
	}
	if user.Role == "" || (user.Role != "administrator" && user.Role != "apoteker") {
		logrus.Error("Model: Role is required and must be 'administrator' or 'apoteker'")
		return false
	}

	return true
}
