package main

import (
	"flag"
	"fmt"
	"example.com/go-http-demo/crawler_distributed/config"
	"example.com/go-http-demo/crawler_distributed/worker"
	"example.com/go-http-demo/crawler_distributed/rpcsupport"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
