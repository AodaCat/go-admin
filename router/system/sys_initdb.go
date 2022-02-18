package system

import (
	v1 "go-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (*InitRouter) InitInitRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	initRouter := public.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)
		initRouter.POST("checkdb", dbApi.CheckDB)
	}
}
