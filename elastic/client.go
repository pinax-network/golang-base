package elastic

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/friendsofgo/errors"
	"github.com/pinax-network/golang-base/log"
	"go.uber.org/zap"
)

type Client struct {
	config *Config
	client *elasticsearch.Client
}

func NewClient(config *Config) (*Client, error) {

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.Hosts,
		Username:  config.User,
		Password:  config.Password,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to initialize elasticsearch Client: %v", err)
	}

	return &Client{config: config, client: client}, nil
}

func (c *Client) BulkIndexDocs(ctx context.Context, mustIndex bool, docs []esutil.BulkIndexerItem) error {

	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		NumWorkers: c.config.NumWorkers,
		Client:     c.client,
		Index:      c.config.Index,
		OnError: func(ctx context.Context, err error) {
			log.CriticalIfError("failed to index document", err)
		},
	})
	if err != nil {
		return errors.WithMessage(err, "failed to create new bulk indexer")
	}

	for _, doc := range docs {
		err = bulkIndexer.Add(ctx, doc)
		log.CriticalIfError("failed to add document to bulk indexer", err, zap.Any("doc", doc))

		if err != nil && mustIndex {
			return errors.WithMessage(err, "failed to bulk index document")
		}
	}

	err = bulkIndexer.Close(ctx)
	if err != nil {
		return errors.WithMessage(err, "failed to flush elastic docs")
	}

	return nil
}
