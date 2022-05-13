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

	response := response.CreateGetVideoByPlaylistResponse(videos)
	responseData := common.DefaultDataResponse{
		Message: "Get video by playlist successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

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

	responseData := common.DefaultDataResponse{
		Message: "Delete video successfully",
		Data:    video,
	}

	return c.JSON(http.StatusOK, responseData)
}
