package utils

import (
	"github.com/opensearch-project/opensearch-go"
	"log"
	"sync"
)

var (
	clientInstance *opensearch.Client
	once           sync.Once
)

func InitClient() *opensearch.Client {
	once.Do(func() {
		config := opensearch.Config{
			Addresses: []string{"http://localhost:9200"},
		}
		client, err := opensearch.NewClient(config)
		if err != nil {
			log.Fatal(err)
		}
		clientInstance = client
	})
	return clientInstance
}

func GetClient() *opensearch.Client {
	if clientInstance == nil {
		return InitClient()
	}
	return clientInstance
}
