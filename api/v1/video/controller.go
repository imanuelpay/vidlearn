package video

import (
	"net/http"
	"strconv"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/api/v1/video/request"
	"vidlearn-final-projcect/api/v1/video/response"
	videoBusiness "vidlearn-final-projcect/business/video"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service videoBusiness.Service
}

func CreateController(service videoBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// GetVideoByPlaylist func for Get Video By Playlist.
// @Description Get Video By Playlist.
// @Summary Get Video By Playlist
// @Tags Video
// @Accept json
// @Produce json
// @Param id path string true "Playlist ID"
// @Success 200 {object} common.DefaultDataResponse{data=response.Playlist}
// @Security ApiKeyAuth
// @Router /v1/video/{id}/playlist [get]
func (controller *Controller) GetVideoByPlaylist(c echo.Context) error {
	playlistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	videos, err := controller.service.GetVideoByPlaylist(playlistID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	if len(videos) < 1 {
		return c.JSON(http.StatusNotFound, common.DefaultDataResponse{
			Message: "Video not found",
			Data:    nil,
		})
	}

	response := response.CreateGetVideoByPlaylistResponse(videos)
	responseData := common.DefaultDataResponse{
		Message: "Get video by playlist successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

// GetByID func for Get Video By ID.
// @Description Get Video By ID.
// @Summary Get Video By ID
// @Tags Video
// @Accept json
// @Produce json
// @Param id path string true "Video ID"
// @Success 200 {object} common.DefaultDataResponse{data=response.GetVideoByIDByPlaylistResponse}
// @Security ApiKeyAuth
// @Router /v1/video/{id} [get]
func (controller *Controller) GetByID(c echo.Context) error {
	videoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	videos, err := controller.service.GetVideoByID(videoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetVideoByIDByPlaylistResponse(videos)
	responseData := common.DefaultDataResponse{
		Message: "Get video by id successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

// Create func for Create Video.
// @Description Create Video.
// @Summary Create Video
// @Tags Video
// @Accept json
// @Produce json
// @Param body body request.CreateVideoRequest true "Create Video"
// @Success 201 {object} common.DefaultDataResponse{data=response.GetVideoByIDByPlaylistResponse}
// @Security ApiKeyAuth
// @Router /v1/video [post]
func (controller *Controller) Create(c echo.Context) error {
	createVideoRequest := new(request.CreateVideoRequest)
	if err := c.Bind(&createVideoRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	request := *createVideoRequest.ToSpec()

	video, err := controller.service.CreateVideo(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetVideoByIDByPlaylistResponse(video)
	responseData := common.DefaultDataResponse{
		Message: "Create video successfully",
		Data:    response,
	}

	return c.JSON(http.StatusCreated, responseData)
}

// Update func for Update Video.
// @Description Update Video.
// @Summary Update Video
// @Tags Video
// @Accept json
// @Produce json
// @Param id path string true "Video ID"
// @Param body body request.CreateVideoRequest true "Video"
// @Success 200 {object} common.DefaultDataResponse{data=response.GetVideoByIDByPlaylistResponse}
// @Security ApiKeyAuth
// @Router /v1/video/{id} [put]
func (controller *Controller) Update(c echo.Context) error {
	videoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	updateVideoRequest := new(request.CreateVideoRequest)
	if err := c.Bind(&updateVideoRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	request := *updateVideoRequest.ToSpec()

	video, err := controller.service.UpdateVideo(&request, videoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetVideoByIDByPlaylistResponse(video)
	responseData := common.DefaultDataResponse{
		Message: "Update video successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

// Delete func for Delete Video.
// @Description Delete Video.
// @Summary Delete Video
// @Tags Video
// @Accept json
// @Produce json
// @Param id path string true "Video ID"
// @Success 200 {object} common.DefaultDataResponse{data=response.GetVideoByIDByPlaylistResponse}
// @Security ApiKeyAuth
// @Router /v1/video/{id} [delete]
func (controller *Controller) Delete(c echo.Context) error {
	videoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	video, err := controller.service.DeleteVideo(videoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetVideoByIDByPlaylistResponse(video)
	responseData := common.DefaultDataResponse{
		Message: "Delete video successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
