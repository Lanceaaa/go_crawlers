package main

import (
	"example.com/go-http-demo/crawler/engine"
	"example.com/go-http-demo/crawler/zhenai/parser"
	"example.com/go-http-demo/crawler/scheduler"
	"example.com/go-http-demo/crawler/persist"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}