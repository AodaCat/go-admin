package system

import (
	v1 "go-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (*ApiRouter) InitApiRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	// public.Use(middleware.OperationRecord())
	publicRecord := public.Group("")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		publicRecord.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		publicRecord.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除Api
		publicRecord.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		publicRecord.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		publicRecord.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		public.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有api
		public.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
