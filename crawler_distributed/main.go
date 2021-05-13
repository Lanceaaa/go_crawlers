package main

import (
	"example.com/go-http-demo/crawler/engine"
	"example.com/go-http-demo/crawler/zhenai/parser"
	"example.com/go-http-demo/crawler/scheduler"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"

	//"example.com/go-http-demo/crawler/persist"
	itemsaver "example.com/go-http-demo/crawler_distribued/persist/client"
	"example.com/go-http-demo/crawler_distribued/config"
	worker "example.com/go-http-demo/crawler_distribued/client"
	"example.com/go-http-demo/crawler_distribued/rpcsupport"
)

var (
	itemSaverHost = flag.String("itemsaver host", "", "itemsaver host")
	workHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connecting to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
 	}

 	out := make(chan *rpc.Client)
 	go func() {
 		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
 	return out
}