package system

import (
	"go-admin/global"
	"go-admin/model/common/response"
	"go-admin/model/system"
	systemRes "go-admin/model/system/response"
	"go-admin/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct{}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /api/menu/getMenu [post]
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	if menus, err := menuService.GetMenuTree(util.GetUserAuthorityId(c)); err != nil {
		global.GA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []system.SysMenu{}
		}
		response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}
