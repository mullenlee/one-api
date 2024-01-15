package common

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
	"time"
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

func IndexingDocs(reqId string, status bool, message string, channelType int, channelId int,
	tokenId int, userId int, group string, model string, maxToken int, temperature float64,
	frequencyPenalty float64, presencePenalty float64, now time.Time) {
	//Build the request body.
	document := struct {
		ReqId            string
		Status           bool
		Message          string
		ChannelType      int
		ChannelId        int
		TokenId          int
		UserId           int
		Group            string
		Model            string
		MaxToken         int
		Temperature      float64
		FrequencyPenalty float64
		PresencePenalty  float64
		Now              time.Time
	}{
		ReqId:            reqId,
		Status:           status,
		Message:          message,
		ChannelType:      channelType,
		ChannelId:        channelId,
		TokenId:          tokenId,
		UserId:           userId,
		Group:            group,
		Model:            model,
		MaxToken:         maxToken,
		Temperature:      temperature,
		FrequencyPenalty: frequencyPenalty,
		PresencePenalty:  presencePenalty,
		Now:              now,
	}
	data, _ := json.Marshal(document)
	//Set up the request object.
	req := esapi.IndexRequest{
		Index:      "one-api",
		DocumentID: reqId,
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

// 示例函数，用于查询"one-api"中的数据，你需要根据实际情况实现该函数
func QueryOneAPI(group string, count string, typeParam bool) (interface{}, error) {
	// 在这里实现查询逻辑，返回查询结果或错误
	// 这里只是一个示例，返回了固定的数据
	return gin.H{"group": group, "count": count, "type": typeParam}, nil
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
