package schemas

// PodcastSchema represents the detailed information of a podcast
type PodcastSchema struct {
	ItemSchema
	Hosts    []string `json:"hosts"`
	Episodes []string `json:"episode_uuids"`
}
