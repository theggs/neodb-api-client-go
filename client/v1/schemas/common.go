package schemas

import (
	"encoding/json"
	"fmt"
)

// SearchResult represents the search result
type SearchResult struct {
	Data  []Item `json:"data"`
	Pages int    `json:"pages"`
	Count int    `json:"count"`
}

// UnmarshalJSON is a custom unmarshal function for SearchResult
func (sr *SearchResult) UnmarshalJSON(data []byte) error {
	var raw struct {
		Data  []json.RawMessage `json:"data"`
		Pages int               `json:"pages"`
		Count int               `json:"count"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	sr.Pages = raw.Pages
	sr.Count = raw.Count

	for _, itemData := range raw.Data {
		var itemType struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(itemData, &itemType); err != nil {
			return err
		}

		var item Item
		switch itemType.Type {
		case "Edition":
			var edition EditionSchema
			if err := json.Unmarshal(itemData, &edition); err != nil {
				return err
			}
			item = edition
		case "Movie":
			var movie MovieSchema
			if err := json.Unmarshal(itemData, &movie); err != nil {
				return err
			}
			item = movie
		// Add other types as needed
		default:
			return fmt.Errorf("unknown item type: %s", itemType.Type)
		}

		sr.Data = append(sr.Data, item)
	}

	return nil
}
