package initialize

import (
	"context"
	"go-admin/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	config := global.GA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GA_REDIS = client
	}
}
