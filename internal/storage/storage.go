package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"log"
	"telem.kmani/internal/models"
	"telem.kmani/internal/utils"
	"time"
)

func StoreDoc[T any](ctx context.Context, doc T) error {
	client := utils.GetClient()
	if client == nil {
		return errors.New("client is not valid")
	}

	var indexName, id string
	var body []byte
	var err error
	switch v := any(doc).(type) {

	case models.Telemetry:
		indexName = "telemetry-" + time.Now().Format("2006.01.02")
		id = fmt.Sprintf("%s-%s", v.DeviceID, v.CreatedAt)
		body, err = json.Marshal(v)

	case models.Event:
		indexName = "events-" + time.Now().Format("2006.01.02")
		id = fmt.Sprintf("%s-%s", v.DeviceID, v.Timestamp)
		body, err = json.Marshal(v)

	case models.DeviceConfig:
		indexName = "device-config"
		id = v.DeviceID
		body, err = json.Marshal(v)

	default:
		return fmt.Errorf("unsupported type: %T", doc)
	}

	if err != nil {
		return err
	}

	req := opensearchapi.IndexRequest{
		Index:      indexName,
		DocumentID: id,
		Body:       bytes.NewReader(body),
	}

	res, err := req.Do(ctx, client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("OpenSearch indexing error: %s", res.String())
	}

	return nil
}
