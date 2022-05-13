package response

import (
	"time"
	"vidlearn-final-projcect/business/playlist"
)

type GetPlaylistByIDResponse struct {
	ID       int `json:"id"`
	Category struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		ImageURL string `json:"image_url"`
	} `json:"category"`
	Name         string    `json:"name"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Description  string    `json:"description"`
	Level        string    `json:"level"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by"`
}

func CreateGetAllPlaylistResponse(playlists []*playlist.Playlist) []*GetPlaylistByIDResponse {
	var playlistsResponse []*GetPlaylistByIDResponse

	for _, playlist := range playlists {
		playlistsResponse = append(playlistsResponse, CreateGetPlaylistByIDResponse(playlist))
	}

	return playlistsResponse
}

func CreateGetPlaylistByIDResponse(playlist *playlist.Playlist) *GetPlaylistByIDResponse {
	return &GetPlaylistByIDResponse{
		ID: playlist.ID,
		Category: struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			ImageURL string `json:"image_url"`
		}{
			ID:       playlist.Category.ID,
			Name:     playlist.Category.Name,
			ImageURL: playlist.Category.ImageURL,
		},

		Name:         playlist.Name,
		ThumbnailURL: playlist.ThumbnailURL,
		Description:  playlist.Description,
		Level:        playlist.Level,
		CreatedAt:    playlist.CreatedAt,
		CreatedBy:    playlist.CreatedBy,
		UpdatedAt:    playlist.UpdatedAt,
		UpdatedBy:    playlist.UpdatedBy,
	}
}
