package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(50);unique;not null"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(100);unique;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(15);unique;not null"`
	FullName    string    `json:"full_name" gorm:"column:full_name;type:varchar(100);not null"`
	Address     string    `json:"address" gorm:"column:address;type:text;not null"`
	DOB         string    `json:"dob" gorm:"column:dob;type:varchar(50);not null"`
	Password    string    `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Role        string    `json:"role" gorm:"column:role;type:enum('user','seller');not null"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (*User) TableName() string {
	return "users"
}

func (l *User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  int `gorm:"primary_key"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              int       `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string    `json:"token" gorm:"type:text" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:text" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l *UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
