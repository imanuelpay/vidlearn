package category_test

import (
	"errors"
	"testing"
	"vidlearn-final-projcect/business/category"
	categoryMocks "vidlearn-final-projcect/business/category/mocks"
	"vidlearn-final-projcect/business/category/spec"

	"github.com/stretchr/testify/mock"
)

var (
	mockCategoryRespository = new(categoryMocks.Repository)
)

func TestGetAllCategory(t *testing.T) {
	categoryData := []*category.Category{}

	t.Run("success", func(t *testing.T) {
		mockCategoryRespository.On("GetAllCategory", mock.Anything).Return(categoryData, nil).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		category, err := categoryService.GetAllCategory()

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(category) != len(categoryData) {
			t.Errorf("Expected lenght data %v, got %v", len(category), len(categoryData))
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockCategoryRespository.On("GetAllCategory", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		_, err := categoryService.GetAllCategory()

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestGetCategoryByID(t *testing.T) {
	categoryData := &category.Category{
		ID:       1,
		Name:     "Test Category",
		ImageURL: "https://img.test/test.png",
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRespository.On("GetCategoryByID", mock.Anything).Return(categoryData, nil).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		category, err := categoryService.GetCategoryByID(categoryData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if category.ID != categoryData.ID {
			t.Errorf("Expected category ID %v, got %v", categoryData.ID, category.ID)
		}

		if category.Name != categoryData.Name {
			t.Errorf("Expected category Name %s, got %s", categoryData.Name, category.Name)
		}

		if category.ImageURL != categoryData.ImageURL {
			t.Errorf("Expected category image URL %s, got %s", categoryData.ImageURL, category.ImageURL)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockCategoryRespository.On("GetCategoryByID", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		_, err := categoryService.GetCategoryByID(categoryData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestCreateCategory(t *testing.T) {
	upsertCategoryData := &spec.UpsertCategorySpec{
		Name:     "Test Category",
		ImageURL: "https://img.test/test.png",
	}

	categoryData := &category.Category{}

	t.Run("success", func(t *testing.T) {
		mockCategoryRespository.On("CreateCategory", mock.Anything).Return(categoryData, nil).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		category, err := categoryService.CreateCategory(upsertCategoryData)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if category.ID != categoryData.ID {
			t.Errorf("Expected category ID %v, got %v", categoryData.ID, category.ID)
		}

		if category.Name != categoryData.Name {
			t.Errorf("Expected category Name %s, got %s", categoryData.Name, category.Name)
		}

		if category.ImageURL != categoryData.ImageURL {
			t.Errorf("Expected category image URL %s, got %s", categoryData.ImageURL, category.ImageURL)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockCategoryRespository.On("CreateCategory", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		_, err := categoryService.CreateCategory(upsertCategoryData)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestUpdateCategory(t *testing.T) {
	categoryData := &category.Category{
		ID: 1,
	}

	upsertCategoryData := &spec.UpsertCategorySpec{
		Name:     "Test Category Update",
		ImageURL: "https://img.test/test-update.png",
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRespository.On("GetCategoryByID", categoryData.ID).Return(categoryData, nil).Once()
		mockCategoryRespository.On("UpdateCategory", mock.Anything, categoryData.ID).Return(categoryData, nil).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		category, err := categoryService.UpdateCategory(upsertCategoryData, categoryData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if category.ID != categoryData.ID {
			t.Errorf("Expected category ID %v, got %v", categoryData.ID, category.ID)
		}

		if category.Name != categoryData.Name {
			t.Errorf("Expected category Name %s, got %s", categoryData.Name, category.Name)
		}

		if category.ImageURL != categoryData.ImageURL {
			t.Errorf("Expected category image URL %s, got %s", categoryData.ImageURL, category.ImageURL)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockCategoryRespository.On("GetCategoryByID", categoryData.ID).Return(nil, errors.New("error unit testing")).Once()
		mockCategoryRespository.On("UpdateCategory", mock.Anything, categoryData.ID).Return(nil, errors.New("error unit testing")).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		_, err := categoryService.UpdateCategory(upsertCategoryData, categoryData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestDeleteCategory(t *testing.T) {
	categoryData := &category.Category{
		ID:       1,
		Name:     "Test Category",
		ImageURL: "https://img.test/test.png",
	}

	t.Run("success", func(t *testing.T) {
		mockCategoryRespository.On("DeleteCategory", mock.Anything).Return(categoryData, nil).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		category, err := categoryService.DeleteCategory(categoryData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if category.ID != categoryData.ID {
			t.Errorf("Expected category ID %v, got %v", categoryData.ID, category.ID)
		}

		if category.Name != categoryData.Name {
			t.Errorf("Expected category Name %s, got %s", categoryData.Name, category.Name)
		}

		if category.ImageURL != categoryData.ImageURL {
			t.Errorf("Expected category image URL %s, got %s", categoryData.ImageURL, category.ImageURL)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockCategoryRespository.On("DeleteCategory", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		categoryService := category.CreateService(mockCategoryRespository)
		_, err := categoryService.DeleteCategory(categoryData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
