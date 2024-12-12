package controller

import (
	"fmt"
	"golang-chap49/config"
	"golang-chap49/database"
	"golang-chap49/helper"
	"golang-chap49/models"
	"golang-chap49/service"
	"golang-chap49/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	Service service.Service
	Log     *zap.Logger
	Cacher  database.Cacher
	Config  config.Configuration
}

func NewUserController(service service.Service, log *zap.Logger, cacher database.Cacher, config config.Configuration) *UserController {
	return &UserController{
		Service: service,
		Log:     log,
		Cacher:  cacher,
		Config:  config,
	}
}

// CreateUserController godoc
// @Summary      Create a new user
// @Description  Register a new user with a provided request body
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        registerRequest  body     models.RegisterRequest  true  "User Registration Request Body"
// @Success      201             {object} helper.HTTPResponse   "User created successfully"
// @Failure      400             {object} helper.HTTPResponse   "Invalid request body"
// @Failure      500             {object} helper.HTTPResponse   "Failed to create user"
// @Router       /auth/register [post]
func (ctrl *UserController) CreateUserController(c *gin.Context) {
	var registerReq models.RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}

	err := ctrl.Service.User.CreateUser(registerReq)
	if err != nil {
		ctrl.Log.Error("Failed to create user", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to create user", http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(c, nil, "User created successfully", http.StatusCreated)
}

// LoginController godoc
// @Summary      User Login
// @Description  Authenticate a user and return a token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        loginRequest  body     models.LoginRequest  true  "Login Request Body"
// @Success      200           {object} helper.HTTPResponse   "User logged in successfully"
// @Failure      400           {object} helper.HTTPResponse   "Invalid request body"
// @Failure      401           {object} helper.HTTPResponse   "Failed to login user"
// @Router       /auth/login [post]
func (ctrl *UserController) LoginController(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}
	otpReq, err := ctrl.Service.User.Login(loginReq)
	if err != nil {
		ctrl.Log.Error("Failed to login user", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to login user", http.StatusUnauthorized)
		return
	}

	if err := utils.SendOTPEmail(ctrl.Config.EmailAPI, ctrl.Config.EmailClient, otpReq); err != nil {
		ctrl.Log.Error("Failed to send OTP email", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to send OTP email", http.StatusInternalServerError)
		return
	}

	ctrl.Service.Otp.CreateOtp(otpReq)
	helper.ResponseOK(c, nil, "OTP sent successfully", http.StatusOK)
}

// CheckEmailUserController godoc
// @Summary      Check if email is already registered
// @Description  Verify if a user with the given email already exists in the system
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        email  body     models.CheckEmailRequest  true  "Email to check"
// @Success      200    {object} helper.HTTPResponse     "Email check result"
// @Failure      400    {object} helper.HTTPResponse     "Invalid request body"
// @Failure      500    {object} helper.HTTPResponse     "Failed to check user email"
// @Router       /auth/check-email [post]
func (ctrl *UserController) CheckEmailUserController(c *gin.Context) {
	request := models.CheckEmailRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}

	existedUser, err := ctrl.Service.User.CheckUserEmail(request.Email)
	if err != nil {
		ctrl.Log.Error("Failed to check user email", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to check user email", http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(c, existedUser, "User email exists", http.StatusOK)
}

// ResetUserPasswordController godoc
// @Summary      Reset user password
// @Description  Reset the password for a user using a provided request body
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        resetRequest  body     models.LoginRequest  true  "User password reset request body"
// @Success      200           {object} helper.HTTPResponse   "User password reset successfully"
// @Failure      400           {object} helper.HTTPResponse   "Invalid request body"
// @Failure      500           {object} helper.HTTPResponse   "Failed to reset user password"
// @Router       /auth/reset-password [PATCH]
func (ctrl *UserController) ResetUserPasswordController(c *gin.Context) {
	var resetReq models.LoginRequest
	if err := c.ShouldBindJSON(&resetReq); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}

	err := ctrl.Service.User.ResetUserPassword(resetReq)
	if err != nil {
		ctrl.Log.Error("Failed to reset user password", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to reset user password", http.StatusInternalServerError)
		return
	}
	helper.ResponseOK(c, nil, "User password reset successfully", http.StatusOK)
}

func (ctrl *UserController) CheckOTP(c *gin.Context) {
	otpStr := c.Param("otp")

	otp := helper.StringToInt(otpStr)
	existedUser, err := ctrl.Service.Otp.GetUserByOTP(otp)
	if err != nil {
		ctrl.Log.Error("Failed to get user by OTP", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Failed to get user by OTP", http.StatusInternalServerError)
		return
	}

	userIDStr := helper.IntToString(existedUser.ID)
	key := fmt.Sprintf("UserID-%s", userIDStr)
	token := helper.GenerateToken(userIDStr, ctrl.Config.SecretKey)
	loginResponse := utils.LoginResponse{
		ID:    key,
		Role:  existedUser.Role,
		Token: token,
	}

	ctrl.Cacher.SaveToken(key, token)
	helper.ResponseOK(c, loginResponse, "OTP verified successfully", http.StatusOK)
}
