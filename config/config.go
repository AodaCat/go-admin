package config

type Server struct {
	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`

	// gorm
	Mysql  Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	DBList []DB  `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
