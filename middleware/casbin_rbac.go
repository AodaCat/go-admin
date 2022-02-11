package middleware

import (
	"go-admin/global"
	"go-admin/model/common/response"
	"go-admin/service"
	"go-admin/util"

	"github.com/gin-gonic/gin"
)

var casbinService = service.GroupApp.SystemGroup.CasbinService

func Casbin(c *gin.Context) {
	//获取用户角色
	sub := util.GetUserAuthorityId(c)
	// 获取请求的path
	obj := c.Request.URL.Path
	//获取请求方法
	act := c.Request.Method
	e := casbinService.Casbin()
	// 判断策略是否存在
	success, _ := e.Enforce(sub, obj, act)
	if global.GA_CONFIG.System.Env == "develop" || success {
		c.Next()
	} else {
		response.FailWithDetailed(gin.H{}, "权限不足", c)
		c.Abort()
		return
	}
}
