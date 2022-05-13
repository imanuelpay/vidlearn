package request

import "vidlearn-final-projcect/business/playlist/spec"

type CreatePlaylistRequest struct {
	Name         string `json:"name"`
	CategoryID   int    `json:"category_id"`
	ThumbnailURL string `json:"thumbnail_url"`
	Description  string `json:"description"`
	Level        string `json:"level"`
}

func (request *CreatePlaylistRequest) ToSpec() *spec.UpsertPlaylistSpec {
	return &spec.UpsertPlaylistSpec{
		Name:         request.Name,
		CategoryID:   request.CategoryID,
		ThumbnailURL: request.ThumbnailURL,
		Description:  request.Description,
		Level:        request.Level,
	}
}
