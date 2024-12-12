package service

import (
	"golang-chap49/repository"

	"go.uber.org/zap"
)

type Service struct {
	User UserService
	Otp  OtpService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		User: NewUserService(repo, log),
		Otp:  NewOtpService(repo, log),
	}
}
