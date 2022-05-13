package request

import "vidlearn-final-projcect/business/playlist_tool/spec"

type CreatePlaylistToolRequest struct {
	PlaylistID int `json:"playlist_id"`
	ToolID     int `json:"tool_id"`
}

func (request *CreatePlaylistToolRequest) ToSpec() *spec.UpsertPlaylistToolSpec {
	return &spec.UpsertPlaylistToolSpec{
		PlaylistID: request.PlaylistID,
		ToolID:     request.ToolID,
	}
}
