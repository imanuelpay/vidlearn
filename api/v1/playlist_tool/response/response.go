package response

import (
	"time"
	"vidlearn-final-projcect/business/playlist_tool"
)

type GetToolsByPlaylist struct {
	PlaylistID   int       `json:"playlist_id"`
	CategoryID   int       `json:"category_id"`
	Name         string    `json:"name"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Description  string    `json:"description"`
	Level        string    `json:"level"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by"`
	Tools        []*Tool   `json:"tools"`
}

type GetToolsByPlaylistID struct {
	PlaylistID   int       `json:"playlist_id"`
	CategoryID   int       `json:"category_id"`
	Name         string    `json:"name"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Description  string    `json:"description"`
	Level        string    `json:"level"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by"`
	Tool         *Tool     `json:"tool"`
}

type GetPlaylistsByTool struct {
	ToolID    int    `json:"tool_id"`
	Name      string `json:"name"`
	ToolURL   string `json:"tool_url"`
	ImageURL  string `json:"image_url"`
	Playlists []*Playlist
}

type Playlist struct {
	PlaylistID   int       `json:"playlist_id"`
	CategoryID   int       `json:"category_id"`
	Name         string    `json:"name"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Description  string    `json:"description"`
	Level        string    `json:"level"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by"`
}

type Tool struct {
	ToolID   int    `json:"tool_id"`
	Name     string `json:"name"`
	ToolURL  string `json:"tool_url"`
	ImageURL string `json:"image_url"`
}

func CreateGetToolsByPlaylistIDResponse(playlistTool *playlist_tool.PlaylistTool) *GetToolsByPlaylistID {
	return &GetToolsByPlaylistID{
		PlaylistID:   playlistTool.Playlist.ID,
		CategoryID:   playlistTool.Playlist.CategoryID,
		Name:         playlistTool.Playlist.Name,
		ThumbnailURL: playlistTool.Playlist.ThumbnailURL,
		Description:  playlistTool.Playlist.Description,
		Level:        playlistTool.Playlist.Level,
		CreatedAt:    playlistTool.Playlist.CreatedAt,
		CreatedBy:    playlistTool.Playlist.CreatedBy,
		UpdatedAt:    playlistTool.Playlist.UpdatedAt,
		UpdatedBy:    playlistTool.Playlist.UpdatedBy,
		Tool: &Tool{
			ToolID:   playlistTool.Tool.ID,
			Name:     playlistTool.Tool.Name,
			ToolURL:  playlistTool.Tool.ToolURL,
			ImageURL: playlistTool.Tool.ImageURL,
		},
	}
}

func CreateGetToolsByPlaylistResponse(playlistTool []*playlist_tool.PlaylistTool) *GetToolsByPlaylist {
	var getToolsByPlaylistResponse *GetToolsByPlaylist
	var tools []*Tool

	for _, playlistTool := range playlistTool {
		toolData := Tool{
			ToolID:   playlistTool.Tool.ID,
			Name:     playlistTool.Tool.Name,
			ToolURL:  playlistTool.Tool.ToolURL,
			ImageURL: playlistTool.Tool.ImageURL,
		}

		tools = append(tools, &toolData)
	}

	getToolsByPlaylistResponse = &GetToolsByPlaylist{
		PlaylistID:   playlistTool[0].Playlist.ID,
		CategoryID:   playlistTool[0].Playlist.CategoryID,
		Name:         playlistTool[0].Playlist.Name,
		ThumbnailURL: playlistTool[0].Playlist.ThumbnailURL,
		Description:  playlistTool[0].Playlist.Description,
		Level:        playlistTool[0].Playlist.Level,
		CreatedAt:    playlistTool[0].Playlist.CreatedAt,
		CreatedBy:    playlistTool[0].Playlist.CreatedBy,
		UpdatedAt:    playlistTool[0].Playlist.UpdatedAt,
		UpdatedBy:    playlistTool[0].Playlist.UpdatedBy,
		Tools:        tools,
	}

	return getToolsByPlaylistResponse
}

func CreateGetPlaylistsByToolResponse(playlistTool []*playlist_tool.PlaylistTool) *GetPlaylistsByTool {
	var getPlaylistsByToolResponse *GetPlaylistsByTool
	var playlists []*Playlist

	for _, playlistTool := range playlistTool {
		playlistData := Playlist{
			PlaylistID:   playlistTool.Playlist.ID,
			CategoryID:   playlistTool.Playlist.CategoryID,
			Name:         playlistTool.Playlist.Name,
			ThumbnailURL: playlistTool.Playlist.ThumbnailURL,
			Description:  playlistTool.Playlist.Description,
			Level:        playlistTool.Playlist.Level,
			CreatedAt:    playlistTool.Playlist.CreatedAt,
			CreatedBy:    playlistTool.Playlist.CreatedBy,
			UpdatedAt:    playlistTool.Playlist.UpdatedAt,
			UpdatedBy:    playlistTool.Playlist.UpdatedBy,
		}

		playlists = append(playlists, &playlistData)
	}

	getPlaylistsByToolResponse = &GetPlaylistsByTool{
		ToolID:    playlistTool[0].Tool.ID,
		Name:      playlistTool[0].Tool.Name,
		ToolURL:   playlistTool[0].Tool.ToolURL,
		ImageURL:  playlistTool[0].Tool.ImageURL,
		Playlists: playlists,
	}

	return getPlaylistsByToolResponse
}
