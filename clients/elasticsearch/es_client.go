package elasticsearch

import (
	"github.com/KestutisKazlauskas/go-utils/logger"
	"github.com/olivere/elastic"
	"context"
	"time"
	"fmt"
)

var (
	Client esClinetInterface = &esClient{}
)

type esClinetInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client * elastic.Client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.99.100:9200"),
	  	elastic.SetHealthcheckInterval(10*time.Second),
	  	elastic.SetSniff(false),
	  	//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
	  	//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(esClient *elastic.Client) {
	c.client = esClient
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
			Index(index).
			Type("doc").
			BodyJson(doc).
			Do(ctx)
	if err != nil {
		// Log error on the first error accurence and nowhere else
		logger.Log.Error(
			fmt.Sprintf("Error on indexing in elastic search index &s", index), err)

		return nil, err
	}

	return result, nil
}