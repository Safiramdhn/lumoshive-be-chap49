package repository

import (
	"golang-chap49/models"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(userInput models.User) error
	GetByEmail(email string) (*models.User, error)
	UpdatePassword(resetPasswordInput models.LoginRequest) error
	CountCustomer() (int, error)
}

type userRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Create implements UserRepository.
func (u *userRepository) Create(userInput models.User) error {
	return u.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userInput).Error; err != nil {
			u.Log.Error("Failed to create user", zap.Error(err))
			return err
		}
		return nil
	})
}

// GetByEmail implements UserRepository.
func (u *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// UpdatePassword implements UserRepository.
func (u *userRepository) UpdatePassword(resetPasswordInput models.LoginRequest) error {
	return u.DB.Model(&models.User{}).Where("email =?", resetPasswordInput.Email).Update("password", resetPasswordInput.Password).Error
}

// CountCustomer implements UserRepository.
func (u *userRepository) CountCustomer() (int, error) {
	var count int64
	err := u.DB.Model(&models.User{}).
		Where("role = ?", "customer").
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", time.Now().Month(), time.Now().Year()).
		Count(&count).Error

	return int(count), err
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{DB: db, Log: log}
}
