package playlist

import (
	"vidlearn-final-projcect/business/playlist"
	"vidlearn-final-projcect/business/video"
	"vidlearn-final-projcect/util"

	"gorm.io/gorm"
)

func PlaylistRepository(dbCon *util.DatabaseConnection) playlist.Repository {
	var playlistRepository playlist.Repository
	playlistRepository = CreateMySQlRepository(dbCon.MySQlDB)

	return playlistRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repository *MySQLRepository) GetAllPlaylist() (playlists []*playlist.Playlist, err error) {
	err = repository.db.Preload("Category").Find(&playlists).Error
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

func (repository *MySQLRepository) GetPlaylistByID(ID int) (playlist *playlist.Playlist, err error) {
	err = repository.db.Preload("Category").First(&playlist, ID).Error
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (repository *MySQLRepository) CreatePlaylist(playlist *playlist.Playlist) (*playlist.Playlist, error) {
	err := repository.db.Create(&playlist).Error
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (repository *MySQLRepository) UpdatePlaylist(playlistCurrent *playlist.Playlist, IDCurrent int) (*playlist.Playlist, error) {
	err := repository.db.Model(&playlistCurrent).Where("id=?", IDCurrent).Updates(playlistCurrent).Error
	if err != nil {
		return nil, err
	}

	return playlistCurrent, nil
}

func (repository *MySQLRepository) DeletePlaylist(ID int) (playlist *playlist.Playlist, err error) {
	err = repository.db.First(&playlist, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Where("playlist_id=?", ID).Delete(&video.Video{}).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&playlist).Error
	if err != nil {
		return nil, err
	}

	return playlist, nil
}
