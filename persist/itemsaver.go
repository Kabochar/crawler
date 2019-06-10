package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false)) // 是否处理集群
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item # %d: %v", itemCount, item)
			itemCount++

			err := save(client, item, index)
			if err != nil {
				log.Printf("Item Saver: err:item, %v, %v", err, item)
			}

		}
	}()
	return out, nil
}

// save content to elasticsearch DB
func save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)
	if item.Id == "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
