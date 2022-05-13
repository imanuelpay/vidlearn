package tool

import (
	"vidlearn-final-projcect/business/tool/spec"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAllTool() (tools []*Tool, err error)
	GetToolByID(ID int) (tool *Tool, err error)
	CreateTool(tool *Tool) (*Tool, error)
	UpdateTool(toolCurrent *Tool, IDCurrent int) (*Tool, error)
	DeleteTool(ID int) (tool *Tool, err error)
}

type Service interface {
	GetAllTool() (tools []*Tool, err error)
	GetToolByID(ID int) (tool *Tool, err error)
	CreateTool(upsertToolSpec *spec.UpsertToolSpec) (*Tool, error)
	UpdateTool(upsertToolSpec *spec.UpsertToolSpec, IDCurrent int) (*Tool, error)
	DeleteTool(ID int) (tool *Tool, err error)
}

type toolService struct {
	repository Repository
	validate   *validator.Validate
}

func CreateService(repository Repository) Service {
	return &toolService{
		repository: repository,
		validate:   validator.New(),
	}
}

func (service *toolService) GetAllTool() (tools []*Tool, err error) {
	return service.repository.GetAllTool()
}

func (service *toolService) GetToolByID(ID int) (tool *Tool, err error) {
	return service.repository.GetToolByID(ID)
}

func (service *toolService) CreateTool(upsertToolSpec *spec.UpsertToolSpec) (*Tool, error) {
	if err := service.validate.Struct(upsertToolSpec); err != nil {
		return nil, err
	}

	tool := NewTool(upsertToolSpec.Name, upsertToolSpec.ToolURL, upsertToolSpec.ImageURL)
	return service.repository.CreateTool(tool)
}

func (service *toolService) UpdateTool(upsertToolSpec *spec.UpsertToolSpec, IDCurrent int) (*Tool, error) {
	if err := service.validate.Struct(upsertToolSpec); err != nil {
		return nil, err
	}

	tool, err := service.repository.GetToolByID(IDCurrent)
	if err != nil {
		return nil, err
	}

	tool = tool.ModifyTool(upsertToolSpec.Name, upsertToolSpec.ToolURL, upsertToolSpec.ImageURL)
	return service.repository.UpdateTool(tool, IDCurrent)
}

func (service *toolService) DeleteTool(ID int) (tool *Tool, err error) {
	return service.repository.DeleteTool(ID)
}
