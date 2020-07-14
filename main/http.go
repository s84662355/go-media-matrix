package main

import (
	"media-matrix/config"
	"media-matrix/http"
	"media-matrix/lib/log"
	"media-matrix/lib/mysql"
	"os"
	"os/signal"
	"syscall"

	"github.com/kataras/iris"
)

func init() {
	log.Logger.Init("log/http", true)

}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	mysql.ConnectMysql(config.MySQL, "default")

	http.HttpApp.Run(iris.Addr(config.LogicHTTPListenIP))

	<-done
	// logger.Sugar.Info("process exit right now")

}
