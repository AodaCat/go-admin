package system

import (
	"go-admin/global"
	"go-admin/model/system"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var User = new(user)

type user struct{}

func (u *user) TableName() string {
	return "sys_user"
}

func (u *user) Initialize() error {
	entities := []system.SysUser{
		{UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{UUID: uuid.NewV4(), Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", NickName: "QMPlusUser", HeaderImg: "https:///qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}
	if err := global.GA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (u *user) CheckDataExist() bool {
	return !errors.Is(global.GA_DB.Where("username = ?", "a303176530").First(&system.SysUser{}).Error, gorm.ErrRecordNotFound)
}