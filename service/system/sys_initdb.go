package system

import (
	"database/sql"
	"fmt"
	"go-admin/global"
	"go-admin/model/example"
	"go-admin/model/system"
	"go-admin/model/system/request"

	adapter "github.com/casbin/gorm-adapter/v3"
)

type InitDBService struct{}

func (initDBService *InitDBService) InitDB(conf request.InitDB) error {
	switch conf.DBType {
	case "mysql":
		return initDBService.initMsqlDB(conf)
	case "pgsql":
		return initDBService.initPgsqlDB(conf)
	default:
		return initDBService.initMsqlDB(conf)
	}
}

func (initDBService *InitDBService) initTables() error {
	return global.GA_DB.AutoMigrate(
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysAuthority{},
		system.JwtBlacklist{},
		system.SysDictionary{},
		system.SysAutoCodeHistory{},
		system.SysOperationRecord{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},

		adapter.CasbinRule{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
}

func (initDBService *InitDBService) createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
