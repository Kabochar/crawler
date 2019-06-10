package engine

import (
	"crawler/fetcher"
	"log"
)

// 工作池
func worker(request Request) (ParseResult, error) {
	// 抓取网页内容
	log.Printf("Fetching %s\n", request.Url)
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch error, Url: %s %v\n",
			request.Url, err)
		return ParseResult{}, err
	}

	return request.ParserFunc(content, request.Url), nil
}
