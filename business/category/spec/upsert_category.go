package spec

type UpsertCategorySpec struct {
	Name     string `validate:"required"`
	ImageURL string `validate:"required"`
}
