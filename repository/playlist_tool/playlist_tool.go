package playlist_tool

import (
	"vidlearn-final-projcect/business/playlist_tool"
	"vidlearn-final-projcect/util"

	"gorm.io/gorm"
)

func PlaylistToolRepository(dbCon *util.DatabaseConnection) playlist_tool.Repository {
	var playlistToolRepository playlist_tool.Repository
	playlistToolRepository = CreateMySQlRepository(dbCon.MySQlDB)

	return playlistToolRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repository *MySQLRepository) GetToolByPlaylist(playlistID int) (playlist_tools []*playlist_tool.PlaylistTool, err error) {
	err = repository.db.Preload("Tool").Preload("Playlist").Where("playlist_id=?", playlistID).Find(&playlist_tools).Error
	if err != nil {
		return nil, err
	}

	return playlist_tools, nil
}

func (repository *MySQLRepository) GetPlaylistByTool(playlistToolID int) (playlistTools []*playlist_tool.PlaylistTool, err error) {
	err = repository.db.Preload("Playlist").Preload("Tool").Where("tool_id=?", playlistToolID).Find(&playlistTools).Error
	if err != nil {
		return nil, err
	}

	return playlistTools, nil
}

func (repository *MySQLRepository) CreatePlaylistTool(playlistTool *playlist_tool.PlaylistTool) (*playlist_tool.PlaylistTool, error) {
	err := repository.db.Create(&playlistTool).Error
	if err != nil {
		return nil, err
	}

	return playlistTool, nil
}

func (repository *MySQLRepository) DeletePlaylistTool(ID int) (playlistTool *playlist_tool.PlaylistTool, err error) {
	err = repository.db.First(&playlistTool, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&playlistTool).Error
	if err != nil {
		return nil, err
	}

	return playlistTool, nil
}
