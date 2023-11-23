package redies

import (
	"context"
	"encoding/json"
	"job-portal-api/internal/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	redisdb *redis.Client
}
//go:generate mockgen -source=redis.go -destination=redis_mock.go -package=redies
type RedisMethods interface {
	AddToRedis(ctx context.Context, jid uint, jobData models.Job) error
	GetDataFromRedis(ctx context.Context, jid uint) (string, error)
}

func NewRedis(redis *redis.Client) RedisMethods {
	return &Redis{
		redisdb: redis,
	}
}
func (r *Redis) AddToRedis(ctx context.Context, jid uint, jobData models.Job) error {
	jobid := strconv.FormatUint(uint64(jid), 10)
	val, err := json.Marshal(jobData)
	if err != nil {
		return err
	}
	err = r.redisdb.Set(ctx, jobid, val, 1*time.Minute).Err()
	return err
}
func (r *Redis) GetDataFromRedis(ctx context.Context, jid uint) (string, error) {
	jobID := strconv.FormatUint(uint64(jid), 10)
	str, err := r.redisdb.Get(ctx, jobID).Result()
	return str, err
}
