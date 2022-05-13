package playlist

import (
	"net/http"
	"strconv"
	"vidlearn-final-projcect/api/common"
	"vidlearn-final-projcect/api/v1/playlist/request"
	"vidlearn-final-projcect/api/v1/playlist/response"
	playlistBusiness "vidlearn-final-projcect/business/playlist"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service playlistBusiness.Service
}

func CreateController(service playlistBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	playlists, err := controller.service.GetAllPlaylist()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetAllPlaylistResponse(playlists)
	responseData := common.DefaultDataResponse{
		Message: "Get all playlists successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) GetByID(c echo.Context) error {
	playlistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	playlist, err := controller.service.GetPlaylistByID(playlistID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetPlaylistByIDResponse(playlist)
	responseData := common.DefaultDataResponse{
		Message: "Get playlist by id successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Create(c echo.Context) error {
	createPlaylistRequest := new(request.CreatePlaylistRequest)
	if err := c.Bind(createPlaylistRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	request := *createPlaylistRequest.ToSpec()

	// TODO: authorize admin
	creared_by := "admin" // static for now

	playlist, err := controller.service.CreatePlaylist(&request, creared_by)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetPlaylistByIDResponse(playlist)
	responseData := common.DefaultDataResponse{
		Message: "Create playlist successfully",
		Data:    response,
	}

	return c.JSON(http.StatusCreated, responseData)
}

func (controller *Controller) Update(c echo.Context) error {
	playlistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	updatePlaylistRequest := new(request.CreatePlaylistRequest)
	if err := c.Bind(updatePlaylistRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	request := *updatePlaylistRequest.ToSpec()

	// TODO: authorize admin
	updated_by := "admin" // static for now

	playlist, err := controller.service.UpdatePlaylist(&request, updated_by, playlistID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetPlaylistByIDResponse(playlist)
	responseData := common.DefaultDataResponse{
		Message: "Update playlist successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}

func (controller *Controller) Delete(c echo.Context) error {
	playlistID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	playlist, err := controller.service.DeletePlaylist(playlistID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.DefaultDataResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	response := response.CreateGetPlaylistByIDResponse(playlist)
	responseData := common.DefaultDataResponse{
		Message: "Delete playlist successfully",
		Data:    response,
	}

	return c.JSON(http.StatusOK, responseData)
}
