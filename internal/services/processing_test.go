package services

import (
	"context"
	"job-portal-api/internal/models"
	"job-portal-api/internal/repository"
	"job-portal-api/redies"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestService_ProcessingJob(t *testing.T) {
	type args struct {
		ctx  context.Context
		rjob []models.RequestJob
	}
	tests := []struct {
		name string
		// s       *Service
		args             args
		want             []models.NewRequestJob
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
		mockRediesResponse func()()
	}{
		{
			name: "succes",
			args: args{
				ctx: context.Background(),
				rjob: []models.RequestJob{
					{Name: "a", JobRole: "b", Description: "c", NoticePeriod: 35, Budget: 2232, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 1},
					{Name: "d", JobRole: "e", Description: "f", NoticePeriod: 45, Budget: 3332, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 2},
				},
			},
			want: []models.NewRequestJob{
				{
					Name: "d",
				},
				{
					Name: "a",
				},
			},
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{Cid: 1, JobRole: "b", Description: "c", Min_Np: 30, Max_Np: 80, Budget: 850000, JobLocation: []models.Location{{Model: gorm.Model{ID: 1}}}, Technology: []models.Technology{{Model: gorm.Model{ID: 1}}}, WorKMode: []models.WorKMode{{Model: gorm.Model{ID: 1}}}, MinExp: 1, MaxExp: 2, Qualification: []models.Qualification{{Model: gorm.Model{ID: 1}}}, Shift: []models.Shift{{Model: gorm.Model{ID: 1}}}, JobType: []models.JobType{{Model: gorm.Model{ID: 1}}}}, nil
			},
		},
		// {
		// 	name: "unsucsses",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		rjob: []models.RequestJob{
		// 			{Name: "a", JobRole: "b", Description: "c", NoticePeriod: 35, Budget: 2232, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 1},
		// 			{Name: "d", JobRole: "e", Description: "f", NoticePeriod: 45, Budget: 3332, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 2},
		// 		},
		// 	},
		// 	want:    []models.NewRequestJob{},
		// 	wantErr: true,
		// 	mockRepoResponse: func() (models.Job, error) {
		// 		return models.Job{}, errors.New("job data unable to fetch")
		// 	},
		// },
		{
			name: "unsucces due to invalid noticePeriod",
			args: args{
				ctx: context.Background(),
				rjob: []models.RequestJob{
					{Name: "a", JobRole: "b", Description: "c", NoticePeriod: 2, Budget: 2232, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 1},
					{Name: "d", JobRole: "e", Description: "f", NoticePeriod: 3, Budget: 3332, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 2},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				//return models.Job{}, errors.New("test fail")
				return models.Job{Cid: 1, JobRole: "b", Description: "c", Min_Np: 30, Max_Np: 80, Budget: 850000, JobLocation: []models.Location{{Model: gorm.Model{ID: 1}}}, Technology: []models.Technology{{Model: gorm.Model{ID: 1}}}, WorKMode: []models.WorKMode{{Model: gorm.Model{ID: 1}}}, MinExp: 1, MaxExp: 2, Qualification: []models.Qualification{{Model: gorm.Model{ID: 1}}}, Shift: []models.Shift{{Model: gorm.Model{ID: 1}}}, JobType: []models.JobType{{Model: gorm.Model{ID: 1}}}}, nil
			},
		},
		{
			name: "unsucces due to invalid budget",
			args: args{
				ctx: context.Background(),
				rjob: []models.RequestJob{
					{Name: "a", JobRole: "b", Description: "c", NoticePeriod: 35, Budget: 10000000000000, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 1},
					{Name: "d", JobRole: "e", Description: "f", NoticePeriod: 45, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 2},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{Cid: 1, JobRole: "b", Description: "c", Min_Np: 30, Max_Np: 80, Budget: 850000, JobLocation: []models.Location{{Model: gorm.Model{ID: 1}}}, Technology: []models.Technology{{Model: gorm.Model{ID: 1}}}, WorKMode: []models.WorKMode{{Model: gorm.Model{ID: 1}}}, MinExp: 1, MaxExp: 2, Qualification: []models.Qualification{{Model: gorm.Model{ID: 1}}}, Shift: []models.Shift{{Model: gorm.Model{ID: 1}}}, JobType: []models.JobType{{Model: gorm.Model{ID: 1}}}}, nil
			},
		},
		{
			name: "unsucces due to invalid Joblocation",
			args: args{
				ctx: context.Background(),
				rjob: []models.RequestJob{
					{Name: "a", JobRole: "b", Description: "c", NoticePeriod: 35, Budget: 2232, JobLocation: []uint{4}, Technology: []uint{1}, WorKMode: 1, Exp: 1},
					{Name: "d", JobRole: "e", Description: "f", NoticePeriod: 45, JobLocation: []uint{4}, Technology: []uint{1}, WorKMode: 1, Exp: 2},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{Cid: 1, JobRole: "b", Description: "c", Min_Np: 30, Max_Np: 80, Budget: 850000, JobLocation: []models.Location{{Model: gorm.Model{ID: 1}}}, Technology: []models.Technology{{Model: gorm.Model{ID: 1}}}, WorKMode: []models.WorKMode{{Model: gorm.Model{ID: 1}}}, MinExp: 1, MaxExp: 2, Qualification: []models.Qualification{{Model: gorm.Model{ID: 1}}}, Shift: []models.Shift{{Model: gorm.Model{ID: 1}}}, JobType: []models.JobType{{Model: gorm.Model{ID: 1}}}}, nil
			},
		},
		{
			name: "unsucces due to invalid technology",
			args: args{
				ctx: context.Background(),
				rjob: []models.RequestJob{
					{Name: "a", JobRole: "b", Description: "c", NoticePeriod: 35, Budget: 2232, JobLocation: []uint{1}, Technology: []uint{2}, WorKMode: 1, Exp: 1},
					{Name: "d", JobRole: "e", Description: "f", NoticePeriod: 45, JobLocation: []uint{1}, Technology: []uint{5}, WorKMode: 1, Exp: 2},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{Cid: 1, JobRole: "b", Description: "c", Min_Np: 30, Max_Np: 80, Budget: 850000, JobLocation: []models.Location{{Model: gorm.Model{ID: 1}}}, Technology: []models.Technology{{Model: gorm.Model{ID: 1}}}, WorKMode: []models.WorKMode{{Model: gorm.Model{ID: 1}}}, MinExp: 1, MaxExp: 2, Qualification: []models.Qualification{{Model: gorm.Model{ID: 1}}}, Shift: []models.Shift{{Model: gorm.Model{ID: 1}}}, JobType: []models.JobType{{Model: gorm.Model{ID: 1}}}}, nil
			},
		},
		{
			name: "unsucces due to invalid experience",
			args: args{
				ctx: context.Background(),
				rjob: []models.RequestJob{
					{Name: "a", JobRole: "b", Description: "c", NoticePeriod: 35, Budget: 2232, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 8},
					{Name: "d", JobRole: "e", Description: "f", NoticePeriod: 45, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 1, Exp: 8},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{Cid: 1, JobRole: "b", Description: "c", Min_Np: 30, Max_Np: 80, Budget: 850000, JobLocation: []models.Location{{Model: gorm.Model{ID: 1}}}, Technology: []models.Technology{{Model: gorm.Model{ID: 1}}}, WorKMode: []models.WorKMode{{Model: gorm.Model{ID: 1}}}, MinExp: 1, MaxExp: 2, Qualification: []models.Qualification{{Model: gorm.Model{ID: 1}}}, Shift: []models.Shift{{Model: gorm.Model{ID: 1}}}, JobType: []models.JobType{{Model: gorm.Model{ID: 1}}}}, nil
			},
		},
		{
			name: "unsucces due to invalid work mode",
			args: args{
				ctx: context.Background(),
				rjob: []models.RequestJob{
					{Name: "a", JobRole: "b", Description: "c", NoticePeriod: 35, Budget: 2232, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 2, Exp: 1},
					{Name: "d", JobRole: "e", Description: "f", NoticePeriod: 45, JobLocation: []uint{1}, Technology: []uint{1}, WorKMode: 3, Exp: 2},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{Cid: 1, JobRole: "b", Description: "c", Min_Np: 30, Max_Np: 80, Budget: 850000, JobLocation: []models.Location{{Model: gorm.Model{ID: 1}}}, Technology: []models.Technology{{Model: gorm.Model{ID: 1}}}, WorKMode: []models.WorKMode{{Model: gorm.Model{ID: 1}}}, MinExp: 1, MaxExp: 2, Qualification: []models.Qualification{{Model: gorm.Model{ID: 1}}}, Shift: []models.Shift{{Model: gorm.Model{ID: 1}}}, JobType: []models.JobType{{Model: gorm.Model{ID: 1}}}}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			mockRedis := redies.NewMockRedisMethods(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobDetailsById(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, mockRedis)
			got, err := s.ProcessingJob(tt.args.ctx, tt.args.rjob)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ProcessingJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ProcessingJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
