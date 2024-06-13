package v1

import (
	"fmt"

	"github.com/theggs/neodb-api-client-go/client/v1/schemas"
)

// GetSearchResults 搜索条目
func (c *NeoDBClient) GetSearchResults(query string, category string, page int) (*schemas.SearchResult, error) {
	endpoint := fmt.Sprintf("/api/catalog/search?query=%s&category=%s&page=%d", query, category, page)
	var searchResult schemas.SearchResult
	err := c.executeRequest(endpoint, &searchResult)
	if err != nil {
		return nil, err
	}
	return &searchResult, nil
}

// FetchItemFromURL 从支持的站点 URL 获取条目
func (c *NeoDBClient) FetchItemFromURL(url string) (*schemas.ItemSchema, error) {
	endpoint := fmt.Sprintf("/api/catalog/fetch?url=%s", url)
	var item schemas.ItemSchema
	err := c.executeRequest(endpoint, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetBook 获取书籍详细信息
func (c *NeoDBClient) GetBook(uuid string) (*schemas.EditionSchema, error) {
	endpoint := fmt.Sprintf("/api/book/%s", uuid)
	var book schemas.EditionSchema
	err := c.executeRequest(endpoint, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// GetMovie 获取电影详细信息
func (c *NeoDBClient) GetMovie(uuid string) (*schemas.MovieSchema, error) {
	endpoint := fmt.Sprintf("/api/movie/%s", uuid)
	var movie schemas.MovieSchema
	err := c.executeRequest(endpoint, &movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

// GetTVShow 获取电视节目详细信息
func (c *NeoDBClient) GetTVShow(uuid string) (*schemas.TVShowSchema, error) {
	endpoint := fmt.Sprintf("/api/tv/%s", uuid)
	var tvShow schemas.TVShowSchema
	err := c.executeRequest(endpoint, &tvShow)
	if err != nil {
		return nil, err
	}
	return &tvShow, nil
}

// GetTVSeason 获取电视季详细信息
func (c *NeoDBClient) GetTVSeason(uuid string) (*schemas.TVSeasonSchema, error) {
	endpoint := fmt.Sprintf("/api/tv/season/%s", uuid)
	var tvSeason schemas.TVSeasonSchema
	err := c.executeRequest(endpoint, &tvSeason)
	if err != nil {
		return nil, err
	}
	return &tvSeason, nil
}

// GetTVEpisode 获取电视集详细信息
func (c *NeoDBClient) GetTVEpisode(uuid string) (*schemas.TVEpisodeSchema, error) {
	endpoint := fmt.Sprintf("/api/tv/episode/%s", uuid)
	var tvEpisode schemas.TVEpisodeSchema
	err := c.executeRequest(endpoint, &tvEpisode)
	if err != nil {
		return nil, err
	}
	return &tvEpisode, nil
}

// GetPodcast 获取播客详细信息
func (c *NeoDBClient) GetPodcast(uuid string) (*schemas.PodcastSchema, error) {
	endpoint := fmt.Sprintf("/api/podcast/%s", uuid)
	var podcast schemas.PodcastSchema
	err := c.executeRequest(endpoint, &podcast)
	if err != nil {
		return nil, err
	}
	return &podcast, nil
}

// GetAlbum 获取专辑详细信息
func (c *NeoDBClient) GetAlbum(uuid string) (*schemas.AlbumSchema, error) {
	endpoint := fmt.Sprintf("/api/album/%s", uuid)
	var album schemas.AlbumSchema
	err := c.executeRequest(endpoint, &album)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

// GetGame 获取游戏详细信息
func (c *NeoDBClient) GetGame(uuid string) (*schemas.GameSchema, error) {
	endpoint := fmt.Sprintf("/api/game/%s", uuid)
	var game schemas.GameSchema
	err := c.executeRequest(endpoint, &game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

// GetPerformance 获取演出详细信息
func (c *NeoDBClient) GetPerformance(uuid string) (*schemas.PerformanceSchema, error) {
	endpoint := fmt.Sprintf("/api/performance/%s", uuid)
	var performance schemas.PerformanceSchema
	err := c.executeRequest(endpoint, &performance)
	if err != nil {
		return nil, err
	}
	return &performance, nil
}

// GetPerformanceProduction 获取演出制作详细信息
func (c *NeoDBClient) GetPerformanceProduction(uuid string) (*schemas.PerformanceProductionSchema, error) {
	endpoint := fmt.Sprintf("/api/performance/production/%s", uuid)
	var performance schemas.PerformanceProductionSchema
	err := c.executeRequest(endpoint, &performance)
	if err != nil {
		return nil, err
	}
	return &performance, nil
}
