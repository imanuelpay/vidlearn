package video

import (
	"time"
	"vidlearn-final-projcect/business/playlist"
)

type Video struct {
	ID          int
	PlaylistID  int
	Title       string
	VideoURL    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Playlist playlist.Playlist
}

func NewVideo(
	playlistID int,
	title string,
	videoURL string,
	description string,
) *Video {
	return &Video{
		PlaylistID:  playlistID,
		Title:       title,
		VideoURL:    videoURL,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

func (oldVideo *Video) ModifyVideo(
	playlistID int,
	title string,
	videoURL string,
	description string,
) *Video {
	return &Video{
		ID:          oldVideo.ID,
		PlaylistID:  playlistID,
		Title:       title,
		VideoURL:    videoURL,
		Description: description,
		CreatedAt:   oldVideo.CreatedAt,
		UpdatedAt:   time.Now(),
	}
}
