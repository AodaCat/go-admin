package main

import (
	"context"
	"go-admin/core"
	"go-admin/global"
	"go-admin/initialize"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	global.GA_VP = core.Viper()
	global.GA_LOG = core.Zap()
	global.GA_DB = initialize.Gorm()

	global.GA_LOG.Info("go admin init")
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.Writer.WriteString("<h1 style='font-size:250'>💩</h1>")
	})
	srv := &http.Server{
		Addr:         ":9001",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.GA_LOG.Error("[GIN]Starting err:", zap.Error(err))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.GA_LOG.Info("[GIN]Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.GA_LOG.Error("[GIN]Forced to shutdown err:", zap.Error(err))
	}
	global.GA_LOG.Info("[GIN]Has been shutdown")
}
