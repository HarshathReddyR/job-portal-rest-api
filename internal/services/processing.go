package services

import (
	"context"
	"encoding/json"
	"fmt"
	"job-portal-api/internal/models"
	"slices"
	"sync"

	"github.com/redis/go-redis/v9"
)

func (s *Service) ProcessingJob(ctx context.Context, rjob []models.RequestJob) ([]models.NewRequestJob, error) {

	wg := new(sync.WaitGroup)

	c1 := make(chan models.NewRequestJob)
	var JobsResult []models.NewRequestJob
	for _, v := range rjob {
		wg.Add(1)
		go func(v models.RequestJob) {
			defer wg.Done()

			// jobData, err := s.userRepo.ViewJobDetailsById(uint64(v.JobId))
			// if err != nil {
			// 	log.Error().Err(err).Interface("applicant", v.Name).Msg("job data unable to fetch")
			// 	return
			// }
			var jobData models.Job
			val, err := s.rdb.GetDataFromRedis(ctx, v.JobId)
			fmt.Println(err)
			if err == redis.Nil {
				dbData, err := s.userRepo.ViewJobDetailsById(uint64(v.JobId))
				if err != nil {
					return
				}
				err = s.rdb.AddToRedis(ctx, v.JobId, dbData)
				if err != nil {
					return
				}
				jobData = dbData
			} else {
				err = json.Unmarshal([]byte(val), &jobData)
				if err == redis.Nil {
					return
				}
				if err != nil {
					return
				}
			}

			result1, t := ProcessingData(jobData, v)
			if t {
				c1 <- result1
			}

		}(v)

	}
	go func() {
		wg.Wait()
		close(c1)
	}()
	for val := range c1 {
		JobsResult = append(JobsResult, val)
	}
	return JobsResult, nil

}
func ProcessingData(j models.Job, rj models.RequestJob) (models.NewRequestJob, bool) {

	if !(j.Min_Np <= rj.NoticePeriod && rj.NoticePeriod <= j.Max_Np) {
		return models.NewRequestJob{}, false
	}
	if !(0 < rj.Budget && rj.Budget <= j.Budget) {
		return models.NewRequestJob{}, false
	}
	if !(j.MinExp <= rj.Exp && rj.Exp <= j.MaxExp) {
		return models.NewRequestJob{}, false
	}
	if !(containsLocation(j.JobLocation, rj.JobLocation)) {
		return models.NewRequestJob{}, false
	}
	// if !(containsQualification(j.Qualification, rj.Qualification)) {
	// 	return models.NewRequestJob{},false
	// }
	if !(containsTechn(j.Technology, rj.Technology)) {
		return models.NewRequestJob{}, false
	}
	// if !(containsShift(j.Shift, rj.Shift)) {
	// 	return models.NewRequestJob{},false
	// }
	// if !(containsJobType(j.JobType, rj.JobType)) {
	// 	return models.NewRequestJob{},false
	// }
	if !(containWorkMode(j.WorKMode, rj.WorKMode)) {
		return models.NewRequestJob{}, false
	}
	return models.NewRequestJob{
		Name: rj.Name,
	}, true

}
func containWorkMode(jw []models.WorKMode, rj uint) bool {
	var jwid []uint
	for _, id := range jw {
		jwid = append(jwid, (id.ID))
	}
	for _, id := range jwid {
		if id == rj {
			return true
		}
	}
	return false
}

func containsLocation(jl []models.Location, request []uint) bool {
	var jlid []uint

	for _, id := range jl {
		jlid = append(jlid, id.ID)
	}
	for _, rid := range request {
		if !slices.Contains(jlid, rid) {
			return false
		}
	}
	return true
}

// func containsQualification(jq []models.Qualification, request []uint) bool {
// 	var jqid []uint

// 	for _, id := range jq {
// 		jqid = append(jqid, id.ID)
// 	}
// 	for _, rid := range request {
// 		if !slices.Contains(jqid, rid) {
// 			return false
// 		}
// 	}
// 	return true
// }

func containsTechn(job []models.Technology, request []uint) bool {
	var jt []uint

	for _, jtid := range job {
		jt = append(jt, jtid.ID)
	}
	for _, jtid := range request {
		if !slices.Contains(jt, jtid) {
			return false
		}
	}
	return true
}

// func containsShift(job []models.Shift, request []uint) bool {
// 	var js []uint

// 	for _, jsid := range job {
// 		js = append(js, jsid.ID)
// 	}
// 	for _, jsid := range request {
// 		if !slices.Contains(js, jsid) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func containsJobType(job []models.JobType, request []uint) bool {
// 	var jjt []uint

// 	for _, jjtid := range job {
// 		jjt = append(jjt, jjtid.ID)
// 	}
// 	for _, jjtid := range request {
// 		if !slices.Contains(jjt, jjtid) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func ProcessingOther5Data(v model.Job, m model.JobRequest) bool {
// 	if !(containsTechn(v.Technology, m.Technology)) {
// 		return false
// 	}
// 	if !(containsShift(v.Shift, m.Shift)) {
// 		return false
// 	}
// 	if !(containsJobType(v.JobType, m.JobType)) {
// 		return false
// 	}
// 	if !(v.WorkMode == m.WorkMode) {
// 		return false
// 	}

// 	return true
// }
