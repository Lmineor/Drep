package global

import (
	"github.com/drep/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	LOG *zap.Logger
	VP *viper.Viper
	CONFIG *config.Server
	REDIS  *redis.Client
)
