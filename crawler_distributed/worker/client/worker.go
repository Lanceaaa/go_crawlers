package client

import (
	//"fmt"
	"example.com/go-http-demo/crawler_distribued/config"
	"example.com/go-http-demo/crawler_distribued/rpcsupport"
	"example.com/go-http-demo/crawler_distribued/worker"
	"example.com/go-http-demo/crawler/engine"
)

func CreateProcessor(clientChan chan []*rpc.CLient) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq :=  worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <- clientChan
		err = c.Call(config2.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
