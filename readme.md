# NeoDB Client for Go

NeoDB Client for Go is a client library for interacting with the [neodb.social](https://neodb.social) API. It provides a set of methods for searching entries and retrieving details about books, movies, TV shows, and more.

**Only implemented the following endpoints:**
- Search entries
- Get book details

## Features

- Search entries
- Get book details
- Get movie details(not fully implemented)
- Get TV show details(not fully implemented)
- Get TV season details(not fully implemented)
- Get TV episode details(not fully implemented)
- Get podcast details(not fully implemented)
- Get album details(not fully implemented)
- Get game details(not fully implemented)
- Get performance details(not fully implemented)
- Get performance production details(not fully implemented)

## Installation

First, make sure you have Go installed. Then, use the `go get` command to install the required dependencies:

```
go get -u github.com/go-resty/resty/v2
```

Clone this project to your local machine:

```
git clone https://github.com/yourusername/neodbclient
cd neodbclient
```

## Usage

### Creating the Client

You need to create a client instance and provide the API base URL.

```
package main

import (
"fmt"
"log"

    client "neodbclient/client/v1" // Using alias client
)

func main() {
apiBaseURL := "https://neodb.social"
neodbClient := client.NewNeoDBClient(apiBaseURL)

    // Example: Search entries
    searchQuery := "Vagabond"
    category := "book"
    page := 1
    searchResult, err := neodbClient.GetSearchResults(searchQuery, category, page)
    if err != nil {
        log.Fatalf("Error searching entries: %v", err)
    }
    fmt.Printf("Search results: %+v\n", searchResult)

    // Example: Get book details
    bookUUID := "1obNw8IqrMSS5DUllN4cyG"
    book, err := neodbClient.GetBook(bookUUID)
    if err != nil {
        log.Fatalf("Error getting book details: %v", err)
    }
    fmt.Printf("Book details: %+v\n", book)

    // Example: Get movie details
    movieUUID := "3BU3rgfjc2PaymhsN2rxwO"
    movie, err := neodbClient.GetMovie(movieUUID)
    if err != nil {
        log.Fatalf("Error getting movie details: %v", err)
    }
    fmt.Printf("Movie details: %+v\n", movie)
}
```

### Searching Entries

Use the `GetSearchResults` method to search for entries.

```
searchQuery := "Vagabond"
category := "book"
page := 1
searchResult, err := neodbClient.GetSearchResults(searchQuery, category, page)
if err != nil {
log.Fatalf("Error searching entries: %v", err)
}
fmt.Printf("Search results: %+v\n", searchResult)
```

### Getting Book Details

Use the `GetBook` method to retrieve book details.

```
bookUUID := "1obNw8IqrMSS5DUllN4cyG"
book, err := neodbClient.GetBook(bookUUID)
if err != nil {
log.Fatalf("Error getting book details: %v", err)
}
fmt.Printf("Book details: %+v\n", book)
```

### Getting Movie Details

Use the `GetMovie` method to retrieve movie details.

```
movieUUID := "3BU3rgfjc2PaymhsN2rxwO"
movie, err := neodbClient.GetMovie(movieUUID)
if err != nil {
log.Fatalf("Error getting movie details: %v", err)
}
fmt.Printf("Movie details: %+v\n", movie)
```
## Schema Documentation
For detailed information about the schemas used in this project, please refer to the Schema Documentation.


## Future Plans

- Add support for more API endpoints
- Implement authorization mechanisms (e.g., OAuth2)

## Contributing

Feel free to submit issues and pull requests to help improve this project.

## License

This project is licensed under the MIT License.
