package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"hogsback.com/whatsapp/pkg/dao"
	"hogsback.com/whatsapp/pkg/dto"
	"hogsback.com/whatsapp/pkg/logger"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/health", checkHealth)
	r.POST("/message/", backupMessage)

	r.Run("localhost:8080")
}

func checkHealth(c *gin.Context) {
	logger.Info("Endpoint Hit: checkHealth")

	c.JSON(http.StatusOK, "")
}

func backupMessage(c *gin.Context) {
	logger.Info("Endpoint Hit: backupMessage")

	var params dto.Parameter
	err := c.BindJSON(&params)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	dao.AddToTable(params, context.TODO())

	c.JSON(http.StatusCreated, "")
}
