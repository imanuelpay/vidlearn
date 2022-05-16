package video_test

import (
	"errors"
	"testing"
	"time"
	"vidlearn-final-projcect/business/video"
	videoMocks "vidlearn-final-projcect/business/video/mocks"
	"vidlearn-final-projcect/business/video/spec"

	"github.com/stretchr/testify/mock"
)

var (
	mockVideoRepository = new(videoMocks.Repository)
)

func TestGetVideoByPlaylist(t *testing.T) {
	playlistID := 1
	videoData := []*video.Video{
		{
			ID:          1,
			PlaylistID:  playlistID,
			Title:       "Test Video 2",
			VideoURL:    "https://vid.test/test.mp4",
			Description: "Test Description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			PlaylistID:  playlistID,
			Title:       "Test Video 2",
			VideoURL:    "https://vid.test/test.mp4",
			Description: "Test Description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	t.Run("success", func(t *testing.T) {
		mockVideoRepository.On("GetVideoByPlaylist", mock.Anything).Return(videoData, nil).Once()

		videoService := video.CreateService(mockVideoRepository)
		video, err := videoService.GetVideoByPlaylist(playlistID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(video) != len(videoData) {
			t.Errorf("Expected length data %v, got %v", len(video), len(videoData))
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockVideoRepository.On("GetVideoByPlaylist", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		videoService := video.CreateService(mockVideoRepository)
		_, err := videoService.GetVideoByPlaylist(playlistID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestGetVideoByID(t *testing.T) {
	videoData := &video.Video{
		ID:          1,
		PlaylistID:  1,
		Title:       "Test Video",
		VideoURL:    "https://vid.test/test.mp4",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockVideoRepository.On("GetVideoByID", mock.Anything).Return(videoData, nil).Once()

		videoService := video.CreateService(mockVideoRepository)
		video, err := videoService.GetVideoByID(videoData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if video.ID != videoData.ID {
			t.Errorf("Expected video ID %v, got %v", videoData.ID, video.ID)
		}

		if video.PlaylistID != videoData.PlaylistID {
			t.Errorf("Expected video playlist ID %v, got %v", videoData.PlaylistID, video.PlaylistID)
		}

		if video.Title != videoData.Title {
			t.Errorf("Expected video title %v, got %v", videoData.Title, video.Title)
		}

		if video.VideoURL != videoData.VideoURL {
			t.Errorf("Expected video URL %v, got %v", videoData.VideoURL, video.VideoURL)
		}

		if video.Description != videoData.Description {
			t.Errorf("Expected video description %v, got %v", videoData.Description, video.Description)
		}

		if video.CreatedAt != videoData.CreatedAt {
			t.Errorf("Expected video created at %v, got %v", videoData.CreatedAt, video.CreatedAt)
		}

		if video.UpdatedAt != videoData.UpdatedAt {
			t.Errorf("Expected video updated at %v, got %v", videoData.UpdatedAt, video.UpdatedAt)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockVideoRepository.On("GetVideoByID", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		videoService := video.CreateService(mockVideoRepository)
		_, err := videoService.GetVideoByID(videoData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestCreateVideo(t *testing.T) {
	upsertVideoData := &spec.UpsertVideoSpec{
		PlaylistID:  1,
		Title:       "Test Video",
		VideoURL:    "https://vid.test/test.mp4",
		Description: "Test Description",
	}

	videoData := &video.Video{}

	t.Run("success", func(t *testing.T) {
		mockVideoRepository.On("CreateVideo", mock.Anything).Return(videoData, nil).Once()

		videoService := video.CreateService(mockVideoRepository)
		video, err := videoService.CreateVideo(upsertVideoData)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if video.ID != videoData.ID {
			t.Errorf("Expected video ID %v, got %v", videoData.ID, video.ID)
		}

		if video.PlaylistID != videoData.PlaylistID {
			t.Errorf("Expected video playlist ID %v, got %v", videoData.PlaylistID, video.PlaylistID)
		}

		if video.Title != videoData.Title {
			t.Errorf("Expected video title %v, got %v", videoData.Title, video.Title)
		}

		if video.VideoURL != videoData.VideoURL {
			t.Errorf("Expected video URL %v, got %v", videoData.VideoURL, video.VideoURL)
		}

		if video.Description != videoData.Description {
			t.Errorf("Expected video description %v, got %v", videoData.Description, video.Description)
		}

		if video.CreatedAt != videoData.CreatedAt {
			t.Errorf("Expected video created at %v, got %v", videoData.CreatedAt, video.CreatedAt)
		}

		if video.UpdatedAt != videoData.UpdatedAt {
			t.Errorf("Expected video updated at %v, got %v", videoData.UpdatedAt, video.UpdatedAt)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockVideoRepository.On("CreateVideo", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		videoService := video.CreateService(mockVideoRepository)
		_, err := videoService.CreateVideo(upsertVideoData)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestUpdateVideo(t *testing.T) {
	videoData := &video.Video{
		ID: 1,
	}

	upsertVideoData := &spec.UpsertVideoSpec{
		PlaylistID:  1,
		Title:       "Test Video",
		VideoURL:    "https://vid.test/test.mp4",
		Description: "Test Description",
	}

	t.Run("success", func(t *testing.T) {
		mockVideoRepository.On("GetVideoByID", videoData.ID).Return(videoData, nil).Once()
		mockVideoRepository.On("UpdateVideo", mock.Anything, videoData.ID).Return(videoData, nil).Once()

		videoService := video.CreateService(mockVideoRepository)
		video, err := videoService.UpdateVideo(upsertVideoData, videoData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if video.ID != videoData.ID {
			t.Errorf("Expected video ID %v, got %v", videoData.ID, video.ID)
		}

		if video.PlaylistID != videoData.PlaylistID {
			t.Errorf("Expected video playlist ID %v, got %v", videoData.PlaylistID, video.PlaylistID)
		}

		if video.Title != videoData.Title {
			t.Errorf("Expected video title %v, got %v", videoData.Title, video.Title)
		}

		if video.VideoURL != videoData.VideoURL {
			t.Errorf("Expected video URL %v, got %v", videoData.VideoURL, video.VideoURL)
		}

		if video.Description != videoData.Description {
			t.Errorf("Expected video description %v, got %v", videoData.Description, video.Description)
		}

		if video.CreatedAt != videoData.CreatedAt {
			t.Errorf("Expected video created at %v, got %v", videoData.CreatedAt, video.CreatedAt)
		}

		if video.UpdatedAt != videoData.UpdatedAt {
			t.Errorf("Expected video updated at %v, got %v", videoData.UpdatedAt, video.UpdatedAt)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockVideoRepository.On("GetVideoByID", videoData.ID).Return(nil, errors.New("error unit testing")).Once()
		mockVideoRepository.On("UpdateVideo", mock.Anything, videoData.ID).Return(nil, errors.New("error unit testing")).Once()

		videoService := video.CreateService(mockVideoRepository)
		_, err := videoService.UpdateVideo(upsertVideoData, videoData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestDeleteVideo(t *testing.T) {
	videoData := &video.Video{
		ID:          1,
		PlaylistID:  1,
		Title:       "Test Video",
		VideoURL:    "https://vid.test/test.mp4",
		Description: "Test Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockVideoRepository.On("DeleteVideo", mock.Anything).Return(videoData, nil).Once()

		videoService := video.CreateService(mockVideoRepository)
		video, err := videoService.DeleteVideo(videoData.ID)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if video.ID != videoData.ID {
			t.Errorf("Expected video ID %v, got %v", videoData.ID, video.ID)
		}

		if video.PlaylistID != videoData.PlaylistID {
			t.Errorf("Expected video playlist ID %v, got %v", videoData.PlaylistID, video.PlaylistID)
		}

		if video.Title != videoData.Title {
			t.Errorf("Expected video title %v, got %v", videoData.Title, video.Title)
		}

		if video.VideoURL != videoData.VideoURL {
			t.Errorf("Expected video URL %v, got %v", videoData.VideoURL, video.VideoURL)
		}

		if video.Description != videoData.Description {
			t.Errorf("Expected video description %v, got %v", videoData.Description, video.Description)
		}

		if video.CreatedAt != videoData.CreatedAt {
			t.Errorf("Expected video created at %v, got %v", videoData.CreatedAt, video.CreatedAt)
		}

		if video.UpdatedAt != videoData.UpdatedAt {
			t.Errorf("Expected video updated at %v, got %v", videoData.UpdatedAt, video.UpdatedAt)
		}
	})

	t.Run("failure", func(t *testing.T) {
		mockVideoRepository.On("DeleteVideo", mock.Anything).Return(nil, errors.New("error unit testing")).Once()

		videoService := video.CreateService(mockVideoRepository)
		_, err := videoService.DeleteVideo(videoData.ID)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
