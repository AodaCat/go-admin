package system

type Group struct {
	JwtService
	CasbinService
	InitDBService
	UserService
	MenuService
}
