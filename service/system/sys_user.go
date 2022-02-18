package system

import (
	"fmt"
	"go-admin/global"
	"go-admin/model/system"
	"go-admin/util"
)

type UserService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	if nil == global.GA_DB {
		return fmt.Errorf("db not init"), nil
	}

	var user system.SysUser
	u.Password = util.MD5V([]byte(u.Password))
	err = global.GA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return err, &user
}
