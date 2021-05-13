package client

import (
	"log"
	"github.com/olivere/elastic/v7"
	"example.com/go-http-demo/crawler/engine"
	"example.com/go-http-demo/crawler_distributed/rpcsupport"
	"example.com/go-http-demo/crawler_distributed/config"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %v：%v", item, err)
			}
		}
	}()
	return out, nil
}
