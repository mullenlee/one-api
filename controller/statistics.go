package controller

import (
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
	group := c.Query("group")
	count := c.Query("count")
	typeParam := c.Query("type")

	// 检查参数是否有效，这里可以根据你的需求进行更详细的验证
	if group == "" || count == "" || typeParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	// 将参数转换为适当的类型，这里假设group和count是字符串，type是布尔值
	var typeParamBool bool
	if typeParam == "true" {
		typeParamBool = true
	} else if typeParam == "false" {
		typeParamBool = false
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type parameter"})
		return
	}

	// 使用参数查询"one-api"中的数据，这里是一个示例函数，你需要替换为实际的查询逻辑
	data, err := common.QueryOneAPI(group, count, typeParamBool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query data"})
		return
	}

	// 返回查询结果给客户端
	c.JSON(http.StatusOK, data)
}
