package services

import (
	"context"
	"job-portal-api/internal/models"

	"gorm.io/gorm"
)

func (s *Service) CreateJob(ctx context.Context, jobData models.NewJob, cid uint64) (models.Job, error) {
	jobDetails := models.Job{
		JobRole:     jobData.JobRole,
		Description: jobData.Description,
		Cid:         uint(cid),
		Min_Np:      jobData.Min_Np,
		Max_Np:      jobData.Max_Np,
		Budget:      jobData.Budget,
		MinExp:      jobData.MinExp,
		MaxExp:      jobData.MaxExp,
	}
	for _, v := range jobData.JobLocation {
		templocation := models.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobDetails.JobLocation = append(jobDetails.JobLocation, templocation)
	}

	for _, v := range jobData.Technology {
		temptechnology := models.Technology{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobDetails.Technology = append(jobDetails.Technology, temptechnology)
	}

	for _, v := range jobData.WorKMode {
		tempworkmode := models.WorKMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobDetails.WorKMode = append(jobDetails.WorKMode, tempworkmode)
	}
	for _, v := range jobData.Qualification {
		tempqualification := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobDetails.WorKMode = append(jobDetails.WorKMode, models.WorKMode(tempqualification))
	}
	for _, v := range jobData.Shift {
		tempshift := models.WorKMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobDetails.WorKMode = append(jobDetails.WorKMode, tempshift)
	}
	for _, v := range jobData.JobType {
		tempjobtype := models.JobType{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobDetails.WorKMode = append(jobDetails.WorKMode, models.WorKMode(tempjobtype))
	}
	jobDetails, err := s.userRepo.CreateJob(jobDetails)
	if err != nil {
		return models.Job{}, err
	}
	return jobDetails, nil
}
func (s *Service) FetchJob() ([]models.Job, error) {
	jobDetails, err := s.userRepo.ViewAllJobs()
	if err != nil {
		return nil, err
	}
	return jobDetails, nil
}
func (s *Service) FetchJobById(jid uint64) (models.Job, error) {
	jobDetails, err := s.userRepo.ViewJobDetailsById(jid)
	if err != nil {
		return models.Job{}, err
	}
	return jobDetails, nil
}
func (s *Service) FetchJobByCompanyId(cid uint64) ([]models.Job, error) {
	jobDetails, err := s.userRepo.ViewJobByCompanyId(cid)
	if err != nil {
		return nil, err
	}
	return jobDetails, nil
}
