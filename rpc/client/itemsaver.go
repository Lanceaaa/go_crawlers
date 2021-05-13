package rpcclient

import (
	"context"
	"errors"
	"log"
	"github.com/olivere/elastic/v7"
	"example.com/go-http-demo/crawler/engine"
)

func ItemSaver() (chan engine.Item, error) {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			// Call Rpc to save item
			err := Save(client, item, index)
			if err != nil {
				log.Printf("Item Saver: error saving item %v：%v", item, err)
			}
		}
	}()
	return out, nil
}

// 保存item
func Save(client *elastic.Client, item engine.Item, index string) (err error) {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	// 存储数据
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.
		Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}