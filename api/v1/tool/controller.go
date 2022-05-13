package tool

import (
	"net/http"
	"strconv"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/api/v1/tool/request"
	"vidlearn-final-projcect/api/v1/tool/response"
	toolBusiness "vidlearn-final-projcect/business/tool"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service toolBusiness.Service
}

func CreateController(service toolBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	tools, err := controller.service.GetAllTool()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: "Get all tool failed",
			Data:    nil,
		})
	}

	response := response.CreateGetAllToolResponse(tools)
	responseData := common.DefaultDataResponse{
		Message: "Get all tools successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) GetByID(c echo.Context) error {
	toolID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	tool, err := controller.service.GetToolByID(toolID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: "Get tool failed",
			Data:    nil,
		})
	}

	response := response.CreateGetToolByIDResponse(tool)
	responseData := common.DefaultDataResponse{
		Message: "Get tool by id successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Create(c echo.Context) error {
	createToolRequest := new(request.CreateToolRequest)
	if err := c.Bind(createToolRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	request := *createToolRequest.ToSpec()

	tool, err := controller.service.CreateTool(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: "Create tool failed",
			Data:    nil,
		})
	}

	response := response.CreateGetToolByIDResponse(tool)
	responseData := common.DefaultDataResponse{
		Message: "Create tool successfully",
		Data:    response,
	}

	return c.JSON(http.StatusCreated, responseData)
}

func (controller *Controller) Update(c echo.Context) error {
	toolID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	updateToolRequest := new(request.CreateToolRequest)
	if err := c.Bind(updateToolRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	request := *updateToolRequest.ToSpec()

	tool, err := controller.service.UpdateTool(&request, toolID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: "Update tool failed",
			Data:    nil,
		})
	}

	response := response.CreateGetToolByIDResponse(tool)
	responseData := common.DefaultDataResponse{
		Message: "Update tool successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Delete(c echo.Context) error {
	toolID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	tool, err := controller.service.DeleteTool(toolID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: "Delete tool failed",
			Data:    nil,
		})
	}

	responseData := common.DefaultDataResponse{
		Message: "Delete tool successfully",
		Data:    tool,
	}

	return c.JSON(http.StatusOK, responseData)
}
