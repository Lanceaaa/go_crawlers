package main

import (
	"example.com/go-http-demo/crawler/engine"
	"example.com/go-http-demo/crawler/zhenai/parser"
	"example.com/go-http-demo/crawler/scheduler"
	"example.com/go-http-demo/crawler/persist"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}