package initialize

import (
	"go-admin/global"

	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}
