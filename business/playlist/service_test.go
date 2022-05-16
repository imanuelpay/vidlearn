package playlist_test

import (
	"errors"
	"testing"
	"time"
	"vidlearn-final-projcect/business/playlist"
	playlistMocks "vidlearn-final-projcect/business/playlist/mocks"
	"vidlearn-final-projcect/business/playlist/spec"

	"github.com/stretchr/testify/mock"
)

var (
	mockPlaylistRespository = new(playlistMocks.Repository)
)

func TestGetAllPlaylist(t *testing.T) {
	playlistData := []*playlist.Playlist{}

	t.Run("success", func(t *testing.T) {
		mockPlaylistRespository.On("GetAllPlaylist", mock.Anything).Return(playlistData, nil).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		playlist, err := playlistService.GetAllPlaylist()

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(playlist) != len(playlistData) {
			t.Errorf("Expected length data %v, got %v", len(playlist), len(playlistData))
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockPlaylistRespository.On("GetAllPlaylist", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		_, err := playlistService.GetAllPlaylist()

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestGetPlaylistByID(t *testing.T) {
	playlistData := &playlist.Playlist{
		ID:           1,
		CategoryID:   1,
		Name:         "Test Playlist",
		ThumbnailURL: "https://img.test/test.png",
		Description:  "Test Description",
		Level:        "Beginner",
		CreatedAt:    time.Now(),
		CreatedBy:    "Admin Test",
		UpdatedAt:    time.Now(),
		UpdatedBy:    "Admin Test",
	}

	t.Run("success", func(t *testing.T) {
		mockPlaylistRespository.On("GetPlaylistByID", mock.Anything).Return(playlistData, nil).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		playlist, err := playlistService.GetPlaylistByID(playlistData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if playlist.ID != playlistData.ID {
			t.Errorf("Expected playlist ID %v, got %v", playlistData.ID, playlist.ID)
		}

		if playlist.CategoryID != playlistData.CategoryID {
			t.Errorf("Expected playlist category ID %v, got %v", playlistData.CategoryID, playlist.CategoryID)
		}

		if playlist.Name != playlistData.Name {
			t.Errorf("Expected playlist name %v, got %v", playlistData.Name, playlist.Name)
		}

		if playlist.ThumbnailURL != playlistData.ThumbnailURL {
			t.Errorf("Expected playlist thumbnail URL %v, got %v", playlistData.ThumbnailURL, playlist.ThumbnailURL)
		}

		if playlist.Description != playlistData.Description {
			t.Errorf("Expected playlist description %v, got %v", playlistData.Description, playlist.Description)
		}

		if playlist.Level != playlistData.Level {
			t.Errorf("Expected playlist level %v, got %v", playlistData.Level, playlist.Level)
		}

		if playlist.CreatedAt != playlistData.CreatedAt {
			t.Errorf("Expected playlist created at %v, got %v", playlistData.CreatedAt, playlist.CreatedAt)
		}

		if playlist.CreatedBy != playlistData.CreatedBy {
			t.Errorf("Expected playlist created by %v, got %v", playlistData.CreatedBy, playlist.CreatedBy)
		}

		if playlist.UpdatedAt != playlistData.UpdatedAt {
			t.Errorf("Expected playlist updated at %v, got %v", playlistData.UpdatedAt, playlist.UpdatedAt)
		}

		if playlist.UpdatedBy != playlistData.UpdatedBy {
			t.Errorf("Expected playlist updated by %v, got %v", playlistData.UpdatedBy, playlist.UpdatedBy)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockPlaylistRespository.On("GetPlaylistByID", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		_, err := playlistService.GetPlaylistByID(playlistData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestCreatePlaylist(t *testing.T) {
	upsertPlaylistData := &spec.UpsertPlaylistSpec{
		CategoryID:   1,
		Name:         "Test Playlist",
		ThumbnailURL: "https://img.test/test.png",
		Description:  "Test Description",
		Level:        "Beginner",
	}

	playlistData := &playlist.Playlist{}

	t.Run("success", func(t *testing.T) {
		mockPlaylistRespository.On("CreatePlaylist", mock.Anything).Return(playlistData, nil).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		playlist, err := playlistService.CreatePlaylist(upsertPlaylistData, "Admin Test")

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if playlist.ID != playlistData.ID {
			t.Errorf("Expected playlist ID %v, got %v", playlistData.ID, playlist.ID)
		}

		if playlist.CategoryID != playlistData.CategoryID {
			t.Errorf("Expected playlist category ID %v, got %v", playlistData.CategoryID, playlist.CategoryID)
		}

		if playlist.Name != playlistData.Name {
			t.Errorf("Expected playlist name %v, got %v", playlistData.Name, playlist.Name)
		}

		if playlist.ThumbnailURL != playlistData.ThumbnailURL {
			t.Errorf("Expected playlist thumbnail URL %v, got %v", playlistData.ThumbnailURL, playlist.ThumbnailURL)
		}

		if playlist.Description != playlistData.Description {
			t.Errorf("Expected playlist description %v, got %v", playlistData.Description, playlist.Description)
		}

		if playlist.Level != playlistData.Level {
			t.Errorf("Expected playlist level %v, got %v", playlistData.Level, playlist.Level)
		}

		if playlist.CreatedAt != playlistData.CreatedAt {
			t.Errorf("Expected playlist created at %v, got %v", playlistData.CreatedAt, playlist.CreatedAt)
		}

		if playlist.CreatedBy != playlistData.CreatedBy {
			t.Errorf("Expected playlist created by %v, got %v", playlistData.CreatedBy, playlist.CreatedBy)
		}

		if playlist.UpdatedAt != playlistData.UpdatedAt {
			t.Errorf("Expected playlist updated at %v, got %v", playlistData.UpdatedAt, playlist.UpdatedAt)
		}

		if playlist.UpdatedBy != playlistData.UpdatedBy {
			t.Errorf("Expected playlist updated by %v, got %v", playlistData.UpdatedBy, playlist.UpdatedBy)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockPlaylistRespository.On("CreatePlaylist", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		_, err := playlistService.CreatePlaylist(upsertPlaylistData, "Admin Test")

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestUpdatePlaylist(t *testing.T) {
	playlistData := &playlist.Playlist{
		ID: 1,
	}

	upsertPlaylistData := &spec.UpsertPlaylistSpec{
		CategoryID:   1,
		Name:         "Test Playlist",
		ThumbnailURL: "https://img.test/test.png",
		Description:  "Test Description",
		Level:        "Beginner",
	}

	t.Run("success", func(t *testing.T) {
		mockPlaylistRespository.On("GetPlaylistByID", playlistData.ID).Return(playlistData, nil).Once()
		mockPlaylistRespository.On("UpdatePlaylist", mock.Anything, playlistData.ID).Return(playlistData, nil).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		playlist, err := playlistService.UpdatePlaylist(upsertPlaylistData, "Admin Test", playlistData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if playlist.ID != playlistData.ID {
			t.Errorf("Expected playlist ID %v, got %v", playlistData.ID, playlist.ID)
		}

		if playlist.CategoryID != playlistData.CategoryID {
			t.Errorf("Expected playlist category ID %v, got %v", playlistData.CategoryID, playlist.CategoryID)
		}

		if playlist.Name != playlistData.Name {
			t.Errorf("Expected playlist name %v, got %v", playlistData.Name, playlist.Name)
		}

		if playlist.ThumbnailURL != playlistData.ThumbnailURL {
			t.Errorf("Expected playlist thumbnail URL %v, got %v", playlistData.ThumbnailURL, playlist.ThumbnailURL)
		}

		if playlist.Description != playlistData.Description {
			t.Errorf("Expected playlist description %v, got %v", playlistData.Description, playlist.Description)
		}

		if playlist.Level != playlistData.Level {
			t.Errorf("Expected playlist level %v, got %v", playlistData.Level, playlist.Level)
		}

		if playlist.CreatedAt != playlistData.CreatedAt {
			t.Errorf("Expected playlist created at %v, got %v", playlistData.CreatedAt, playlist.CreatedAt)
		}

		if playlist.CreatedBy != playlistData.CreatedBy {
			t.Errorf("Expected playlist created by %v, got %v", playlistData.CreatedBy, playlist.CreatedBy)
		}

		if playlist.UpdatedAt != playlistData.UpdatedAt {
			t.Errorf("Expected playlist updated at %v, got %v", playlistData.UpdatedAt, playlist.UpdatedAt)
		}

		if playlist.UpdatedBy != playlistData.UpdatedBy {
			t.Errorf("Expected playlist updated by %v, got %v", playlistData.UpdatedBy, playlist.UpdatedBy)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockPlaylistRespository.On("GetPlaylistByID", playlistData.ID).Return(nil, errors.New("error unit testing")).Once()
		mockPlaylistRespository.On("UpdatePlaylist", mock.Anything, playlistData.ID).Return(nil, errors.New("error unit testing")).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		_, err := playlistService.UpdatePlaylist(upsertPlaylistData, "Admin Test", playlistData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestDeletePlaylist(t *testing.T) {
	playlistData := &playlist.Playlist{
		ID:           1,
		CategoryID:   1,
		Name:         "Test Playlist",
		ThumbnailURL: "https://img.test/test.png",
		Description:  "Test Description",
		Level:        "Beginner",
		CreatedAt:    time.Now(),
		CreatedBy:    "Admin Test",
		UpdatedAt:    time.Now(),
		UpdatedBy:    "Admin Test",
	}

	t.Run("success", func(t *testing.T) {
		mockPlaylistRespository.On("DeletePlaylist", mock.Anything).Return(playlistData, nil).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		playlist, err := playlistService.DeletePlaylist(playlistData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if playlist.ID != playlistData.ID {
			t.Errorf("Expected playlist ID %v, got %v", playlistData.ID, playlist.ID)
		}

		if playlist.CategoryID != playlistData.CategoryID {
			t.Errorf("Expected playlist category ID %v, got %v", playlistData.CategoryID, playlist.CategoryID)
		}

		if playlist.Name != playlistData.Name {
			t.Errorf("Expected playlist name %v, got %v", playlistData.Name, playlist.Name)
		}

		if playlist.ThumbnailURL != playlistData.ThumbnailURL {
			t.Errorf("Expected playlist thumbnail URL %v, got %v", playlistData.ThumbnailURL, playlist.ThumbnailURL)
		}

		if playlist.Description != playlistData.Description {
			t.Errorf("Expected playlist description %v, got %v", playlistData.Description, playlist.Description)
		}

		if playlist.Level != playlistData.Level {
			t.Errorf("Expected playlist level %v, got %v", playlistData.Level, playlist.Level)
		}

		if playlist.CreatedAt != playlistData.CreatedAt {
			t.Errorf("Expected playlist created at %v, got %v", playlistData.CreatedAt, playlist.CreatedAt)
		}

		if playlist.CreatedBy != playlistData.CreatedBy {
			t.Errorf("Expected playlist created by %v, got %v", playlistData.CreatedBy, playlist.CreatedBy)
		}

		if playlist.UpdatedAt != playlistData.UpdatedAt {
			t.Errorf("Expected playlist updated at %v, got %v", playlistData.UpdatedAt, playlist.UpdatedAt)
		}

		if playlist.UpdatedBy != playlistData.UpdatedBy {
			t.Errorf("Expected playlist updated by %v, got %v", playlistData.UpdatedBy, playlist.UpdatedBy)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockPlaylistRespository.On("DeletePlaylist", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		playlistService := playlist.CreateService(mockPlaylistRespository)
		_, err := playlistService.DeletePlaylist(playlistData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
