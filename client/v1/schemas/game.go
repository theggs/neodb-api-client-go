package schemas

// GameSchema represents the detailed information of a game
type GameSchema struct {
	ItemSchema
	Platforms   []string `json:"platforms"`
	Developers  []string `json:"developers"`
	Genres      []string `json:"genres"`
	ReleaseDate string   `json:"release_date"`
}
