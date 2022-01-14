package global

import (
	"go-admin/config"
	"sync"

	"github.com/go-redis/redis"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GA_VP      *viper.Viper
	GA_CONFIG  config.Server
	GA_DB      *gorm.DB
	GA_DB_LIST map[string]*gorm.DB
	GA_REDIS   *redis.Client
	GA_LOG     *zap.Logger
	BlackCache local_cache.Cache
	lock       sync.RWMutex
)
