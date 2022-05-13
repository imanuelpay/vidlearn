package playlist

import (
	"vidlearn-final-projcect/business/playlist/spec"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAllPlaylist() (playlists []*Playlist, err error)
	GetPlaylistByID(ID int) (playlist *Playlist, err error)
	CreatePlaylist(playlist *Playlist) (*Playlist, error)
	UpdatePlaylist(playlistCurrent *Playlist, IDCurrent int) (*Playlist, error)
	DeletePlaylist(ID int) (playlist *Playlist, err error)
}

type Service interface {
	GetAllPlaylist() (playlists []*Playlist, err error)
	GetPlaylistByID(ID int) (playlist *Playlist, err error)
	CreatePlaylist(upsertPlaylistSpec *spec.UpsertPlaylistSpec, createdBy string) (*Playlist, error)
	UpdatePlaylist(upsertPlaylistSpec *spec.UpsertPlaylistSpec, updatedBy string, IDCurrent int) (*Playlist, error)
	DeletePlaylist(ID int) (playlist *Playlist, err error)
}

type playlistService struct {
	repository Repository
	validate   *validator.Validate
}

func CreateService(repository Repository) Service {
	return &playlistService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (service *playlistService) GetAllPlaylist() (playlists []*Playlist, err error) {
	return service.repository.GetAllPlaylist()
}

func (service *playlistService) GetPlaylistByID(ID int) (playlist *Playlist, err error) {
	return service.repository.GetPlaylistByID(ID)
}

func (service *playlistService) CreatePlaylist(upsertPlaylistSpec *spec.UpsertPlaylistSpec, createdBy string) (*Playlist, error) {
	if err := service.validate.Struct(upsertPlaylistSpec); err != nil {
		return nil, err
	}

	playlist := NewPlaylist(
		upsertPlaylistSpec.CategoryID,
		upsertPlaylistSpec.Name,
		upsertPlaylistSpec.ThumbnailURL,
		upsertPlaylistSpec.Description,
		upsertPlaylistSpec.Level,
		createdBy,
	)

	return service.repository.CreatePlaylist(playlist)
}

func (service *playlistService) UpdatePlaylist(upsertPlaylistSpec *spec.UpsertPlaylistSpec, updatedBy string, IDCurrent int) (*Playlist, error) {
	if err := service.validate.Struct(upsertPlaylistSpec); err != nil {
		return nil, err
	}

	playlistCurrent, err := service.repository.GetPlaylistByID(IDCurrent)
	if err != nil {
		return nil, err
	}

	playlist := playlistCurrent.ModifyPlaylist(
		upsertPlaylistSpec.CategoryID,
		upsertPlaylistSpec.Name,
		upsertPlaylistSpec.ThumbnailURL,
		upsertPlaylistSpec.Description,
		upsertPlaylistSpec.Level,
		updatedBy,
	)

	return service.repository.UpdatePlaylist(playlist, IDCurrent)
}

func (service *playlistService) DeletePlaylist(ID int) (playlist *Playlist, err error) {
	return service.repository.DeletePlaylist(ID)
}
