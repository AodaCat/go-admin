package system

import (
	"fmt"
	"go-admin/config"
	"go-admin/global"
	model "go-admin/model/system"
	"go-admin/model/system/request"
	"go-admin/source/example"
	"go-admin/source/system"
	"go-admin/util"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func (*InitDBService) writeMysqlConfig(mysql config.Mysql) error {
	global.GA_CONFIG.Mysql = mysql
	cs := util.StructToMap(global.GA_CONFIG)
	for k, v := range cs {
		global.GA_VP.Set(k, v)
	}
	global.GA_VP.Set("jwt.signing-key", uuid.NewV4().String())
	return global.GA_VP.WriteConfig()
}

func (initDBService *InitDBService) initMsqlDB(conf request.InitDB) error {
	dsn := conf.MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := initDBService.createDatabase(dsn, "mysql", createSql); err != nil {
		return err
	}
	mysqlConfig := conf.ToMysqlConfig()
	if mysqlConfig.Dbname == "" {
		return nil
	}
	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConfig.Dsn(), // DSN data source name
		DefaultStringSize:         191,               // string 类型字段的默认长度
		SkipInitializeWithVersion: true,              // 根据版本自动配置
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}}); err != nil {
		return nil
	} else {
		global.GA_DB = db
	}
	if err := initDBService.initTables(); err != nil {
		global.GA_DB = nil
		return err
	}
	if err := initDBService.initMysqlData(); err != nil {
		global.GA_DB = nil
		return err
	}
	if err := initDBService.writeMysqlConfig(mysqlConfig); err != nil {
		return err
	}

	global.GA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return nil
}

func (initDBService *InitDBService) initMysqlData() error {
	return model.MysqlDataInitialize(
		system.Api,
		system.User,
		system.Casbin,
		system.BaseMenu,
		system.Authority,
		system.Dictionary,
		system.UserAuthority,
		system.DataAuthorities,
		system.AuthoritiesMenus,
		system.DictionaryDetail,
		system.ViewAuthorityMenuMysql,
		example.FileMysql,
	)
}
