package wrappers

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"saver/domain"
)

type Config struct {
	Host     string
	Username string
	Password string
	Index    string
}

type ElasticWrapper struct {
	client *elastic.Client
	index  string
}

const mapping = `
{
    "settings": {
        "index": {
            "number_of_shards": 2,
            "number_of_replicas": 1
        }
    },
    "mappings": {
        "properties": {
            "url": {
                "type": "keyword"
            },
            "title": {
                "type": "text",
            }
        }
    }
}
`

func NewElasticWrapper(cfg Config) *ElasticWrapper {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetBasicAuth(cfg.Username, cfg.Password))
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(cfg.Host).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(cfg.Host)
	if err != nil {
		// Handle error
		logrus.Fatalf("error %+v", err.Error())
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(cfg.Index).Do(ctx)
	if err != nil {
		logrus.Fatalf("error %+v", err.Error())
	}

	if !exists {
		// Create a new index.
		_, err := client.CreateIndex(cfg.Index).BodyString(mapping).Do(ctx)
		fmt.Println("Create")
		if err != nil {
			// Handle error
			logrus.Fatalf("error %+v", err.Error())
		}
	}
	return &ElasticWrapper{client, cfg.Index}
}

func (e *ElasticWrapper) Save(title, url string) {
	info := domain.Info{Url: url, Title: title}
	ctx := context.Background()

	_, err := e.client.Index().
		Index(e.index).
		BodyJson(info).
		Do(ctx)
	if err != nil {
		return
	}
}
