package system

import (
	"go-admin/global"
	"go-admin/model/system"

	"go.uber.org/zap"
)

func LoadAll() {
	var data []string
	err := global.GA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
