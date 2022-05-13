package category

import (
	"net/http"
	"strconv"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/api/v1/category/request"
	"vidlearn-final-projcect/api/v1/category/response"
	categoryBusiness "vidlearn-final-projcect/business/category"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service categoryBusiness.Service
}

func CreateController(service categoryBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	categories, err := controller.service.GetAllCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetAllCategoryResponse(categories)
	responseData := common.DefaultDataResponse{
		Message: "Get all categories successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) GetByID(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	category, err := controller.service.GetCategoryByID(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetCategoryByIDResponse(category)
	responseData := common.DefaultDataResponse{
		Message: "Get category by id successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Create(c echo.Context) error {
	createCategoryRequest := new(request.CreateCategoryRequest)
	if err := c.Bind(createCategoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	request := *createCategoryRequest.ToSpec()

	category, err := controller.service.CreateCategory(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetCategoryByIDResponse(category)
	responseData := common.DefaultDataResponse{
		Message: "Create category successfully",
		Data:    response,
	}

	return c.JSON(http.StatusCreated, responseData)
}

func (controller *Controller) Update(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	updateCategoryRequest := new(request.CreateCategoryRequest)
	if err := c.Bind(updateCategoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	request := *updateCategoryRequest.ToSpec()

	category, err := controller.service.UpdateCategory(&request, categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetCategoryByIDResponse(category)
	responseData := common.DefaultDataResponse{
		Message: "Update category successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Delete(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	category, err := controller.service.DeleteCategory(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetCategoryByIDResponse(category)
	responseData := common.DefaultDataResponse{
		Message: "Delete category successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
