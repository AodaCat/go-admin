package system

import (
	"go-admin/model/common/response"

	"github.com/gin-gonic/gin"
)

type SystemApiApi struct{}

// @Tage SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {object} response.Response{msg=string} "创建基础api"
// @Router /api/api/createApi [post]
func (s *SystemApiApi) CreateApi(c *gin.Context) {
	response.OkWithMessage("创建成功", c)
}

func (s *SystemApiApi) DeleteApi(c *gin.Context) {

}

func (s *SystemApiApi) GetApiList(c *gin.Context) {

}

func (s *SystemApiApi) GetApiById(c *gin.Context) {

}

func (s *SystemApiApi) UpdateApi(c *gin.Context) {

}

func (s *SystemApiApi) GetAllApis(c *gin.Context) {

}

func (s *SystemApiApi) DeleteApisByIds(c *gin.Context) {

}
