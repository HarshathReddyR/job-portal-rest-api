package handlers

import (
	"errors"
	"fmt"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/repository"
	"job-portal-api/internal/services"
	"job-portal-api/redies"

	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service services.ServiceMethod
	auth    *auth.Auth
	// redis   redies.RedisMethodsv
}

func NewHandler(service services.ServiceMethod, auth *auth.Auth, rd *redies.Redis) (*Handler, error) {
	if service == nil || auth == nil {
		return nil, errors.New("interface and structure cannot be null")
	}
	return &Handler{
		service: service,
		auth:    auth,
		// redis:   rd,
	}, nil
}

func API(a *auth.Auth, sc repository.UserRepo, redisLayer redies.RedisMethods) *gin.Engine {
	r := gin.New() //create a new engine
	m, err := middleware.NewMid(a)
	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}

	// rdb := database.ConnectionToRedis()

	// redisLayer := redies.NewRedis(rdb)

	ms, err := services.NewService(sc, redisLayer)
	if err != nil {
		log.Panic().Msg("handlers not setup")
	}
	h := Handler{
		service: ms,
		auth:    a,
		// redis:   redisLayer,
	}

	r.Use(m.Log(), gin.Recovery())
	r.GET("/check", check)
	r.POST("/api/register", h.Signup)
	r.POST("/api/login", h.Login)
	r.POST("/api/companies", m.Authenticate(h.CreateCompany))
	r.GET("/api/companies", m.Authenticate(h.FetchCompanies))
	r.GET("/api/companies/:id", m.Authenticate(h.FetchCompanyById))
	r.POST("/api/companies/:id/jobs", m.Authenticate(h.CreateJob))
	r.GET("/api/jobs", m.Authenticate(h.FetchJob))
	r.GET("/api/jobs/:id", m.Authenticate(h.FetchJobById))
	r.GET("/api/companies/:id/jobs", m.Authenticate(h.FetchJobByCompanyId))
	r.POST("/api/jobpost/", (h.ProcessingJob))
	r.POST("/api/forgotpassword",h.ForgotPassword)
	r.POST("/api/updatepassword",h.UpdatePassword)
	return r //return prepared gin engine

}
func check(c *gin.Context) {
	time.Sleep(time.Second * 3)
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})
	}
}
