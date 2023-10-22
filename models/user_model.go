package models

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	Id           int            `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Name         string         `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	Username     string         `gorm:"type:varchar(25);not null" json:"username" form:"username"`
	Password     string         `gorm:"type:varchar(25);not null" json:"password" form:"password"`
	Email        string         `gorm:"type:varchar(50);not null" json:"email" form:"email"`
	Phone        string         `gorm:"type:varchar(15);not null" json:"phone" form:"phone"`
	Address      string         `gorm:"type:varchar(255);not null" json:"address" form:"address"`
	Role         string         `gorm:"type:varchar(25);not null" json:"role" form:"role"`
	CreatedAt    time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt    time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Transactions []Transaction  `gorm:"foreignKey:IdUser;references:Id"`
	ReqProducts  []ReqProduct   `gorm:"foreignKey:IdUser;references:Id"`
}

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// Declaration Contract
type UserModelInterface interface {
	Login(username string, password string) *User
	Insert(newItem User) *User
	SelectAll() []User
	SelectById(userId int) *User
}

// Interaction with db
type UsersModel struct {
	db *gorm.DB
}

// New Instance from UsersModel
func NewUsersModel(db *gorm.DB) UserModelInterface {
	return &UsersModel{
		db: db,
	}
}

func (um *UsersModel) Login(username string, password string) *User {
	var data = User{}
	if err := um.db.Where("username = ?", username).First(&data).Error; err != nil {
		logrus.Error("Model : Login data error, ", err.Error())
		return nil
	}
	if data.Id == 0 {
		logrus.Error("Model : Login data error, ", nil)
		return nil
	}

	return &data
}

func (um *UsersModel) Insert(newUser User) *User {
	if err := um.db.Create(&newUser).Error; err != nil {
		logrus.Error("Model : Insert data error, ", err.Error())
		return nil
	}

	return &newUser
}

func (um *UsersModel) SelectAll() []User {
	var data = []User{}
	if err := um.db.Find(&data).Error; err != nil {
		logrus.Error("Model : Error get all users, ", err.Error())
		return nil
	}

	return data
}

func (um *UsersModel) SelectById(userId int) *User {
	var data = User{}
	if err := um.db.Where("id = ?", userId).First(&data).Error; err != nil {
		logrus.Error("Model : Error get user by id, ", err.Error())
		return nil
	}

	return &data
}