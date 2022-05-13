package playlist_tool

import (
	"net/http"
	"strconv"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/api/v1/playlist_tool/response"
	playlistToolBusiness "vidlearn-final-projcect/business/playlist_tool"
	"vidlearn-final-projcect/business/playlist_tool/spec"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service playlistToolBusiness.Service
}

func CreateController(service playlistToolBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAllToolsByPlaylist(c echo.Context) error {
	playlistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	toolsByPlaylist, err := controller.service.GetToolByPlaylist(playlistID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetToolsByPlaylistResponse(toolsByPlaylist)
	responseData := common.DefaultDataResponse{
		Message: "Get all tools by playlist successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) GetAllPlaylistsByTool(c echo.Context) error {
	toolID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	playlistsByTool, err := controller.service.GetPlaylistByTool(toolID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetPlaylistsByToolResponse(playlistsByTool)
	responseData := common.DefaultDataResponse{
		Message: "Get all playlists by tool successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) CreatePlaylistTool(c echo.Context) error {
	var upsertPlaylistToolSpec spec.UpsertPlaylistToolSpec
	if err := c.Bind(&upsertPlaylistToolSpec); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := c.Validate(&upsertPlaylistToolSpec); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	playlistTool, err := controller.service.CreatePlaylistTool(&upsertPlaylistToolSpec)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetToolsByPlaylistIDResponse(playlistTool)
	responseData := common.DefaultDataResponse{
		Message: "Create playlist tool successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) DeletePlaylistTool(c echo.Context) error {
	playlistToolID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	playlistTool, err := controller.service.DeletePlaylistTool(playlistToolID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetToolsByPlaylistIDResponse(playlistTool)

	responseData := common.DefaultDataResponse{
		Message: "Delete playlist tool successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
