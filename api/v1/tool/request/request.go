package request

import "vidlearn-final-projcect/business/tool/spec"

type CreateToolRequest struct {
	Name     string `json:"name"`
	ToolURL  string `json:"tool_url"`
	ImageURL string `json:"image_url"`
}

func (request *CreateToolRequest) ToSpec() *spec.UpsertToolSpec {
	return &spec.UpsertToolSpec{
		Name:     request.Name,
		ToolURL:  request.ToolURL,
		ImageURL: request.ImageURL,
	}
}
