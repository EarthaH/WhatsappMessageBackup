package main

import (
	"hogsback.com/whatsapp/pkg/api"
	"hogsback.com/whatsapp/pkg/dao"
	"hogsback.com/whatsapp/pkg/logger"
)

func main() {
	logger.Init()
	dao.Init()
	api.HandleRequest()
}
