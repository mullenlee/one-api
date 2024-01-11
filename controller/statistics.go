package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"one-api/common"
)

func StoreLog(c *gin.Context, relayMode int, req *http.Request, resp *http.Response, err error, textRequest GeneralOpenAIRequest) {
	if relayMode != RelayModeChatCompletions {
		return
	}

	common.IndexingDocs()

}
