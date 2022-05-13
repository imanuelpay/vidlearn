package spec

type UpsertToolSpec struct {
	Name     string `validate:"required"`
	ToolURL  string `validate:"required"`
	ImageURL string `validate:"required"`
}
