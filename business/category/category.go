package category

type Category struct {
	ID       int
	Name     string
	ImageURL string
}

func NewCategory(
	name string,
	imageURL string,
) *Category {
	return &Category{
		Name:     name,
		ImageURL: imageURL,
	}
}

func (oldCategory *Category) ModifyCategory(
	name string,
	imageURL string,
) *Category {
	return &Category{
		ID:       oldCategory.ID,
		Name:     name,
		ImageURL: imageURL,
	}
}
