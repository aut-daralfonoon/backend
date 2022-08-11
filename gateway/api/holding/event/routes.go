package event

import "github.com/gin-gonic/gin"

func AddRoutes(parent *gin.RouterGroup){
	eventPublicRoutingGroup := parent.Group("/event")

	eventPublicRoutingGroup.GET("/")
	eventPublicRoutingGroup.POST("/register")
}