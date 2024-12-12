package models

import "time"

type OTPRequest struct {
	OTP         int       `json:"otp" gorm:"type:INTEGER;not null"`
	UserEmail   string    `json:"user_email" gorm:"type:VARCHAR(255);not null"`
	ExpiredTime time.Time `json:"expired_time" gorm:"type:TIMESTAMP;not null"`
}
