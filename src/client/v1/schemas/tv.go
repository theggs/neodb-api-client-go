package schemas

// TVShowSchema represents the detailed information of a TV show
type TVShowSchema struct {
	ItemSchema
	Seasons []string `json:"season_uuids"`
}

// TVSeasonSchema represents the detailed information of a TV season
type TVSeasonSchema struct {
	ItemSchema
	Episodes     []string `json:"episode_uuids"`
	SeasonNumber int      `json:"season_number"`
}

// TVEpisodeSchema represents the detailed information of a TV episode
type TVEpisodeSchema struct {
	ItemSchema
	EpisodeNumber int    `json:"episode_number"`
	Duration      int    `json:"duration"` // in minutes
	AirDate       string `json:"air_date"`
}
