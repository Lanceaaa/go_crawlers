package main

import (
	"testing"
	"example.com/go-http-demo/crawler_distributed/rpcsupport"
	"example.com/go-http-demo/crawler/engine"
	"time"
	"example.com/go-http-demo/crawler_distributed/config"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	go serverRpc(host, "test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url: "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id: "108906739",
		Payload: model.Profile{
			Age : 34,
			Height: 162,
			Weight: 57,
			Income: "3001-5000元",
			Gender: "女",
			Name: "安静的雪",
			Xinzuo: "牡羊座",
			Occupation: "人事/行政",
			Marriage: "离异",
			House: "已购房",
			Hokou: "山东菏泽",
			Education: "大学本科",
			Car: "未购车",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}