package playlist_tool

import (
	"vidlearn-final-projcect/business/playlist"
	"vidlearn-final-projcect/business/tool"
)

type PlaylistTool struct {
	ID         int
	PlaylistID int
	ToolID     int

	Playlist playlist.Playlist
	Tool     tool.Tool
}

func NewPlaylistTool(playlistID int, toolID int) *PlaylistTool {
	return &PlaylistTool{
		PlaylistID: playlistID,
		ToolID:     toolID,
	}
}
