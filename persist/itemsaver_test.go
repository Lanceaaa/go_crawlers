package persist

import (
	"context"
	//"encoding/json"
	"testing"
	"example.com/go-http-demo/crawler/model"
	"github.com/olivere/elastic/v7"
	"example.com/go-http-demo/crawler/engine"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
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
	// 创建elastic客户端
	client, err := elastic.NewClient(
		// 必须在docker中关闭嗅探
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = save(client, expected, index)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	/*var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}*/
}