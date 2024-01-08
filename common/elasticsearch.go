package common

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"os"
	"strings"
)

var ES *elasticsearch.Client
var EsEnabled = true

// InitESClient This function is called after init()
func InitESClient() (error error) {
	if os.Getenv("ES_CONN_STRING") == "" {
		EsEnabled = false
		SysLog("ES_CONN_STRING not set, ES is not enabled")
		return nil
	}
	SysLog("ES is enabled")

	cfg := elasticsearch.Config{
		Addresses: []string{"ES_CONN_STRING"},
	}

	ES, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// use Ping() method check Elasticsearch
	_, err = ES.Ping()
	if err != nil {
		return err
	}

	return err
}

func indexingDocs() {
	document := struct {
		Name string `json:"name"`
	}{
		"go-elasticsearch",
	}
	data, _ := json.Marshal(document)
	ES.Index("my_index", bytes.NewReader(data))
}

func gettingDocs() {
	ES.Get("my_index", "id")
}

func searchDocs() {
	query := `{ "query": { "match_all": {} } }`
	ES.Search(
		ES.Search.WithIndex("my_index"),
		ES.Search.WithBody(strings.NewReader(query)),
	)
}

func updatingDocs() {
	ES.Update("my_index", "id", strings.NewReader(`{doc: { language: "Go" }}`))
}

func deletingDocs() {
	ES.Delete("my_index", "id")
}

func deleteIndex() {
	ES.Indices.Delete([]string{"my_index"})
}
