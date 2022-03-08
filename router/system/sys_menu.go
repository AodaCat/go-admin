package system

import (
	v1 "go-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	menuRouter := private.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("getMenu", authorityMenuApi.GetMenu)
	}
}
