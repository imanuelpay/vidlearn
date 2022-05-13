package spec

type UpsertVideoSpec struct {
	PlaylistID  int    `validate:"required"`
	Title       string `validate:"required"`
	VideoURL    string `validate:"required"`
	Description string `validate:"required"`
}
