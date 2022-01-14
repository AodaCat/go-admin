package config

type Server struct {
	// auto
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
}
