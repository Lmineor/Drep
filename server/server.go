package server

import (
	"fmt"
	"github.com/drep/core/initial"
	"github.com/drep/global"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func Run(){
	if global.DB != nil {
		db,_ :=  global.DB.DB()
		defer db.Close()
		initial.MigrateTables(global.DB)
	}
	runServer()
}


func runServer(){
	Router := initial.Routers()
	address := fmt.Sprintf("127.0.0.1:%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎%s`, address)
	global.LOG.Error(s.ListenAndServe().Error())
}


func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}