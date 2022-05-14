package auth

import (
	"net/http"
	"time"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/api/v1/auth/request"
	"vidlearn-final-projcect/api/v1/auth/response"
	userBusiness "vidlearn-final-projcect/business/user"
	"vidlearn-final-projcect/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	userService userBusiness.Service
	config      *config.AppConfig
}

func CreateController(userService userBusiness.Service, config *config.AppConfig) *Controller {
	return &Controller{
		userService: userService,
		config:      config,
	}
}

// Login func for login user.
// @Description Login user.
// @Summary Login user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body request.LoginUserRequest true "Login User"
// @Success 201 {object} common.DefaultDataResponse{data=response.LoginResponse}
// @Router /v1/auth/login [post]
func (controller *Controller) Login(c echo.Context) error {
	var loginUserRequest *request.LoginUserRequest
	if err := c.Bind(&loginUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := loginUserRequest.ToSpec()

	user, err := controller.userService.GetUserByEmailAndPassword(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	claims := jwt.MapClaims{
		"userId": user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
	}

	JWTKey := []byte(controller.config.App.Key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString(JWTKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateLoginResponse(user.Name, user.Email, jwt)
	responseData := common.DefaultDataResponse{
		Message: "Login successfully",
		Data:    response,
	}

	return c.JSON(http.StatusCreated, responseData)
}

// Register func for register user.
// @Description Register user.
// @Summary Register user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body request.RegisterUserRequest true "Register User"
// @Success 201 {object} common.DefaultDataResponse{data=response.RegisterResponse}
// @Router /v1/auth/register [post]
func (controller *Controller) Register(c echo.Context) error {
	var registerUserRequest *request.RegisterUserRequest
	if err := c.Bind(&registerUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := registerUserRequest.ToSpec()

	user, err := controller.userService.CreateUser(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateRegisterResponse(user.Name, user.Email)
	responseData := common.DefaultDataResponse{
		Message: "Register successfully, please check your email to verify your account",
		Data:    response,
	}

	return c.JSON(http.StatusCreated, responseData)
}

// Verify func for register user.
// @Description Verify user.
// @Summary Verify user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param token path string true "Verify Token"
// @Success 200 {object} common.DefaultDataResponse
// @Router /v1/verify/{token} [get]
func (controller *Controller) Verify(c echo.Context) error {
	verifyCode := c.Param("token")
	if verifyCode == "" {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: "Verify code is empty",
			Data:    nil,
		})
	}

	user, err := controller.userService.VerifyEmail(verifyCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.DefaultDataResponse{
		Message: "Verify successfully, please login",
		Data:    user.Email,
	})
}

// ForgotPassword func for Forgot Password user.
// @Description Forgot Password user.
// @Summary Forgot Password user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body request.ForgotPasswordRequest true "Forgot Password"
// @Success 200 {object} common.DefaultDataResponse
// @Router /v1/forgot-password [post]
func (controller *Controller) ForgotPassword(c echo.Context) error {
	var forgotPasswordRequest *request.ForgotPasswordRequest

	if err := c.Bind(&forgotPasswordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := forgotPasswordRequest.ToSpec()

	user, err := controller.userService.ForgotPassword(request.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.DefaultDataResponse{
		Message: "Forgot password successfully, please check your email to reset your password",
		Data:    user.Email,
	})
}

// ResetPassword func for Reset Password user.
// @Description Reset Password user.
// @Summary Reset Password user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param token path string true "Reset Token"
// @Param body body request.ResetPasswordRequest true "Reset Password"
// @Success 200 {object} common.DefaultDataResponse
// @Router /v1/reset-password/{token} [POST]
func (controller *Controller) ResetPassword(c echo.Context) error {
	verifyCode := c.Param("token")
	if verifyCode == "" {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: "Verify code is empty",
			Data:    nil,
		})
	}

	var resetPasswordRequest *request.ResetPasswordRequest
	if err := c.Bind(&resetPasswordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := resetPasswordRequest.ToSpec()

	user, err := controller.userService.ResetPassword(request, verifyCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.DefaultDataResponse{
		Message: "Reset password successfully, please login",
		Data:    user.Email,
	})
}
