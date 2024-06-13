package schemas

// EditionSchema represents the detailed information of a book
type EditionSchema struct {
	ItemSchema
	Author        []string `json:"author"`
	Translator    []string `json:"translator"`
	Editor        []string `json:"editor"`
	ISBN          string   `json:"isbn"`
	PubHouse      string   `json:"pub_house"`
	Pages         int      `json:"pages"`
	Binding       string   `json:"binding"`
	PubYear       int      `json:"pub_year"`
	PubMonth      int      `json:"pub_month"`
	OriginalTitle string   `json:"orig_title"`
	Language      string   `json:"language"`
	Genres        []string `json:"genres"`
	Summary       string   `json:"summary"`
	Tags          []string `json:"tags"`
	CoverImageURL string   `json:"cover_image_url"`
	Rating        *float64 `json:"rating"`
	RatingCount   int      `json:"rating_count"`
	Series        string   `json:"series"`
	SeriesOrder   int      `json:"series_order"`
	ExternalLinks []struct {
		URL string `json:"url"`
	} `json:"external_resources"`
	Subtitle string `json:"subtitle"`
	Price    string `json:"price"`
	Imprint  string `json:"imprint"`
}

// GetType returns the type of the item
func (e EditionSchema) GetType() string {
	return "Edition"
}
