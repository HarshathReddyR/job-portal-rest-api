package handlers

// import (
// 	"context"
// 	"errors"
// 	"job-portal-api/internal/middleware"
// 	"job-portal-api/internal/models"
// 	"job-portal-api/internal/services"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert/v2"
// 	"go.uber.org/mock/gomock"
// )

// func TestHandler_FetchJobById(t *testing.T) {
// 	// type args struct {
// 	// 	c *gin.Context
// 	// }
// 	tests := []struct {
// 		name string
// 		// h                  *Handler
// 		// args               args
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name: "missing trace id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"msg":"Internal Server Error"}`,
// 		},
// 		{
// 			name: "invalid job id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
// 				// ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "a"})

// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := services.NewMockServiceMethod(mc)

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"msg":"error found at conversion.."}`,
// 		},
// 		{
// 			name: "featching job failed",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})

// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := services.NewMockServiceMethod(mc)
// 				ms.EXPECT().FetchJobById(gomock.Any()).Return(models.Job{}, errors.New("error found at conversion..")).AnyTimes()
// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"msg":"job fetching failed"}`,
// 		},
// 		{
// 			name: "success",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})

// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := services.NewMockServiceMethod(mc)
// 				ms.EXPECT().FetchJobById(gomock.Any()).Return(models.Job{
// 					Cid:         1,
// 					JobRole:     "hhh",
// 					Description: "llllll",
// 				}, nil).AnyTimes()
// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   ` {"cid":1,"Role":"hhh","description":"llllll","minimum_notice_period":0,"maximum_notice_period":0,"budget":0,"JobLocation":null,"Technology":null,"WorKMode":null,"minimum_experience":0,"maximum_experience":0,"Qualification":null,"Shift":null,"JobType":null}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, sm := tt.setup()

// 			h := Handler{
// 				service: sm,
// 			}

// 			h.FetchJobById(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, strings.TrimSpace(tt.expectedResponse), strings.TrimSpace(rr.Body.String()))
// 		})
// 	}
// }

// func TestHandler_FetchJobByCompanyId(t *testing.T) {
// 	// type args struct {
// 	// 	c *gin.Context
// 	// }
// 	tests := []struct {
// 		name string
// 		// h                  *Handler
// 		// args               args
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name: "missing trace id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"msg":"Internal Server Error"}`,
// 		},
// 		{
// 			name: "invalid company id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
// 				// ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "a"})

// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := services.NewMockServiceMethod(mc)

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"msg":"error found at conversion.."}`,
// 		},
// 		{
// 			name: "featching company failed",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})

// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := services.NewMockServiceMethod(mc)
// 				ms.EXPECT().FetchJobByCompanyId(gomock.Any()).Return([]models.Job{}, errors.New("error found at conversion..")).AnyTimes()
// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"msg":"job fetching with company failed"}`,
// 		},
// 		{
// 			name: "success",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")

// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
// 				mc := gomock.NewController(t)
// 				ms := services.NewMockServiceMethod(mc)

// 				ms.EXPECT().FetchJobByCompanyId(gomock.Any()).Return([]models.Job{}, nil).AnyTimes()
// 				return c, rr, ms
// 			},
// 			expectedStatusCode: 200,
// 			expectedResponse:   `[]`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, sm := tt.setup()

// 			h := Handler{
// 				service: sm,
// 			}
// 			h.FetchJobByCompanyId(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			// fmt.Println("EE::", tt.expectedResponse)
// 			// fmt.Println("GG::", strings.TrimSpace(rr.Body.String()) == strings.TrimSpace(tt.expectedResponse))
// 			// fmt.Printf("%q", tt.expectedResponse)
// 			// fmt.Println("")
// 			// fmt.Printf("%q", rr.Body.String())
// 			// t.Error("")
// 			assert.Equal(t, strings.TrimSpace(tt.expectedResponse), strings.TrimSpace(rr.Body.String()))
// 		})
// 	}
// }

// func TestHandler_ProcessingJob(t *testing.T) {
// 	// type args struct {
// 	// 	c *gin.Context
// 	// }
// 	tests := []struct {
// 		name string
// 		// h    *Handler
// 		// args args
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name: "error in converting json to struct",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", strings.NewReader(`{}`))
// 				// ctx := httpRequest.Context()
// 				// ctx = context.WithValue(ctx, middleware.TraceIdKey, "123")
// 				c.Request = httpRequest
// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"msg2":"decode fail"}`,
// 		},
// 		{
// 			name: "error while fetching jobs from service",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", strings.NewReader(`[
// 					{
// 						"name":"a",
// 						"jobid":1,
// 						"role": "Software developer",
// 						"description": "java development",
// 						"noticPeriod": 30,
// 						"budget": 80000,
// 						"job_locations": [1],
// 						"job_technology": [1,2],
// 						"job_workmode": 1,
// 						"experience": 3,
// 						"job_qualification": [1],
// 						"job_type": [1]
// 					}
// 				]`))
// 				// ctx := httpRequest.Context()
// 				// ctx = context.WithValue(ctx, middleware.TraceIdKey, "123")
// 				c.Request = httpRequest
// 				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "abc"})
// 				mc := gomock.NewController(t)
// 				ms := services.NewMockServiceMethod(mc)

// 				ms.EXPECT().ProcessingJob(gomock.Any(), gomock.Any()).Return([]models.NewRequestJob{}, errors.New("test service error")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"msg":"Internal Server Error"}`,
// 		},
// 		{
// 			name: "success",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", strings.NewReader(`
// 					[{
// 						"name":"a",
// 						"jobid":1,
// 						"role": "Software developer",
// 						"description": "java development",
// 						"noticPeriod": 30,
// 						"budget": 80000,
// 						"job_locations": [1],
// 						"job_technology": [1,2],
// 						"job_workmode": 1,
// 						"experience": 3,
// 						"job_qualification": [1],
// 						"job_type": [1]
// 					}]
// 				`))
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := services.NewMockServiceMethod(mc)
// 				ms.EXPECT().ProcessingJob(gomock.Any(), gomock.Any()).Return([]models.NewRequestJob{}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: 200,
// 			expectedResponse:   "[]",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()
// 			h := &Handler{
// 				service: ms,
// 			}
// 			h.ProcessingJob(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, strings.TrimSpace(tt.expectedResponse), strings.TrimSpace(rr.Body.String()))
// 		})
// 	}
// }

// func TestHandler_FetchJob(t *testing.T) {
// 	// type args struct {
// 	// 	c *gin.Context
// 	// }
// 	tests := []struct {
// 		name string
// 		// h    *Handler
// 		// args args
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name: "missing trace id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.ServiceMethod) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"msg":"Internal Server Error"}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, sm := tt.setup()

// 			h := Handler{
// 				service: sm,
// 			}

// 			h.FetchJob(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, strings.TrimSpace(tt.expectedResponse), strings.TrimSpace(rr.Body.String()))
// 		})
// 	}
// }
