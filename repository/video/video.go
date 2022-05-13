package video

import (
	"vidlearn-final-projcect/business/video"
	"vidlearn-final-projcect/util"

	"gorm.io/gorm"
)

func VideoRepository(dbCon *util.DatabaseConnection) video.Repository {
	var videoRepository video.Repository
	videoRepository = CreateMySQlRepository(dbCon.MySQlDB)

	return videoRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repostitory *MySQLRepository) GetVideoByPlaylist(playlistID int) (videos []*video.Video, err error) {
	err = repostitory.db.Preload("Playlist").Where("playlist_id=?", playlistID).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (repository *MySQLRepository) GetVideoByID(ID int) (video *video.Video, err error) {
	err = repository.db.Preload("Playlist").First(&video, ID).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (repository *MySQLRepository) CreateVideo(video *video.Video) (*video.Video, error) {
	err := repository.db.Create(&video).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (repository *MySQLRepository) UpdateVideo(videoCurrent *video.Video, IDCurrent int) (*video.Video, error) {
	err := repository.db.Model(&videoCurrent).Where("id=?", IDCurrent).Updates(videoCurrent).Error
	if err != nil {
		return nil, err
	}

	return videoCurrent, nil
}

func (repository *MySQLRepository) DeleteVideo(ID int) (video *video.Video, err error) {
	err = repository.db.First(&video, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&video).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}
