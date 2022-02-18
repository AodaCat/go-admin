package system

import "go-admin/service"

type ApiGroup struct {
	DBApi
	BaseApi
	SystemApiApi
}

var (
	initDBService = service.GroupApp.SystemGroup.InitDBService
	jwtService    = service.GroupApp.SystemGroup.JwtService
	userService   = service.GroupApp.SystemGroup.UserService
)
