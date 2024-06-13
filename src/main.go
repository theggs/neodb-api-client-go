package main

import (
	"fmt"
	"log"

	client "github.com/theggs/neodb-api-client-go/client/v1" // 使用别名 client
)

func main() {
	apiBaseURL := "https://neodb.social"
	neodbClient := client.NewNeoDBClient(apiBaseURL)

	// 示例：搜索条目
	searchQuery := "浪客行"
	category := "book"
	page := 1
	searchResult, err := neodbClient.GetSearchResults(searchQuery, category, page)
	if err != nil {
		log.Fatalf("Error searching entries: %v", err)
	}
	fmt.Printf("Search results: %+v\n", searchResult)

	// 示例：获取书籍详细信息
	bookUUID := "1obNw8IqrMSS5DUllN4cyG"
	book, err := neodbClient.GetBook(bookUUID)
	if err != nil {
		log.Fatalf("Error getting book details: %v", err)
	}
	fmt.Printf("Book details: %+v\n", book)

	// 示例：获取电影详细信息
	movieUUID := "3BU3rgfjc2PaymhsN2rxwO"
	movie, err := neodbClient.GetMovie(movieUUID)
	if err != nil {
		log.Fatalf("Error getting movie details: %v", err)
	}
	fmt.Printf("Movie details: %+v\n", movie)
}
