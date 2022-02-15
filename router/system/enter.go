package system

import "github.com/gin-gonic/gin"

type Group struct {
	ApiRouter
}

func (g *Group) RegisterRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	g.ApiRouter.InitApiRouter(public, private)
}
