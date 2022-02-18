package request

import (
	"fmt"
	"go-admin/config"
)

type InitDB struct {
	DBType   string `json:"dbType"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password"`
	DBName   string `json:"dbName" binding:"required"`
}

// MysqlEmptyDsn msyql 空数据库 建库链接
// Author SliverHorn
func (i *InitDB) MysqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}

// PgsqlEmptyDsn pgsql 空数据库 建库链接
// Author SliverHorn
func (i *InitDB) PgsqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return "host=" + i.Host + " user=" + i.UserName + " password=" + i.Password + " port=" + i.Port + " " + "sslmode=disable TimeZone=Asia/Shanghai"
}

// ToMysqlConfig 转换 config.Mysql
// Author [SliverHorn](https://github.com/SliverHorn)
func (i *InitDB) ToMysqlConfig() config.Mysql {
	return config.Mysql{
		Path:     i.Host,
		Port:     i.Port,
		Dbname:   i.DBName,
		Username: i.UserName,
		Password: i.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}
}

// ToPgsqlConfig 转换 config.Pgsql
// Author [SliverHorn](https://github.com/SliverHorn)
func (i *InitDB) ToPgsqlConfig() config.Pgsql {
	return config.Pgsql{
		Path:     i.Host,
		Port:     i.Port,
		Dbname:   i.DBName,
		Username: i.UserName,
		Password: i.Password,
		Config:   "sslmode=disable TimeZone=Asia/Shanghai",
	}
}
