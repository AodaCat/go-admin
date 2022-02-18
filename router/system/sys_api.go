package system

import (
	v1 "go-admin/api/v1"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (*ApiRouter) InitApiRouter(public *gin.RouterGroup, private *gin.RouterGroup) {
	// public.Use(middleware.OperationRecord())
	privateRecord := private.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		privateRecord.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		privateRecord.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除Api
		privateRecord.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		privateRecord.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		privateRecord.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		private.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有api
		private.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
