package core

import (
	"fmt"
	"go-admin/global"
	"go-admin/initialize"
	"go-admin/service/system"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {

	if global.GA_CONFIG.System.UseMultipoint {
		initialize.Redis()
	}
	if global.GA_DB != nil {
		system.LoadAll()
	}
	router := initialize.Routers()
	router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GA_CONFIG.System.Addr)
	s := initServer(address, router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	欢迎使用 https://github.com/AodaCat/go-admin
	当前版本:v0.0.1
	默认自动化文档地址:http://127.0.0.1%s/api/swagger/index.html`,
		address)
	err := s.ListenAndServe()
	if err != nil {
		global.GA_LOG.Error(err.Error())
	}
}
