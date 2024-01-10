package common

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"os"
	"strings"
)

var ES *elasticsearch.Client
var EsEnabled = true
var IndexOneAPi = "one-api"

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

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	ES = client

	esInfo, err := ES.Info()
	if err != nil {
		return err
	}

	log.Printf("Error getting response: %s", esInfo)

	return err
}

func IndexExists(indexs []string) (bool, error) {
	res, err := esapi.IndicesExistsRequest{
		Index: indexs,
	}.Do(context.Background(), ES)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	return res.StatusCode == 200, err
}

func initIndex() {
	indexs := []string{IndexOneAPi}
	exists, err := IndexExists(indexs)
	if err != nil {
		log.Fatalf("Error initIndex response: %s", err)
	}

	if !exists {
		// create one-api index
	}

}

func IndexingDocs() {
	//Build the request body.
	document := struct {
		Name string
	}{
		Name: "go-elasticsearch",
	}
	data, _ := json.Marshal(document)
	//Set up the request object.
	req := esapi.IndexRequest{
		Index:      "one-api",
		DocumentID: "1",
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}
	//Perform the request with the client.
	res, err := req.Do(context.Background(), ES)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status(), "1")
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}

func GettingDocs() {
	ES.Get("my_index", "id")
}

func SearchDocs(index string, query string) {
	ES.Search(
		ES.Search.WithIndex(index),
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
