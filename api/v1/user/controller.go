package user

import (
	"net/http"
	"strconv"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/api/v1/user/request"
	"vidlearn-final-projcect/api/v1/user/response"
	userBusiness "vidlearn-final-projcect/business/user"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service userBusiness.Service
}

func CreateController(service userBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	users, err := controller.service.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetAllUserResponse(users)
	responseData := common.DefaultDataResponse{
		Message: "Get all users successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) GetByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	user, err := controller.service.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetUserByIDResponse(user)
	responseData := common.DefaultDataResponse{
		Message: "Get user by id successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Update(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	updateUserRequest := new(request.UpdateUserRequest)
	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := *updateUserRequest.ToSpec()

	user, err := controller.service.UpdateUser(&request, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetUserByIDResponse(user)
	responseData := common.DefaultDataResponse{
		Message: "Update user successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) UpdatePassword(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	updatePasswordRequest := new(request.UpdatePasswordRequest)
	if err := c.Bind(updatePasswordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := *updatePasswordRequest.ToSpec()

	user, err := controller.service.UpdateUserPassword(&request, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetUserByIDResponse(user)
	responseData := common.DefaultDataResponse{
		Message: "Update user password successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Delete(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	user, err := controller.service.DeleteUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	responseData := common.DefaultDataResponse{
		Message: "Delete user successfully",
		Data:    user,
	}

	return c.JSON(http.StatusOK, responseData)
}
