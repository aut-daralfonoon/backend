package api

import (
	"github.com/AUT-Daralfonoon/back-end/gateway/api/holding"
	"github.com/gin-gonic/gin"
)


func AddRoutes(master *gin.RouterGroup){
	r:= master.Group("/api", nil)

	holding.AddRoutes(r)
}