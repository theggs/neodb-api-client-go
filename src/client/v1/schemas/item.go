package schemas

// Item is a generic interface for all item types
type Item interface {
	GetType() string
}

// ItemSchema represents a brief information about an item
type ItemSchema struct {
	Title         string  `json:"title"`
	Brief         string  `json:"brief"`
	CoverImageURL string  `json:"cover_image_url"`
	Rating        float64 `json:"rating"`
	RatingCount   int     `json:"rating_count"`
	ID            string  `json:"id"`
	Type          string  `json:"type"`
	UUID          string  `json:"uuid"`
	URL           string  `json:"url"`
	APIURL        string  `json:"api_url"`
	Category      string  `json:"category"`
	ParentUUID    string  `json:"parent_uuid"`
	DisplayTitle  string  `json:"display_title"`
}
