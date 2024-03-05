package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"one-api/common"
	"time"
)

func StoreLog(c *gin.Context, relayMode int, resp *http.Response, err error, textRequest GeneralOpenAIRequest) {
	if relayMode != RelayModeChatCompletions {
		return
	}

	status := true
	message := ""

	if err != nil || resp == nil {
		// access fail
		status = false
		message = err.Error()
	}

	reqId := c.Request.Header.Get("X-Request-Id")

	channelType := c.GetInt("channel")
	channelId := c.GetInt("channel_id")
	tokenId := c.GetInt("token_id")
	userId := c.GetInt("id")
	group := c.GetString("group")

	model := textRequest.Model
	maxToken := textRequest.MaxTokens
	temperature := textRequest.Temperature
	frequencyPenalty := textRequest.FrequencyPenalty
	presencePenalty := textRequest.PresencePenalty
	now := time.Now()

	common.IndexingDocs(reqId, status, message, channelType, channelId, tokenId, userId, group,
		model, maxToken, temperature, frequencyPenalty, presencePenalty, now)

}

func ChatStatistics(c *gin.Context) {
	// 解析URL参数
	channelId := c.Query("channelId") //渠道
	model := c.Query("model")         //模型

	// 检查参数是否有效，这里可以根据你的需求进行更详细的验证
	if channelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	// 使用参数查询"one-api"中的数据，这里是一个示例函数，你需要替换为实际的查询逻辑
	data, err := common.QueryOneAPI(channelId, model)
	jsonStr, _ := json.Marshal(data)
	var jsonData interface{}
	err = json.Unmarshal([]byte(jsonStr), &jsonData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query data"})
		return
	}

	// 返回查询结果给客户端
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    jsonData,
	})

}
