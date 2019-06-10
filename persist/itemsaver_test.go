package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1193358432",
		Type: "zhenai",
		Id:   "1193358432",
		Payload: model.Profile{
			Age:        49,
			Height:     160,
			Weight:     64,
			Income:     "3-5千",
			Gender:     "女",
			Name:       "隐形的翅膀",
			Xinzuo:     "天秤座",
			Occupation: "自由职业",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "湖南衡阳",
			Education:  "高中及以下",
			Car:        "未购车",
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// save expected item
	const index = "dating_profile"
	err = save(client, expected, index)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	autualProfile, _ := model.FromJsonObj(
		actual.Payload)
	actual.Payload = autualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
