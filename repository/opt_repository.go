package repository

import (
	"golang-chap49/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OtpRepository interface {
	Create(OtpInput models.OTPRequest) error
	GetByOTP(otp int) (models.OTPRequest, error)
}

type otpRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// Create implements OtpRepository.
func (o *otpRepository) Create(OtpInput models.OTPRequest) error {
	return o.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&OtpInput).Error
	})
}

// GetByOTP implements OtpRepository.
func (o *otpRepository) GetByOTP(otp int) (models.OTPRequest, error) {
	var otpData models.OTPRequest
	err := o.DB.Where("otp =?", otp).First(&otpData).Error
	return otpData, err
}

func NewOtpRepository(db *gorm.DB, log *zap.Logger) OtpRepository {
	return &otpRepository{DB: db, Log: log}
}
