package system

import "go-admin/service"

type ApiGroup struct {
	DBApi
	BaseApi
	SystemApiApi
	AuthorityMenuApi
}

var (
	initDBService = service.GroupApp.SystemGroup.InitDBService
	jwtService    = service.GroupApp.SystemGroup.JwtService
	userService   = service.GroupApp.SystemGroup.UserService
	menuService   = service.GroupApp.SystemGroup.MenuService
)
