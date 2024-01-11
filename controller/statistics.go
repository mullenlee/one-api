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
