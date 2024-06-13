package v1

import "fmt"

// GetSearchResults 搜索条目
func (c *NeoDBClient) GetSearchResults(query string, category string, page int) (*SearchResult, error) {
	endpoint := fmt.Sprintf("/api/catalog/search?query=%s&category=%s&page=%d", query, category, page)
	var searchResult SearchResult
	err := c.executeRequest(endpoint, &searchResult)
	if err != nil {
		return nil, err
	}
	return &searchResult, nil
}

// FetchItemFromURL 从支持的站点 URL 获取条目
func (c *NeoDBClient) FetchItemFromURL(url string) (*ItemSchema, error) {
	endpoint := fmt.Sprintf("/api/catalog/fetch?url=%s", url)
	var item ItemSchema
	err := c.executeRequest(endpoint, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetBook 获取书籍详细信息
func (c *NeoDBClient) GetBook(uuid string) (*EditionSchema, error) {
	endpoint := fmt.Sprintf("/api/book/%s", uuid)
	var book EditionSchema
	err := c.executeRequest(endpoint, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// GetMovie 获取电影详细信息
func (c *NeoDBClient) GetMovie(uuid string) (*MovieSchema, error) {
	endpoint := fmt.Sprintf("/api/movie/%s", uuid)
	var movie MovieSchema
	err := c.executeRequest(endpoint, &movie)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// GetTVShow 获取电视节目详细信息
func (c *NeoDBClient) GetTVShow(uuid string) (*TVShowSchema, error) {
	endpoint := fmt.Sprintf("/api/tv/%s", uuid)
	var tvShow TVShowSchema
	err := c.executeRequest(endpoint, &tvShow)
	if err != nil {
		return nil, err
	}
	return &tvShow, nil
}

// GetTVSeason 获取电视季详细信息
func (c *NeoDBClient) GetTVSeason(uuid string) (*TVSeasonSchema, error) {
	endpoint := fmt.Sprintf("/api/tv/season/%s", uuid)
	var tvSeason TVSeasonSchema
	err := c.executeRequest(endpoint, &tvSeason)
	if err != nil {
		return nil, err
	}
	return &tvSeason, nil
}

// GetTVEpisode 获取电视集详细信息
func (c *NeoDBClient) GetTVEpisode(uuid string) (*TVEpisodeSchema, error) {
	endpoint := fmt.Sprintf("/api/tv/episode/%s", uuid)
	var tvEpisode TVEpisodeSchema
	err := c.executeRequest(endpoint, &tvEpisode)
	if err != nil {
		return nil, err
	}
	return &tvEpisode, nil
}

// GetPodcast 获取播客详细信息
func (c *NeoDBClient) GetPodcast(uuid string) (*PodcastSchema, error) {
	endpoint := fmt.Sprintf("/api/podcast/%s", uuid)
	var podcast PodcastSchema
	err := c.executeRequest(endpoint, &podcast)
	if err != nil {
		return nil, err
	}
	return &podcast, nil
}

// GetAlbum 获取专辑详细信息
func (c *NeoDBClient) GetAlbum(uuid string) (*AlbumSchema, error) {
	endpoint := fmt.Sprintf("/api/album/%s", uuid)
	var album AlbumSchema
	err := c.executeRequest(endpoint, &album)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

// GetGame 获取游戏详细信息
func (c *NeoDBClient) GetGame(uuid string) (*GameSchema, error) {
	endpoint := fmt.Sprintf("/api/game/%s", uuid)
	var game GameSchema
	err := c.executeRequest(endpoint, &game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

// GetPerformance 获取演出详细信息
func (c *NeoDBClient) GetPerformance(uuid string) (*PerformanceSchema, error) {
	endpoint := fmt.Sprintf("/api/performance/%s", uuid)
	var performance PerformanceSchema
	err := c.executeRequest(endpoint, &performance)
	if err != nil {
		return nil, err
	}
	return &performance, nil
}

// GetPerformanceProduction 获取演出制作详细信息
func (c *NeoDBClient) GetPerformanceProduction(uuid string) (*PerformanceSchema, error) {
	endpoint := fmt.Sprintf("/api/performance/production/%s", uuid)
	var performance PerformanceSchema
	err := c.executeRequest(endpoint, &performance)
	if err != nil {
		return nil, err
	}
	return &performance, nil
}
