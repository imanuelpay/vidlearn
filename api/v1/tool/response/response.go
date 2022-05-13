package response

import "vidlearn-final-projcect/business/tool"

type GetToolByIDResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ToolURL  string `json:"tool_url"`
	ImageURL string `json:"image_url"`
}

func CreateGetAllToolResponse(tools []*tool.Tool) []*GetToolByIDResponse {
	var toolsResponse []*GetToolByIDResponse

	for _, tool := range tools {
		toolsResponse = append(toolsResponse, CreateGetToolByIDResponse(tool))
	}

	return toolsResponse
}

func CreateGetToolByIDResponse(tool *tool.Tool) *GetToolByIDResponse {
	return &GetToolByIDResponse{
		ID:       tool.ID,
		Name:     tool.Name,
		ToolURL:  tool.ToolURL,
		ImageURL: tool.ImageURL,
	}
}
