package request

import "vidlearn-final-projcect/business/video/spec"

type CreateVideoRequest struct {
	PlaylistID  int    `json:"playlist_id"`
	Title       string `json:"title"`
	VideoURL    string `json:"video_url"`
	Description string `json:"description"`
}

func (request *CreateVideoRequest) ToSpec() *spec.UpsertVideoSpec {
	return &spec.UpsertVideoSpec{
		PlaylistID:  request.PlaylistID,
		Title:       request.Title,
		VideoURL:    request.VideoURL,
		Description: request.Description,
	}
}
