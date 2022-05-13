package category

import (
	"vidlearn-final-projcect/business/category"
	"vidlearn-final-projcect/util"

	"gorm.io/gorm"
)

func CategoryRepository(dbCon *util.DatabaseConnection) category.Repository {
	var categoryRepository category.Repository
	categoryRepository = CreateMySQlRepository(dbCon.MySQlDB)

	return categoryRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repository *MySQLRepository) GetAllCategory() (categories []*category.Category, err error) {
	err = repository.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (repository *MySQLRepository) GetCategoryByID(ID int) (category *category.Category, err error) {
	err = repository.db.First(&category, ID).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (repository *MySQLRepository) CreateCategory(category *category.Category) (*category.Category, error) {
	err := repository.db.Create(&category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (repository *MySQLRepository) UpdateCategory(categoryCurrent *category.Category, IDCurrent int) (*category.Category, error) {
	err := repository.db.Model(&categoryCurrent).Where("id=?", IDCurrent).Updates(categoryCurrent).Error
	if err != nil {
		return nil, err
	}

	return categoryCurrent, nil
}

func (repository *MySQLRepository) DeleteCategory(ID int) (category *category.Category, err error) {
	err = repository.db.First(&category, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}
