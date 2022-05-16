// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	video "vidlearn-final-projcect/business/video"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateVideo provides a mock function with given fields: _a0
func (_m *Repository) CreateVideo(_a0 *video.Video) (*video.Video, error) {
	ret := _m.Called(_a0)

	var r0 *video.Video
	if rf, ok := ret.Get(0).(func(*video.Video) *video.Video); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*video.Video)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*video.Video) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVideo provides a mock function with given fields: ID
func (_m *Repository) DeleteVideo(ID int) (*video.Video, error) {
	ret := _m.Called(ID)

	var r0 *video.Video
	if rf, ok := ret.Get(0).(func(int) *video.Video); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*video.Video)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVideoByID provides a mock function with given fields: ID
func (_m *Repository) GetVideoByID(ID int) (*video.Video, error) {
	ret := _m.Called(ID)

	var r0 *video.Video
	if rf, ok := ret.Get(0).(func(int) *video.Video); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*video.Video)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVideoByPlaylist provides a mock function with given fields: playlistID
func (_m *Repository) GetVideoByPlaylist(playlistID int) ([]*video.Video, error) {
	ret := _m.Called(playlistID)

	var r0 []*video.Video
	if rf, ok := ret.Get(0).(func(int) []*video.Video); ok {
		r0 = rf(playlistID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*video.Video)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(playlistID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateVideo provides a mock function with given fields: videoCurrent, IDCurrent
func (_m *Repository) UpdateVideo(videoCurrent *video.Video, IDCurrent int) (*video.Video, error) {
	ret := _m.Called(videoCurrent, IDCurrent)

	var r0 *video.Video
	if rf, ok := ret.Get(0).(func(*video.Video, int) *video.Video); ok {
		r0 = rf(videoCurrent, IDCurrent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*video.Video)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*video.Video, int) error); ok {
		r1 = rf(videoCurrent, IDCurrent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t testing.TB) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}