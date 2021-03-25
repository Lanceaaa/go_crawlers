package main

import (
	"example.com/go-http-demo/crawler/engine"
	"example.com/go-http-demo/crawler/zhenai/parser"
	"example.com/go-http-demo/crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}