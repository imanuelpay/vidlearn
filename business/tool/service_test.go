package tool_test

import (
	"errors"
	"testing"
	"vidlearn-final-projcect/business/tool"
	toolMocks "vidlearn-final-projcect/business/tool/mocks"
	"vidlearn-final-projcect/business/tool/spec"

	"github.com/stretchr/testify/mock"
)

var (
	mockToolRespository = new(toolMocks.Repository)
)

func TestGetAllTool(t *testing.T) {
	toolData := []*tool.Tool{}

	t.Run("success", func(t *testing.T) {
		mockToolRespository.On("GetAllTool", mock.Anything).Return(toolData, nil).Once()

		toolService := tool.CreateService(mockToolRespository)
		tool, err := toolService.GetAllTool()

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(tool) != len(toolData) {
			t.Errorf("Expected length data %v, got %v", len(tool), len(toolData))
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockToolRespository.On("GetAllTool", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		toolService := tool.CreateService(mockToolRespository)
		_, err := toolService.GetAllTool()

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestGetToolByID(t *testing.T) {
	toolData := &tool.Tool{
		ID:       1,
		Name:     "Test Tool",
		ToolURL:  "https://tool.test/download",
		ImageURL: "https://img.test/test.png",
	}

	t.Run("success", func(t *testing.T) {
		mockToolRespository.On("GetToolByID", mock.Anything).Return(toolData, nil).Once()

		toolService := tool.CreateService(mockToolRespository)
		tool, err := toolService.GetToolByID(toolData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if tool.ID != toolData.ID {
			t.Errorf("Expected tool ID %v, got %v", toolData.ID, tool.ID)
		}

		if tool.Name != toolData.Name {
			t.Errorf("Expected tool Name %s, got %s", toolData.Name, tool.Name)
		}

		if tool.ToolURL != toolData.ToolURL {
			t.Errorf("Expected tool URL %s, got %s", toolData.ToolURL, tool.ToolURL)
		}

		if tool.ImageURL != toolData.ImageURL {
			t.Errorf("Expected tool image URL %s, got %s", toolData.ImageURL, tool.ImageURL)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockToolRespository.On("GetToolByID", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		toolService := tool.CreateService(mockToolRespository)
		_, err := toolService.GetToolByID(toolData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestCreateTool(t *testing.T) {
	upsertToolData := &spec.UpsertToolSpec{
		Name:     "Test Tool",
		ToolURL:  "https://tool.test/download",
		ImageURL: "https://img.test/test.png",
	}

	toolData := &tool.Tool{}

	t.Run("success", func(t *testing.T) {
		mockToolRespository.On("CreateTool", mock.Anything).Return(toolData, nil).Once()

		toolService := tool.CreateService(mockToolRespository)
		tool, err := toolService.CreateTool(upsertToolData)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if tool.ID != toolData.ID {
			t.Errorf("Expected tool ID %v, got %v", toolData.ID, tool.ID)
		}

		if tool.Name != toolData.Name {
			t.Errorf("Expected tool Name %s, got %s", toolData.Name, tool.Name)
		}

		if tool.ToolURL != toolData.ToolURL {
			t.Errorf("Expected tool URL %s, got %s", toolData.ToolURL, tool.ToolURL)
		}

		if tool.ImageURL != toolData.ImageURL {
			t.Errorf("Expected tool image URL %s, got %s", toolData.ImageURL, tool.ImageURL)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockToolRespository.On("CreateTool", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		toolService := tool.CreateService(mockToolRespository)
		_, err := toolService.CreateTool(upsertToolData)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestUpdateTool(t *testing.T) {
	toolData := &tool.Tool{
		ID: 1,
	}

	upsertToolData := &spec.UpsertToolSpec{
		Name:     "Test Tool Updated",
		ToolURL:  "https://tool.test/download-updated",
		ImageURL: "https://img.test/test-updated.png",
	}

	t.Run("success", func(t *testing.T) {
		mockToolRespository.On("GetToolByID", toolData.ID).Return(toolData, nil).Once()
		mockToolRespository.On("UpdateTool", mock.Anything, toolData.ID).Return(toolData, nil).Once()

		toolService := tool.CreateService(mockToolRespository)
		tool, err := toolService.UpdateTool(upsertToolData, toolData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if tool.ID != toolData.ID {
			t.Errorf("Expected tool ID %v, got %v", toolData.ID, tool.ID)
		}

		if tool.Name != toolData.Name {
			t.Errorf("Expected tool Name %s, got %s", toolData.Name, tool.Name)
		}

		if tool.ToolURL != toolData.ToolURL {
			t.Errorf("Expected tool URL %s, got %s", toolData.ToolURL, tool.ToolURL)
		}

		if tool.ImageURL != toolData.ImageURL {
			t.Errorf("Expected tool image URL %s, got %s", toolData.ImageURL, tool.ImageURL)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockToolRespository.On("GetToolByID", toolData.ID).Return(nil, errors.New("error unit testing")).Once()
		mockToolRespository.On("UpdateTool", mock.Anything, toolData.ID).Return(nil, errors.New("error unit testing")).Once()

		toolService := tool.CreateService(mockToolRespository)
		_, err := toolService.UpdateTool(upsertToolData, toolData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestDeleteTool(t *testing.T) {
	toolData := &tool.Tool{
		ID:       1,
		Name:     "Test Tool",
		ToolURL:  "https://tool.test/download",
		ImageURL: "https://img.test/test.png",
	}

	t.Run("success", func(t *testing.T) {
		mockToolRespository.On("DeleteTool", mock.Anything).Return(toolData, nil).Once()

		toolService := tool.CreateService(mockToolRespository)
		tool, err := toolService.DeleteTool(toolData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if tool.ID != toolData.ID {
			t.Errorf("Expected tool ID %v, got %v", toolData.ID, tool.ID)
		}

		if tool.Name != toolData.Name {
			t.Errorf("Expected tool Name %s, got %s", toolData.Name, tool.Name)
		}

		if tool.ToolURL != toolData.ToolURL {
			t.Errorf("Expected tool URL %s, got %s", toolData.ToolURL, tool.ToolURL)
		}

		if tool.ImageURL != toolData.ImageURL {
			t.Errorf("Expected tool image URL %s, got %s", toolData.ImageURL, tool.ImageURL)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockToolRespository.On("DeleteTool", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		toolService := tool.CreateService(mockToolRespository)
		_, err := toolService.DeleteTool(toolData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
