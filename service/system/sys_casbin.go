package system

import (
	"go-admin/global"
	"sync"

	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinService struct{}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (*CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormAdapter.NewAdapterByDB(global.GA_DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.GA_CONFIG.Casbin.ModelPath, a)
	})
	syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
