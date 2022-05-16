package auth

import (
	"net/http"
	"strings"
	"time"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/api/v1/auth/request"
	"vidlearn-final-projcect/api/v1/auth/response"
	requestUser "vidlearn-final-projcect/api/v1/user/request"
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

// GetProfile func for get user profile.
// @Description Get user profile.
// @Summary Get user profile
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} common.DefaultDataResponse{data=response.ProfileResponse}
// @Security ApiKeyAuth
// @Router /v1/auth/profile [get]
func (controller *Controller) GetProfile(c echo.Context) error {
	signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(signature) < 2 {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Message: "Invalid token",
			Data:    nil,
		})
	}

	if signature[0] != "Bearer" {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Message: "Invalid token",
			Data:    nil,
		})
	}

	jwtToken := signature[1]
	user, err := controller.userService.GetUserLogin(jwtToken)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateProfileResponse(user)
	responseData := common.DefaultDataResponse{
		Message: "Get profile successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

// UpdateProfile func for Update Profile.
// @Description Update Profile.
// @Summary Update Profile
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body request.UpdateProfileRequest true "Update Profile"
// @Success 200 {object} common.DefaultDataResponse{data=response.ProfileResponse}
// @Router /v1/auth/update-profile [put]
func (controller *Controller) UpdateProfile(c echo.Context) error {
	signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(signature) < 2 {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Message: "Invalid token",
			Data:    nil,
		})
	}

	if signature[0] != "Bearer" {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Message: "Invalid token",
			Data:    nil,
		})
	}

	jwtToken := signature[1]
	user, err := controller.userService.GetUserLogin(jwtToken)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	updateUserRequest := new(request.UpdateProfileRequest)
	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := *updateUserRequest.ToSpec(user.Role)

	user, err = controller.userService.UpdateUser(&request, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateProfileResponse(user)
	responseData := common.DefaultDataResponse{
		Message: "Update user successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
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

	response := response.CreateRegisterResponse(user.Name, user.Email, user.VerifyCode)
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

	response := response.CreateVerifyResponse(user.Name, user.Email, user.VerifiedAt)
	responseData := common.DefaultDataResponse{
		Message: "Verify successfully, please login",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
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

	response := response.CreateForgotPasswordResponse(user.Name, user.Email, user.VerifyCode)
	responseData := common.DefaultDataResponse{
		Message: "Forgot password successfully, please check your email to verify your account",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
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

	response := response.CreateResetPasswordResponse(user.Name, user.Email)
	responseData := common.DefaultDataResponse{
		Message: "Reset password successfully, please login",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

// ChangePassword func Change Password
// @Description Change Password
// @Summary Change Password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body requestUser.UpdatePasswordRequest true "Change Password"
// @Success 200 {object} common.DefaultDataResponse{data=response.ResetPasswordResponse}
// @Security ApiKeyAuth
// @Router /api/v1/auth/profile/change-password [put]
func (controller *Controller) ChangePassword(c echo.Context) error {
	signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(signature) < 2 {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Message: "Invalid token",
			Data:    nil,
		})
	}

	if signature[0] != "Bearer" {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Message: "Invalid token",
			Data:    nil,
		})
	}

	jwtToken := signature[1]
	user, err := controller.userService.GetUserLogin(jwtToken)
	if err != nil {
		return c.JSON(http.StatusForbidden, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}
	updatePasswordRequest := new(requestUser.UpdatePasswordRequest)
	if err := c.Bind(updatePasswordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := *updatePasswordRequest.ToSpec()

	user, err = controller.userService.UpdateUserPassword(&request, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateResetPasswordResponse(user.Name, user.Email)
	responseData := common.DefaultDataResponse{
		Message: "Change password successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
