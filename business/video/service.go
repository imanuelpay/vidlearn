package video

import (
	"vidlearn-final-projcect/business/video/spec"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetVideoByPlaylist(playlistID int) (videos []*Video, err error)
	GetVideoByID(ID int) (video *Video, err error)
	CreateVideo(video *Video) (*Video, error)
	UpdateVideo(videoCurrent *Video, IDCurrent int) (*Video, error)
	DeleteVideo(ID int) (video *Video, err error)
}

type Service interface {
	GetVideoByPlaylist(playlistID int) (videos []*Video, err error)
	GetVideoByID(ID int) (video *Video, err error)
	CreateVideo(upsertVideoSpec *spec.UpsertVideoSpec) (*Video, error)
	UpdateVideo(upsertVideoSpec *spec.UpsertVideoSpec, IDCurrent int) (*Video, error)
	DeleteVideo(ID int) (video *Video, err error)
}

type videoService struct {
	repository Repository
	validate   *validator.Validate
}

func CreateService(repository Repository) Service {
	return &videoService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (service *videoService) GetVideoByPlaylist(playlistID int) (videos []*Video, err error) {
	return service.repository.GetVideoByPlaylist(playlistID)
}

func (service *videoService) GetVideoByID(ID int) (video *Video, err error) {
	return service.repository.GetVideoByID(ID)
}

func (service *videoService) CreateVideo(upsertVideoSpec *spec.UpsertVideoSpec) (*Video, error) {
	if err := service.validate.Struct(upsertVideoSpec); err != nil {
		return nil, err
	}

	video := NewVideo(
		upsertVideoSpec.PlaylistID,
		upsertVideoSpec.Title,
		upsertVideoSpec.VideoURL,
		upsertVideoSpec.Description,
	)

	return service.repository.CreateVideo(video)
}

func (service *videoService) UpdateVideo(upsertVideoSpec *spec.UpsertVideoSpec, IDCurrent int) (*Video, error) {
	if err := service.validate.Struct(upsertVideoSpec); err != nil {
		return nil, err
	}

	videoCurrent, err := service.repository.GetVideoByID(IDCurrent)
	if err != nil {
		return nil, err
	}

	video := videoCurrent.ModifyVideo(
		upsertVideoSpec.PlaylistID,
		upsertVideoSpec.Title,
		upsertVideoSpec.VideoURL,
		upsertVideoSpec.Description,
	)

	return service.repository.UpdateVideo(video, IDCurrent)
}

func (service *videoService) DeleteVideo(ID int) (video *Video, err error) {
	return service.repository.DeleteVideo(ID)
}
