package handlers

import (
	"encoding/json"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// Signup is a method for the handler struct which handles user registration
func (h *Handler) Signup(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		// If the traceId isn't found in the request, log an error and return
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Define a NewUser variable
	var userData models.NewUser

	// Attempt to decode JSON from the request body into the NewUser variable
	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		// If there is an error in decoding, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Create a new validator and validate the NewUser variable
	validate := validator.New()
	err = validate.Struct(userData)
	if err != nil {
		// If validation fails, log the error and return
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Name, Email and Password"})
		return
	}

	// Attempt to create the user
	usr, err := h.service.CreateUser(ctx, userData)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "user signup failed"})
		return
	}
	// If everything goes right, respond with the created user
	c.JSON(http.StatusOK, usr)
}

// Login is a method for the handler struct which handles user login
func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Define a new struct for login data
	var login struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	// Attempt to decode JSON from the request body into the login variable
	err := json.NewDecoder(c.Request.Body).Decode(&login)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Create a new validator and validate the login variable
	validate := validator.New()
	err = validate.Struct(login)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Email and Password"})
		return
	}

	// Attempt to authenticate the user with the email and password
	claims, err := h.service.UserLogin(ctx, login.Email, login.Password)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "login failed"})
		return
	}

	// Define a new struct for the token
	var tkn struct {
		Token string `json:"token"`
	}

	// Generate a new token and put it in the Token field of the token struct
	tkn.Token, err = h.auth.GenerateToken(claims)
	if err != nil {
		log.Error().Err(err).Msg("generating token")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	// If everything goes right, respond with the token
	c.JSON(http.StatusOK, tkn)

}
func (h *Handler) ForgotPassword(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var ru1 models.Recive1

	//here we are decoding the Json body
	err := json.NewDecoder(c.Request.Body).Decode(&ru1)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	//here we are validating the every fields
	validate := validator.New()
	err = validate.Struct(ru1)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Email and DOB"})
		return
	}
	err = h.service.ForgotPassword(ctx, ru1)
	if err != nil {
		log.Error().Err(err).Msg("Forgot Password")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, "OTP Sent")
}
func (h *Handler) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var ru2 models.Recive2
	err := json.NewDecoder(c.Request.Body).Decode(&ru2)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()
	err = validate.Struct(ru2)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "please provide Email and Password"})
		return
	}
	err = h.service.UpdatePassword(ctx, ru2)
	if err != nil {
		log.Error().Err(err).Msg("update Password ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, "Password is updated")
}
