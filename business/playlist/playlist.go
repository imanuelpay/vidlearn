package playlist

import (
	"time"
	"vidlearn-final-projcect/business/category"
)

type Playlist struct {
	ID           int
	CategoryID   int
	Name         string
	ThumbnailURL string
	Description  string
	Level        string
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string

	Category category.Category
}

func NewPlaylist(
	categoryID int,
	name string,
	thumbnailURL string,
	description string,
	level string,
	createdBy string,
) *Playlist {
	return &Playlist{
		CategoryID:   categoryID,
		Name:         name,
		ThumbnailURL: thumbnailURL,
		Description:  description,
		Level:        level,
		CreatedAt:    time.Now(),
		CreatedBy:    createdBy,
	}
}

func (oldPlaylist *Playlist) ModifyPlaylist(
	categoryID int,
	name string,
	thumbnailURL string,
	description string,
	level string,
	updatedBy string,
) *Playlist {
	return &Playlist{
		ID:           oldPlaylist.ID,
		CategoryID:   categoryID,
		Name:         name,
		ThumbnailURL: thumbnailURL,
		Description:  description,
		Level:        level,
		CreatedAt:    oldPlaylist.CreatedAt,
		CreatedBy:    oldPlaylist.CreatedBy,
		UpdatedAt:    time.Now(),
		UpdatedBy:    updatedBy,
	}
}
