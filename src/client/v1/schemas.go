package v1

// SearchResult 表示搜索结果
type SearchResult struct {
	Data  []ItemSchema `json:"data"`
	Pages int          `json:"pages"`
	Count int          `json:"count"`
}

// ItemSchema 表示一个条目的详细信息
type ItemSchema struct {
	Title             string             `json:"title"`
	Brief             string             `json:"brief"`
	CoverImageURL     string             `json:"cover_image_url"`
	Rating            float64            `json:"rating"`
	RatingCount       int                `json:"rating_count"`
	ID                string             `json:"id"`
	Type              string             `json:"type"`
	UUID              string             `json:"uuid"`
	URL               string             `json:"url"`
	APIURL            string             `json:"api_url"`
	Category          string             `json:"category"`
	ParentUUID        string             `json:"parent_uuid"`
	DisplayTitle      string             `json:"display_title"`
	ExternalResources []ExternalResource `json:"external_resources"`
}

// ExternalResource 表示一个外部资源
type ExternalResource struct {
	URL string `json:"url"`
}

// EditionSchema 表示书籍的详细信息
type EditionSchema struct {
	ItemSchema
	Authors    []string `json:"author"`
	Translator []string `json:"translator"`
	ISBN       string   `json:"isbn"`
}

// MovieSchema 表示电影的详细信息
type MovieSchema struct {
	ItemSchema
	Directors []string `json:"director"`
	Actors    []string `json:"actor"`
}

// TVShowSchema 表示电视节目的详细信息
type TVShowSchema struct {
	ItemSchema
	Seasons []string `json:"season_uuids"`
}

// TVSeasonSchema 表示电视季的详细信息
type TVSeasonSchema struct {
	ItemSchema
	Episodes []string `json:"episode_uuids"`
}

// TVEpisodeSchema 表示电视集的详细信息
type TVEpisodeSchema struct {
	ItemSchema
	EpisodeNumber int `json:"episode_number"`
}

// PodcastSchema 表示播客的详细信息
type PodcastSchema struct {
	ItemSchema
	Hosts []string `json:"hosts"`
}

// AlbumSchema 表示专辑的详细信息
type AlbumSchema struct {
	ItemSchema
	Tracks []string `json:"track_list"`
}

// GameSchema 表示游戏的详细信息
type GameSchema struct {
	ItemSchema
	Platforms []string `json:"platform"`
}

// PerformanceSchema 表示演出的详细信息
type PerformanceSchema struct {
	ItemSchema
	Performers []string `json:"performer"`
}
