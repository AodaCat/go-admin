package system

import (
	v1 "go-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (*ApiRouter) InitApiRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	// public.Use(middleware.OperationRecord())
	apiPrivate := private.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiPrivate.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		apiPrivate.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除Api
		apiPrivate.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		apiPrivate.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		apiPrivate.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiPrivate.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有api
		apiPrivate.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
