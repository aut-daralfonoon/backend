package api

import (
	"github.com/AUT-Daralfonoon/back-end/gateway/api/holding"
	"github.com/AUT-Daralfonoon/back-end/gateway/api/holding/event"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	EventManager *event.EventController
}

func NewGateway() *ServiceController {
	return &ServiceController{}
}

func InitServer() *gin.Engine {
	router := gin.Default()
	AddRoutes(&router.RouterGroup)

	return router
}

func AddRoutes(master *gin.RouterGroup) {
	r := master.Group("/api")

	holding.AddRoutes(r)
}
