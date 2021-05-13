package main

import (
	"fmt"
	"testing"
	"example.com/go-http-demo/crawler_distributed/rpcsupport"
	"example.com/go-http-demo/crawler_distributed/worker"
	"example.com/go-http-demo/crawler_distributed/config"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/108906739",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "安静的雪",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
