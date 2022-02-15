package router

import (
	"go-admin/router/autocode"
	"go-admin/router/example"
	"go-admin/router/system"

	"github.com/gin-gonic/gin"
)

type Group struct {
	System   system.Group
	Example  example.Group
	Autocode autocode.Group
}

var RouterApp = new(Group)

func (g *Group) RegisterRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	g.System.RegisterRouter(public, private)
	g.Example.RegisterRouter(public, private)
	g.Autocode.RegisterRouter(public, private)
}
