package initialize

import (
	"go-admin/config"
	"go-admin/global"
	"go-admin/util"

	"go.uber.org/zap"
)

func Timer() {
	if global.GA_CONFIG.Timer.Start {
		for i := range global.GA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.GA_Timer.AddTaskByFunc("ClearDB", global.GA_CONFIG.Timer.Spec, func() {
					err := util.ClearTable(global.GA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						global.GA_LOG.Error("[Timer]ClearDB err:", zap.Error(err))
					}
				})

			}(global.GA_CONFIG.Timer.Detail[i])
		}
	}
}
