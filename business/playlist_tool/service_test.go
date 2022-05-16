package playlist_tool_test

import (
	"errors"
	"testing"
	"vidlearn-final-projcect/business/playlist_tool"
	playlistToolMocks "vidlearn-final-projcect/business/playlist_tool/mocks"
	"vidlearn-final-projcect/business/playlist_tool/spec"

	"github.com/stretchr/testify/mock"
)

var (
	mockPlaylistToolRepository = new(playlistToolMocks.Repository)
)

func TestGetToolByPlaylist(t *testing.T) {
	playlistID := 1
	playlistToolData := []*playlist_tool.PlaylistTool{
		{
			ID:         1,
			PlaylistID: playlistID,
			ToolID:     1,
		},
		{
			ID:         2,
			PlaylistID: playlistID,
			ToolID:     2,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockPlaylistToolRepository.On("GetToolByPlaylist", mock.Anything).Return(playlistToolData, nil).Once()

		playlistToolService := playlist_tool.CreateService(mockPlaylistToolRepository)
		playlistTool, err := playlistToolService.GetToolByPlaylist(playlistID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(playlistTool) != len(playlistToolData) {
			t.Errorf("Expected lenght data %v, got %v", len(playlistTool), len(playlistToolData))
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockPlaylistToolRepository.On("GetToolByPlaylist", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		playlistToolService := playlist_tool.CreateService(mockPlaylistToolRepository)
		_, err := playlistToolService.GetToolByPlaylist(playlistID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestGetPlaylistByTool(t *testing.T) {
	toolID := 1
	playlistToolData := []*playlist_tool.PlaylistTool{
		{
			ID:         1,
			PlaylistID: 1,
			ToolID:     toolID,
		},
		{
			ID:         2,
			PlaylistID: 2,
			ToolID:     toolID,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockPlaylistToolRepository.On("GetPlaylistByTool", mock.Anything).Return(playlistToolData, nil).Once()

		playlistToolService := playlist_tool.CreateService(mockPlaylistToolRepository)
		playlistTool, err := playlistToolService.GetPlaylistByTool(toolID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(playlistTool) != len(playlistToolData) {
			t.Errorf("Expected lenght data %v, got %v", len(playlistTool), len(playlistToolData))
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockPlaylistToolRepository.On("GetPlaylistByTool", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		playlistToolService := playlist_tool.CreateService(mockPlaylistToolRepository)
		_, err := playlistToolService.GetPlaylistByTool(toolID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestCreatePlaylistTool(t *testing.T) {
	playlistID := 1
	upsertPlaylistToolData := &spec.UpsertPlaylistToolSpec{
		ToolID: 1,
	}

	playlistToolData := &playlist_tool.PlaylistTool{}

	t.Run("success", func(t *testing.T) {
		mockPlaylistToolRepository.On("CreatePlaylistTool", mock.Anything).Return(playlistToolData, nil).Once()

		playlistToolService := playlist_tool.CreateService(mockPlaylistToolRepository)
		playlistTool, err := playlistToolService.CreatePlaylistTool(upsertPlaylistToolData, playlistID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if playlistTool.ID != playlistToolData.ID {
			t.Errorf("Expected id %v, got %v", playlistTool.ID, playlistToolData.ID)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockPlaylistToolRepository.On("CreatePlaylistTool", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		playlistToolService := playlist_tool.CreateService(mockPlaylistToolRepository)
		_, err := playlistToolService.CreatePlaylistTool(upsertPlaylistToolData, playlistID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestDeletePlaylistTool(t *testing.T) {
	playlistToolData := &playlist_tool.PlaylistTool{
		ID:         1,
		PlaylistID: 1,
		ToolID:     1,
	}

	t.Run("success", func(t *testing.T) {
		mockPlaylistToolRepository.On("DeletePlaylistTool", playlistToolData.PlaylistID, playlistToolData.ToolID).Return(playlistToolData, nil).Once()

		playlistToolService := playlist_tool.CreateService(mockPlaylistToolRepository)
		playlistTool, err := playlistToolService.DeletePlaylistTool(playlistToolData.PlaylistID, playlistToolData.ToolID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if playlistTool.ID != playlistToolData.ID {
			t.Errorf("Expected id %v, got %v", playlistToolData.ID, playlistTool.ID)
		}

		if playlistTool.PlaylistID != playlistToolData.PlaylistID {
			t.Errorf("Expected playlist id %v, got %v", playlistToolData.PlaylistID, playlistTool.PlaylistID)
		}

		if playlistTool.ToolID != playlistToolData.ToolID {
			t.Errorf("Expected tool id %v, got %v", playlistToolData.ToolID, playlistTool.ToolID)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockPlaylistToolRepository.On("DeletePlaylistTool", playlistToolData.PlaylistID, playlistToolData.ToolID).Return(nil, errors.New("error unit testing")).Once()

		playlistToolService := playlist_tool.CreateService(mockPlaylistToolRepository)
		_, err := playlistToolService.DeletePlaylistTool(playlistToolData.PlaylistID, playlistToolData.ToolID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
