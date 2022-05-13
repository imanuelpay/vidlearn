package tool

type Tool struct {
	ID       int
	Name     string
	ToolURL  string
	ImageURL string
}

func NewTool(
	name string,
	toolURL string,
	imageURL string,
) *Tool {
	return &Tool{
		Name:     name,
		ToolURL:  toolURL,
		ImageURL: imageURL,
	}
}

func (oldTool *Tool) ModifyTool(
	name string,
	toolURL string,
	imageURL string,
) *Tool {
	return &Tool{
		ID:       oldTool.ID,
		Name:     name,
		ToolURL:  toolURL,
		ImageURL: imageURL,
	}
}
