package main

import (
	"example.com/go-http-demo/crawler_distributed/rpcsupport"
	"example.com/go-http-demo/crawler_distributed/persist"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"example.com/go-http-demo/crawler_distributed/config"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(serverRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.NewClient(host, &persist.ItemSaverService{
		Client: client,
		Index: index,
	})
}