package engine

import (
	"log"
	"example.com/go-http-demo/crawler/fetcher"
)

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching: %s", r.Url)
	// 每个request去fetch获取页面结果
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}

	// 通过ParseFunc来获取最终结果
	return r.Parser.Parse(body, r.Url), nil
}