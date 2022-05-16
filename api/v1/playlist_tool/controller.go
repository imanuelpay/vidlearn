package playlist_tool

import (
	"net/http"
	"strconv"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/api/v1/playlist_tool/request"
	"vidlearn-final-projcect/api/v1/playlist_tool/response"
	playlistToolBusiness "vidlearn-final-projcect/business/playlist_tool"

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

// GetAllToolsByPlaylist Get Tools by Playlist
// @Description Get Tools by Playlist
// @Summary Get Tools by Playlist
// @Tags Playlist
// @Accept json
// @Produce json
// @Param id path int true "Playlist ID"
// @Success 200 {object} common.DefaultDataResponse{data=response.GetToolsByPlaylist}
// @Security ApiKeyAuth
// @Router /v1/playlist/{id}/tool [get]
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

// GetAllPlaylistsByTool Get Playlists by Tool
// @Description Get Playlists by Tool
// @Summary Get Playlists by Tool
// @Tags Tool
// @Accept json
// @Produce json
// @Param id path int true "Tool ID"
// @Success 200 {object} common.DefaultDataResponse{data=response.GetPlaylistsByTool}
// @Security ApiKeyAuth
// @Router /v1/tool/{id}/playlist [get]
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

	if len(playlistsByTool) < 1 {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: "Playlist not found",
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

// CreatePlaylistTool Create Playlist Tool
// @Description Create Playlist Tool
// @Summary Create Playlist Tool
// @Tags Playlist
// @Accept json
// @Produce json
// @Param id path string true "Playlist ID"
// @Param body body request.CreatePlaylistToolRequest true "Create Playlist Tool"
// @Success 200 {object} common.DefaultDataResponse{data=response.GetToolsByPlaylistID}
// @Security ApiKeyAuth
// @Router /v1/playlist/{id}/tool [post]
func (controller *Controller) CreatePlaylistTool(c echo.Context) error {
	playlistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	createPlaylistToolRequest := new(request.CreatePlaylistToolRequest)
	if err := c.Bind(createPlaylistToolRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := *createPlaylistToolRequest.ToSpec()

	playlistTool, err := controller.service.CreatePlaylistTool(&request, playlistID)
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

// DeletePlaylistTool Delete Playlist Tool
// @Description Delete Playlist Tool
// @Summary Delete Playlist Tool
// @Tags Playlist
// @Accept json
// @Produce json
// @Param id_playlist path string true "Playlist ID"
// @Param id_tool path string true "Tool ID"
// @Success 200 {object} common.DefaultDataResponse{data=response.GetToolsByPlaylistID}
// @Security ApiKeyAuth
// @Router /v1/playlist/{id_playlist}/tool/{id_tool} [delete]
func (controller *Controller) DeletePlaylistTool(c echo.Context) error {
	playlistID, err := strconv.Atoi(c.Param("id_playlist"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	toolID, err := strconv.Atoi(c.Param("id_tool"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	playlistTool, err := controller.service.DeletePlaylistTool(playlistID, toolID)
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
