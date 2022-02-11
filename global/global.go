package global

import (
	"go-admin/config"
	"go-admin/util/timer"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GA_VP                  *viper.Viper
	GA_CONFIG              config.Server
	GA_DB                  *gorm.DB
	GA_DB_LIST             map[string]*gorm.DB
	GA_REDIS               *redis.Client
	GA_LOG                 *zap.Logger
	GA_Timer               timer.Timer = timer.NewTimerTask()
	GA_Concurrency_Control             = &singleflight.Group{}
	BlackCache             local_cache.Cache
	lock                   sync.RWMutex
)
