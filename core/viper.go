package core

import (
	"flag"
	"fmt"
	"go-admin/global"
	"go-admin/util"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(util.ConfigEnv); configEnv == "" {
				config = util.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", util.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GA_CONFIG); err != nil {
		fmt.Println(err)
	}
	// root 适配性
	// 根据root位置去找到对应迁移位置,保证root路径有效
	global.GA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(global.GA_CONFIG.JWT.ExpiresTime)),
	)
	return v
}
