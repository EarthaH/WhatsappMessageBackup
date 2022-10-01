package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hogsback.com/whatsapp/pkg/dto"
	"hogsback.com/whatsapp/pkg/logger"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/health/", checkHealth)
	r.POST("/message/", backupMessage)

	r.Run("localhost:8080")
}

func checkHealth(c *gin.Context) {
	logger.Info("Endpoint Hit: checkHealth")

	c.JSON(http.StatusOK, "")
}

func backupMessage(c *gin.Context) {
	logger.Info("Endpoint Hit: backupMessage")

	var message dto.Parameter
	err := c.BindJSON(&message)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	println(message)

	c.JSON(http.StatusCreated, "")
}
