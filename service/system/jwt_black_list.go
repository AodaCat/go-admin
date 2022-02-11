package system

import (
	"context"
	"go-admin/global"
	"go-admin/model/system"
	"time"

	"go.uber.org/zap"
)

type JwtService struct{}

func (*JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (*JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

func (*JwtService) GetRedisJWT(username string) (string, error) {
	redisJWT, err := global.GA_REDIS.Get(context.Background(), username).Result()
	return redisJWT, err
}

func (*JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
