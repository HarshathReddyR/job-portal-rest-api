package services

import (
	"context"
	"errors"
	"job-portal-api/internal/models"
	"job-portal-api/internal/repository"
	"job-portal-api/redies"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	userRepo repository.UserRepo
	rdb      redies.RedisMethods
}

//go:generate mockgen -source=service.go -destination=service_mock.go -package=services

type ServiceMethod interface {
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	UserLogin(ctx context.Context, email, password string) (jwt.RegisteredClaims, error)

	CreateCompany(ctx context.Context, nc models.NewCompany) (models.Company, error)
	FetchCompanies() ([]models.Company, error)
	FetchCompanyById(cid uint64) (models.Company, error)

	CreateJob(ctx context.Context, nj models.NewJob, cid uint64) (models.Job, error)
	FetchJob() ([]models.Job, error)
	FetchJobById(jid uint64) (models.Job, error)
	FetchJobByCompanyId(cid uint64) ([]models.Job, error)
	ProcessingJob(ctx context.Context, rjob []models.RequestJob) ([]models.NewRequestJob, error)
	
	ForgotPassword(ctx context.Context,ru1 models.Recive1) error
}

func NewService(userRepo repository.UserRepo,rdb redies.RedisMethods ) (ServiceMethod, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		userRepo: userRepo,
		rdb: rdb,
	}, nil
}
