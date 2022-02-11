package middleware

import (
	"go-admin/global"
	"go-admin/model/common/response"
	"go-admin/model/system"
	"go-admin/service"
	"go-admin/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var jwtService = service.GroupApp.SystemGroup.JwtService

func JWTAuth(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	if token == "" {
		response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
		c.Abort()
		return
	}
	if jwtService.IsBlacklist(token) {
		response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		c.Abort()
		return
	}
	j := util.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == util.ErrTokenExpired {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}
	if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
		claims.ExpiresAt = time.Now().Unix() + global.GA_CONFIG.JWT.ExpiresTime
		newToken, _ := j.CreateTokenByOldToken(token, *claims)
		newClaims, _ := j.ParseToken(newToken)
		c.Header("new-token", newToken)
		c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		if global.GA_CONFIG.System.UseMultipoint {
			RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
			if err != nil {
				global.GA_LOG.Error("get redis jwt failed", zap.Error(err))
			} else { // 当之前的取成功时才进行拉黑操作
				_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
			}
			// 无论如何都要记录当前的活跃状态
			_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
		}
	}
	c.Set("claims", claims)
	c.Next()
}
