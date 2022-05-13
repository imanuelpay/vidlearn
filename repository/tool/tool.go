package tool

import (
	"vidlearn-final-projcect/business/tool"
	"vidlearn-final-projcect/util"

	"gorm.io/gorm"
)

func ToolRepository(dbCon *util.DatabaseConnection) tool.Repository {
	var toolRepository tool.Repository
	toolRepository = CreateMySQlRepository(dbCon.MySQlDB)

	return toolRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repository *MySQLRepository) GetAllTool() (tools []*tool.Tool, err error) {
	err = repository.db.Find(&tools).Error
	if err != nil {
		return nil, err
	}

	return tools, nil
}

func (repository *MySQLRepository) GetToolByID(ID int) (tool *tool.Tool, err error) {
	err = repository.db.First(&tool, ID).Error
	if err != nil {
		return nil, err
	}

	return tool, nil
}

func (repository *MySQLRepository) CreateTool(tool *tool.Tool) (*tool.Tool, error) {
	err := repository.db.Create(&tool).Error
	if err != nil {
		return nil, err
	}

	return tool, nil
}

func (repository *MySQLRepository) UpdateTool(toolCurrent *tool.Tool, IDCurrent int) (*tool.Tool, error) {
	err := repository.db.Model(&toolCurrent).Where("id=?", IDCurrent).Updates(toolCurrent).Error
	if err != nil {
		return nil, err
	}

	return toolCurrent, nil
}

func (repository *MySQLRepository) DeleteTool(ID int) (tool *tool.Tool, err error) {
	err = repository.db.First(&tool, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&tool).Error
	if err != nil {
		return nil, err
	}

	return tool, nil
}
