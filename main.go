package main

import (
	"fmt"
	"github.com/drep/core"
	"github.com/drep/core/initial"
	"github.com/drep/global"
	"github.com/drep/server"
)

func main() {
	global.VP = initial.Viper() // 初始化Viper
	global.LOG = initial.Zap()  // 初始化zap日志库
	global.DB = core.DB()       // gorm连接数据库
	server.Run()
	fmt.Println("init project")
}
