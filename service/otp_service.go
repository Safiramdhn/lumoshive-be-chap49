package service

import (
	"golang-chap49/models"
	"golang-chap49/repository"

	"go.uber.org/zap"
)

type OtpService interface {
	CreateOtp(OtpInput models.OTPRequest) error
	GetUserByOTP(otp int) (*models.User, error)
}

type otpService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CreateOtp implements OtpService.
func (o *otpService) CreateOtp(OtpInput models.OTPRequest) error {
	err := o.Repo.Otp.Create(OtpInput)
	if err != nil {
		o.Log.Error("Failed to create OTP", zap.Error(err))
	}
	return err
}

// GetOTPRequest implements OtpService.
func (o *otpService) GetUserByOTP(otp int) (*models.User, error) {
	otpRequest, err := o.Repo.Otp.GetByOTP(otp)
	if err != nil {
		o.Log.Error("Failed to get OTP request", zap.Error(err))
		return nil, err
	}

	user, err := o.Repo.User.GetByEmail(otpRequest.UserEmail)
	if err != nil {
		o.Log.Error("Failed to get user by email", zap.Error(err))
		return nil, err
	}
	return user, err
}

func NewOtpService(repo repository.Repository, log *zap.Logger) OtpService {
	return &otpService{Repo: repo, Log: log}
}
