package holding

import (
	"github.com/AUT-Daralfonoon/back-end/gateway/api/holding/event"
	"github.com/gin-gonic/gin"
)

func AddRoutes(parent *gin.RouterGroup){
	r := parent.Group("/holding", nil)

	event.AddRoutes(r)
}