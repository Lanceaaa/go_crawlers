package main

import (
	"example.com/go-http-demo/crawler/engine"
	"example.com/go-http-demo/crawler/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList(),
	})
}