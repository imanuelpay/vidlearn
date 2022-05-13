package spec

type UpsertPlaylistSpec struct {
	CategoryID   int    `validate:"required"`
	Name         string `validate:"required"`
	ThumbnailURL string `validate:"required"`
	Description  string `validate:"required"`
	Level        string `validate:"required"`
}
