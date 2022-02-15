package initialize

import (
	"go-admin/global"
	"go-admin/middleware"
	"net/http"

	_ "go-admin/docs"
	routerGroup "go-admin/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	router := gin.Default()
	// 静态地址
	router.StaticFS(global.GA_CONFIG.Local.Path, http.Dir(global.GA_CONFIG.Local.Path))
	global.GA_LOG.Info("use middleware logger")
	// 直接放行全部跨域请求
	router.Use(middleware.Cors)
	global.GA_LOG.Info("use middleware cors")
	// swagger
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GA_LOG.Info("register swagger handler")

	// 获取路由组实例
	// 开放路由
	publicGroup := router.Group("/api")
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	// 鉴权
	privateGroup := router.Group("/api")
	privateGroup.Use(middleware.JWTAuth).Use(middleware.Casbin)
	routerGroup.RouterApp.RegisterRouter(publicGroup, privateGroup)

	// .Use(middleware.CasbinHandler())
	// 安装插件
	// InstallPlugin(publicGroup, privateGroup)
	global.GA_LOG.Info("router register success")

	return router
}
