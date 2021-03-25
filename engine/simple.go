package engine

import (
	"example.com/go-http-demo/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request)  {
	// 维护了requests队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got Item: %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching: %s", r.Url)
	// 每个request去fetch获取页面结果
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}

	// 通过ParseFunc来获取最终结果
	return r.ParseFunc(body), nil
}