package schemas

// AlbumSchema represents the detailed information of an album
type AlbumSchema struct {
	ItemSchema
	Tracks      []string `json:"track_list"`
	Artists     []string `json:"artists"`
	Genres      []string `json:"genres"`
	ReleaseDate string   `json:"release_date"`
}
