package startserver

import (
	"log"

	"github.com/AUT-Daralfonoon/back-end/gateway/api"
	"github.com/gin-gonic/gin"
)

func initServer() (*gin.Engine){
	router := gin.Default()

	api.AddRoutes(&router.RouterGroup)

	return router
}

func ServerRouter(){
	r:= initServer()
	if err:= r.Run("locahlhost:8080"); err!= nil{
		log.Fatalf("server could not start")
	}
}