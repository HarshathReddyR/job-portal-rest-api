package models

import "gorm.io/gorm"

type NewUser struct {
	Name     string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	DOB      string `json:"dob" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	gorm.Model
	Name         string `json:"username"`
	Email        string `json:"email"`
	DOB          string `json:"dob"`
	PasswordHash string `json:"-"`
}
type Recive1 struct {
	Email string `json:"email"`
	DOB   string `json:"dob"`
}

// type User1 struct {
// 	gorm.Model
// 	Name         string `json:"username"`
// 	Email        string `json:"email"`
// 	DOB          string `json:"dob"`
// 	PasswordHash string `json:"-"`
// }
type Recive2 struct {
	OTP             string `json:"otp"`
	NewPassword     string `json:"newpassword"`
	ConformPassword string `json:"Conformpassword"`
}
