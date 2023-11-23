package services

import (
	"context"
	"errors"
	"job-portal-api/internal/models"
	"job-portal-api/internal/repository"
	"job-portal-api/redies"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_FetchJobById(t *testing.T) {
	type args struct {
		jid uint64
	}
	tests := []struct {
		name             string
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		{
			name: "success",
			want: models.Job{
				Company: models.Company{
					CompanyName: "ll",
					Location:    "kkk",
				},
				Cid:         1,
				JobRole:     "xyz",
				Description: "abcd",
			},
			args: args{
				jid: 15,
			},
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{
					Company: models.Company{
						CompanyName: "ll",
						Location:    "kkk",
					},
					Cid:         1,
					JobRole:     "xyz",
					Description: "abcd",
				}, nil
			},
		},
		{
			name: "invalid job id",
			want: models.Job{},
			args: args{
				jid: 5,
			},
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("error test")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRedis:=redies.NewMockRedisMethods(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobDetailsById(tt.args.jid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo,mockRedis)
			got, err := s.FetchJobById(tt.args.jid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.FetchJobById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.FetchJobById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_FetchJob(t *testing.T) {
	tests := []struct {
		name             string
		want             []models.Job
		wantErr          bool
		mockRepoResponse func() ([]models.Job, error)
	}{
		{
			name: "database success",
			want: []models.Job{
				{
					Cid:         1,
					JobRole:     "aa",
					Description: "pp",
				},
				{
					Cid:         2,
					JobRole:     "bb",
					Description: "qq",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{
					{
						Cid:         1,
						JobRole:     "aa",
						Description: "pp",
					},
					{
						Cid:         2,
						JobRole:     "bb",
						Description: "qq",
					},
				}, nil
			},
		},
		{
			name:    "database failure",
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Job, error) {
				return nil, errors.New("error test")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRedis:=redies.NewMockRedisMethods(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewAllJobs().Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo,mockRedis)
			got, err := s.FetchJob()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.FetchJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.FetchJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_FetchJobByCompanyId(t *testing.T) {
	type args struct {
		cid uint64
	}
	tests := []struct {
		name             string
		args             args
		want             []models.Job
		wantErr          bool
		mockRepoResponse func() ([]models.Job, error)
	}{
		{
			name: "success",
			args: args{
				cid: 2,
			},
			want: []models.Job{
				{
					Cid:         1,
					JobRole:     "aa",
					Description: "pp",
				},
				{
					Cid:         2,
					JobRole:     "bb",
					Description: "qq",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{
					{
						Cid:         1,
						JobRole:     "aa",
						Description: "pp",
					},
					{
						Cid:         2,
						JobRole:     "bb",
						Description: "qq",
					},
				}, nil
			},
		},
		{
			name: "failure",
			args: args{
				cid: 10,
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Job, error) {
				return nil, errors.New("data is not there")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRedis:=redies.NewMockRedisMethods(mc)
			
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobByCompanyId(tt.args.cid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo,mockRedis)
			got, err := s.FetchJobByCompanyId(tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.FetchJobByCompanyId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.FetchJobByCompanyId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateJob(t *testing.T) {
	type args struct {
		ctx     context.Context
		jobData models.NewJob
		cid     uint64
	}
	tests := []struct {
		name             string
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				jobData: models.NewJob{
					JobRole:       "a",
					Description:   "mm",
					Min_Np:        2,
					Max_Np:        6,
					Budget:        22000,
					JobLocation:   []uint{1},
					Technology:    []uint{1, 2},
					WorKMode:      []uint{1, 2},
					MinExp:        2,
					MaxExp:        5,
					Qualification: []uint{1},
					Shift:         []uint{1, 2},
					JobType:       []uint{1},
				},
				cid: 1,
			},
			want: models.Job{
				JobRole:       "a",
				Description:   "mm",
				Min_Np:        2,
				Max_Np:        6,
				Budget:        22000,
				JobLocation:   []models.Location{},
				Technology:    []models.Technology{},
				WorKMode:      []models.WorKMode{},
				MinExp:        2,
				MaxExp:        5,
				Qualification: []models.Qualification{},
				Shift:         []models.Shift{},
				JobType:       []models.JobType{},
			},
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{
					JobRole:       "a",
				Description:   "mm",
				Min_Np:        2,
				Max_Np:        6,
				Budget:        22000,
				JobLocation:   []models.Location{},
				Technology:    []models.Technology{},
				WorKMode:      []models.WorKMode{},
				MinExp:        2,
				MaxExp:        5,
				Qualification: []models.Qualification{},
				Shift:         []models.Shift{},
				JobType:       []models.JobType{},
				}, nil
			},
		},
		{
			name: "failure",
			args: args{
				ctx: context.Background(),
				jobData: models.NewJob{
					JobRole:     "a",
					Description: "mm",
				},
				cid: 1,
			},
			want:    models.Job{},
			wantErr: true,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("job is not created")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRedis:=redies.NewMockRedisMethods(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateJob(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo,mockRedis)
			got, err := s.CreateJob(tt.args.ctx, tt.args.jobData, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.CreateJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
