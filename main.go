package main

import (
	"go-admin/core"
	"go-admin/global"
	"go-admin/initialize"
)

func main() {

	global.GA_VP = core.Viper()
	global.GA_LOG = core.Zap()
	global.GA_DB = initialize.Gorm()
	initialize.Timer()
	initialize.DBList()
	if global.GA_DB != nil {
		initialize.RegisterTables(global.GA_DB)
		// 程序结束前关闭数据库链接
		db, _ := global.GA_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
