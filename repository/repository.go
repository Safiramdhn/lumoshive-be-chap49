package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	User UserRepository
	Otp  OtpRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		User: NewUserRepository(db, log),
		Otp:  NewOtpRepository(db, log),
	}
}
