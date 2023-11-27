package repository

import (
	"job-portal-api/internal/models"
	"job-portal-api/internal/pkg"

	"errors"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(userData models.User) (models.User, error) {
	result := r.db.Create(&userData).Error
	//calling default create method
	if result != nil {
		log.Info().Err(result).Send()
		return models.User{}, errors.New("unable to create new user")
	}
	// Successfully created the record, return the user.
	return userData, nil
}
func (r *Repo) UserLogin(email string) (models.User, error) {
	// We attempt to find the User record where the email
	// matches the provided email.
	var user models.User
	result := r.db.Where("email = ?", email).First(&user).Error
	if result != nil {
		log.Info().Err(result).Send()
		return models.User{}, errors.New("email not found")
	}
	return user, nil
}
func (r *Repo) ForgotPassword(ru1 models.Recive1) error {
	var user models.User
	result := r.db.Where("email = ? ", ru1.Email).First(&user).Error
	if result != nil {
		log.Info().Err(result).Send()
		return errors.New("email not found")
	}
	return nil
}
func (r *Repo) CopmarePassword(ru2 models.Recive2) error {
	var user models.User
	conformhassedpassword, err := pkg.HashPassword(ru2.ConformPassword)
	if err != nil {
		return err
	}
	r.db.Where("email = ? ", ru2.Email).First(&user)
	if user.Email != conformhassedpassword {
		return nil
	}
	return errors.New("password does not match")
}
func (r *Repo) UpdateNewPassword(ru2 models.Recive2) error {
	//var user models.User
	conformhassedpassword, err := pkg.HashPassword(ru2.ConformPassword)
	if err != nil {
		return err
	}
	err = r.db.Model(&models.User{}).Where("email = ? ", ru2.Email).Update("PasswordHash", conformhassedpassword).Error
	if err != nil {
		return err
	}
	return nil
}
