package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
