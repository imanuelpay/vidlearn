package request

import "vidlearn-final-projcect/business/playlist_tool/spec"

type CreatePlaylistToolRequest struct {
	ToolID int `json:"tool_id"`
}

func (request *CreatePlaylistToolRequest) ToSpec() *spec.UpsertPlaylistToolSpec {
	return &spec.UpsertPlaylistToolSpec{
		ToolID: request.ToolID,
	}
}
