package persist

import (
	"context"
	"errors"
	"log"
	"github.com/olivere/elastic/v7"
	"example.com/go-http-demo/crawler/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	// 创建elastic客户端
	client, err := elastic.NewClient(
		// 必须在docker中关闭嗅探
		elastic.SetSniff(false))
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

			err := save(client, item, index)
			if err != nil {
				log.Printf("Item Saver: error saving item %v：%v", item, err)
			}
		}
	}()
	return out, nil
}

// 保存item
func save(client *elastic.Client, item engine.Item, index string) (err error) {
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