package initialize

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	router := gin.Default()
	return router
}
