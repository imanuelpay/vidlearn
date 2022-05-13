package spec

type UpsertPlaylistToolSpec struct {
	PlaylistID int `validate:"required"`
	ToolID     int `validate:"required"`
}
