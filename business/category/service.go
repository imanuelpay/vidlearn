package category

import (
	"vidlearn-final-projcect/business/category/spec"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAllCategory() (categories []*Category, err error)
	GetCategoryByID(ID int) (category *Category, err error)
	CreateCategory(category *Category) (*Category, error)
	UpdateCategory(categoryCurrent *Category, IDCurrent int) (*Category, error)
	DeleteCategory(ID int) (category *Category, err error)
}

type Service interface {
	GetAllCategory() (categories []*Category, err error)
	GetCategoryByID(ID int) (category *Category, err error)
	CreateCategory(upsertCategorySpec *spec.UpsertCategorySpec) (*Category, error)
	UpdateCategory(upsertCategorySpec *spec.UpsertCategorySpec, IDCurrent int) (*Category, error)
	DeleteCategory(ID int) (category *Category, err error)
}

type categoryService struct {
	repository Repository
	validate   *validator.Validate
}

func CreateService(repository Repository) Service {
	return &categoryService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (service *categoryService) GetAllCategory() (categories []*Category, err error) {
	return service.repository.GetAllCategory()
}

func (service *categoryService) GetCategoryByID(ID int) (category *Category, err error) {
	return service.repository.GetCategoryByID(ID)
}

func (service *categoryService) CreateCategory(upsertCategorySpec *spec.UpsertCategorySpec) (*Category, error) {
	if err := service.validate.Struct(upsertCategorySpec); err != nil {
		return nil, err
	}

	category := NewCategory(upsertCategorySpec.Name, upsertCategorySpec.ImageURL)
	return service.repository.CreateCategory(category)
}

func (service *categoryService) UpdateCategory(upsertCategorySpec *spec.UpsertCategorySpec, IDCurrent int) (*Category, error) {
	if err := service.validate.Struct(upsertCategorySpec); err != nil {
		return nil, err
	}

	category, err := service.repository.GetCategoryByID(IDCurrent)
	if err != nil {
		return nil, err
	}

	category = category.ModifyCategory(upsertCategorySpec.Name, upsertCategorySpec.ImageURL)
	return service.repository.UpdateCategory(category, IDCurrent)
}

func (service *categoryService) DeleteCategory(ID int) (category *Category, err error) {
	return service.repository.DeleteCategory(ID)
}
