package schemas

// MovieSchema represents the detailed information of a movie
type MovieSchema struct {
	ItemSchema
	Directors   []string `json:"director"`
	Actors      []string `json:"actor"`
	Genres      []string `json:"genre"`
	Duration    string   `json:"duration"` // in minutes, stored as string
	ReleaseDate string   `json:"release_date"`
}

// GetType returns the type of the item
func (m MovieSchema) GetType() string {
	return "Movie"
}
