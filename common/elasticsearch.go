package common

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
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
		Addresses: []string{os.Getenv("ES_CONN_STRING")},
	}

	ES, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	esInfo, err := ES.Info()
	if err != nil {
		return err
	}

	SysLog("ES_CONNED , INFO IS " + esInfo.String())

	return err
}

func IndexingDocs() {
	document := struct {
		Name string `json:"name"`
	}{
		"go-elasticsearch",
	}
	data, _ := json.Marshal(document)
	ES.Index("my_index", bytes.NewReader(data))
}

func GettingDocs() {
	ES.Get("my_index", "id")
}

func SearchDocs() {
	query := `{ "query": { "match_all": {} } }`
	ES.Search(
		ES.Search.WithIndex("services"),
		ES.Search.WithBody(strings.NewReader(query)),
	)
}

func UpdatingDocs() {
	ES.Update("my_index", "id", strings.NewReader(`{doc: { language: "Go" }}`))
}

func DeletingDocs() {
	ES.Delete("my_index", "id")
}

func DeleteIndex() {
	ES.Indices.Delete([]string{"my_index"})
}
