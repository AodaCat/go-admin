package main

import (
	"context"
	"fmt"
	"go-admin/core"
	"go-admin/global"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	global.GA_VP = core.Viper()
	global.GA_LOG = core.Zap()

	fmt.Println("go admin init")
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
			fmt.Println("[GIN]Starting err:", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("[GIN]Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("[GIN]Forced to shutdown err:", err)
	}
	fmt.Println("[GIN]Has been shutdown")
}
