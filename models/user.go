package models

import (
	"golang-chap49/helper"
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement" swaggerignore:"true"`
	Name      string    `json:"name" gorm:"size:100;not null" binding:"required" example:"John Doe"`
	Email     string    `json:"email" gorm:"size:100;unique" binding:"required,email" example:"johndoe@example.com"`
	Password  string    `json:"-" gorm:"size:255;not null" binding:"required" swaggerignore:"true"`
	Role      string    `json:"role" gorm:"size:20;check:role IN ('admin','customer', 'staff')" binding:"required,oneof=admin customer staff" example:"admin"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"johndoe@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"password123"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"johndoe@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"password123"`
}

type CheckEmailRequest struct {
	Email string `json:"email" binding:"required,email" example:"johndoe@example"`
}

func UserSeed() []User {
	return []User{
		{
			Name:     "Admin User",
			Email:    "admin@example.com",
			Password: helper.HashPassword("adminpassword"),
			Role:     "admin",
		},
		{
			Name:     "Staff User 1",
			Email:    "staff1@example.com",
			Password: helper.HashPassword("staffpassword1"),
			Role:     "staff",
		},
		{
			Name:     "Staff User 2",
			Email:    "staff2@example.com",
			Password: helper.HashPassword("staffpassword2"),
			Role:     "staff",
		},
		{
			Name:     "User 1",
			Email:    "user1@example.com",
			Password: helper.HashPassword("userpassword1"),
			Role:     "customer",
		},
		{
			Name:     "User 2",
			Email:    "user2@example.com",
			Password: helper.HashPassword("userpassword2"),
			Role:     "customer",
		},
		{
			Name:     "User 3",
			Email:    "user3@example.com",
			Password: helper.HashPassword("userpassword3"),
			Role:     "customer",
		},
		{
			Name:     "User 4",
			Email:    "user4@example.com",
			Password: helper.HashPassword("userpassword4"),
			Role:     "customer",
		},
		{
			Name:     "User 5",
			Email:    "user5@example.com",
			Password: helper.HashPassword("userpassword5"),
			Role:     "customer",
		},
		{
			Name:     "User 6",
			Email:    "user6@example.com",
			Password: helper.HashPassword("userpassword6"),
			Role:     "customer",
		},
		{
			Name:     "User 7",
			Email:    "user7@example.com",
			Password: helper.HashPassword("userpassword7"),
			Role:     "customer",
		},
		{
			Name:     "User 8",
			Email:    "user8@example.com",
			Password: helper.HashPassword("userpassword8"),
			Role:     "customer",
		},
		{
			Name:     "User 9",
			Email:    "user9@example.com",
			Password: helper.HashPassword("userpassword9"),
			Role:     "customer",
		},
		{
			Name:     "User 10",
			Email:    "user10@example.com",
			Password: helper.HashPassword("userpassword10"),
			Role:     "customer",
		},
	}
}
