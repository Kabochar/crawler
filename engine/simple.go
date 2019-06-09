package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

// 任务执行函数
func (e SimpleEngine) Run(seeds ...Request) {
	// 建立任务队列
	var requests []Request
	// 把传入的任务添加到任务队列
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 只要任务队列不为空就一直爬取
	for len(requests) > 0 {

		request := requests[0]
		requests = requests[1:]

		// 开启 worker
		parseResult, err := worker(request)
		if err != nil {
			continue
		}

		// 把解析出的请求添加到请求队列
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v\n", item)
		}
	}
}

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
	return request.ParseFunc(content), nil
}