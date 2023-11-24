package services

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"job-portal-api/internal/models"
	"job-portal-api/internal/pkg"
	"math/big"
	"net/smtp"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) CreateUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	//method that creates a new record in  db
	hashedPass, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return models.User{}, err
	}
	//prepare user record
	userDetails := models.User{
		Name:         userData.Name,
		Email:        userData.Email,
		DOB:          userData.DOB,
		PasswordHash: string(hashedPass),
	}
	userDetails, err = s.userRepo.CreateUser(userDetails)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, nil
}
func (s *Service) UserLogin(ctx context.Context, email, password string) (jwt.RegisteredClaims, error) {
	//checking the email in database
	userDetails, err := s.userRepo.UserLogin(email)
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	err = pkg.CheckPassword(password, userDetails.PasswordHash)
	if err != nil {
		log.Info().Err(err).Send()
		return jwt.RegisteredClaims{}, errors.New("entered password is wrong")
	}
	claims := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	return claims, nil

}

func (s *Service) ForgotPassword(ctx context.Context, ru1 models.Recive1) error {
	err := s.userRepo.ForgotPassword(ru1)
	if err != nil {
		return errors.New("invalid email")
	}
	otp, err := GenerateOTP()
	if err != nil {
		return err
	}
	// Message content
	a := ru1.Email

	err = s.rdb.AddOTPToRedis(ctx, a, otp)
	if err != nil {
		return err
	}
	err = SendMail(ctx, ru1, otp)
	if err != nil {
		return err
	}
	return nil
}
func SendMail(ctx context.Context, ru1 models.Recive1, otp string) error {
	// Sender's email address and password
	from := "harshathreddy18@gmail.com"
	password := "xhod ymbp xarp nehf"

	// Recipient's email address
	to := ru1.Email

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	// otp,err:=GenerateOTP()
	// if err!=nil{
	// 	return err
	// }
	// // Message content

	// err=redies.RedisMethods.AddOTPToRedis(ctx,ru1.Email,otp)

	message := []byte("One Time Password" + otp)

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err := smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if err != nil {
		// fmt.Println("Error sending email:", err)
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
func GenerateOTP() (string, error) {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(999999))
	if err != nil {
		return "", err
	}

	// Add leading zeros if necessary
	otp := fmt.Sprintf("%06d", randomNumber)

	return otp, nil
}
