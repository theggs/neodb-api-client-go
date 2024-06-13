# Schema Documentation

## EditionSchema

EditionSchema represents the detailed information of a book. Below are the fields included in the schema:

- **Authors**: List of authors of the book.
- **Translators**: List of translators of the book.
- **Editors**: List of editors of the book.
- **ISBN**: ISBN number of the book.
- **Publisher**: Publisher of the book.
- **Pages**: Number of pages in the book.
- **Binding**: Binding type of the book.
- **PublicationDate**: Publication date of the book.
- **OriginalTitle**: Original title of the book.
- **Language**: Language of the book.
- **Genres**: Genres or categories of the book.
- **Summary**: Summary or brief description of the book.
- **Tags**: List of tags associated with the book.
- **CoverImageURL**: URL of the cover image of the book.
- **Rating**: Rating of the book.
- **RatingCount**: Number of ratings of the book.
- **Series**: Series the book belongs to.
- **SeriesOrder**: Order of the book in the series.
- **ExternalLinks**: List of external links related to the book, including site name and URL.

### Example

```go
package schemas

// EditionSchema represents the detailed information of a book
type EditionSchema struct {
ItemSchema
Authors         []string `json:"authors"`
Translators     []string `json:"translators"`
Editors         []string `json:"editors"`
ISBN            string   `json:"isbn"`
Publisher       string   `json:"publisher"`
Pages           int      `json:"pages"`
Binding         string   `json:"binding"`
PublicationDate string   `json:"publication_date"`
OriginalTitle   string   `json:"original_title"`
Language        string   `json:"language"`
Genres          []string `json:"genres"`
Summary         string   `json:"summary"`
Tags            []string `json:"tags"`
CoverImageURL   string   `json:"cover_image_url"`
Rating          float64  `json:"rating"`
RatingCount     int      `json:"rating_count"`
Series          string   `json:"series"`
SeriesOrder     int      `json:"series_order"`
ExternalLinks   []struct {
Site string `json:"site"`
URL  string `json:"url"`
} `json:"external_links"`
}
```