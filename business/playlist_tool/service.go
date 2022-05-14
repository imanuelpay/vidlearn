package playlist_tool

import (
	"vidlearn-final-projcect/business/playlist_tool/spec"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetToolByPlaylist(playlistID int) (playlist_tools []*PlaylistTool, err error)
	GetPlaylistByTool(toolID int) (playlist_tools []*PlaylistTool, err error)
	CreatePlaylistTool(playlist_tools *PlaylistTool) (*PlaylistTool, error)
	DeletePlaylistTool(playlistID, toolID int) (playlist_tools *PlaylistTool, err error)
}

type Service interface {
	GetToolByPlaylist(playlistID int) (playlist_tools []*PlaylistTool, err error)
	GetPlaylistByTool(toolID int) (playlist_tools []*PlaylistTool, err error)
	CreatePlaylistTool(upsertPlaylistToolSpec *spec.UpsertPlaylistToolSpec, PlaylistID int) (*PlaylistTool, error)
	DeletePlaylistTool(playlistID, toolID int) (playlist_tools *PlaylistTool, err error)
}

type playlistToolService struct {
	repository Repository
	validate   *validator.Validate
}

func CreateService(repository Repository) Service {
	return &playlistToolService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (service *playlistToolService) GetToolByPlaylist(playlistID int) (playlist_tools []*PlaylistTool, err error) {
	return service.repository.GetToolByPlaylist(playlistID)
}

func (service *playlistToolService) GetPlaylistByTool(toolID int) (playlist_tools []*PlaylistTool, err error) {
	return service.repository.GetPlaylistByTool(toolID)
}

func (service *playlistToolService) CreatePlaylistTool(upsertPlaylistToolSpec *spec.UpsertPlaylistToolSpec, PlaylistID int) (*PlaylistTool, error) {
	if err := service.validate.Struct(upsertPlaylistToolSpec); err != nil {
		return nil, err
	}

	playlist_tool := NewPlaylistTool(
		PlaylistID,
		upsertPlaylistToolSpec.ToolID,
	)

	return service.repository.CreatePlaylistTool(playlist_tool)
}

func (service *playlistToolService) DeletePlaylistTool(playlistID, toolID int) (playlist_tools *PlaylistTool, err error) {
	return service.repository.DeletePlaylistTool(playlistID, toolID)
}
