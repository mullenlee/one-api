package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/songquanpeng/one-api/common"
	"net/http"
)

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
