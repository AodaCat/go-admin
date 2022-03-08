package system

import "github.com/gin-gonic/gin"

type Group struct {
	ApiRouter
	BaseRouter
	InitRouter
	MenuRouter
}

func (g *Group) RegisterRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	g.ApiRouter.InitApiRouter(public, private)
	g.BaseRouter.InitBaseRouter(public, private)
	g.InitRouter.InitInitRouter(public, private)
	g.MenuRouter.InitMenuRouter(public, private)
}
