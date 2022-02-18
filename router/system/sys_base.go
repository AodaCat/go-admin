package system

import (
	v1 "go-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (*BaseRouter) InitBaseRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	baseRouter := public.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
}
