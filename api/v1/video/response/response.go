package response

import (
	"time"
	"vidlearn-final-projcect/business/video"
)

type GetVideoByIDByPlaylistResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	VideoURL    string    `json:"video_url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Playlist struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnail_url"`
	Level        string `json:"level"`
	Videos       []*GetVideoByIDByPlaylistResponse
}

func CreateGetVideoByPlaylistResponse(videos []*video.Video) *Playlist {
	var videosResponse []*GetVideoByIDByPlaylistResponse

	for _, video := range videos {
		videosResponse = append(videosResponse, CreateGetVideoByIDByPlaylistResponse(video))
	}

	return &Playlist{
		ID:           videos[0].Playlist.ID,
		Name:         videos[0].Playlist.Name,
		ThumbnailURL: videos[0].Playlist.ThumbnailURL,
		Level:        videos[0].Playlist.Level,
		Videos:       videosResponse,
	}
}

func CreateGetVideoByIDByPlaylistResponse(video *video.Video) *GetVideoByIDByPlaylistResponse {
	return &GetVideoByIDByPlaylistResponse{
		ID:          video.ID,
		Title:       video.Title,
		VideoURL:    video.VideoURL,
		Description: video.Description,
		CreatedAt:   video.CreatedAt,
		UpdatedAt:   video.UpdatedAt,
	}
}
