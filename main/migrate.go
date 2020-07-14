package main

import (
	"os"
	"os/signal"
	"syscall"

	"media-matrix/config"
	"media-matrix/lib/mysql"
	"media-matrix/logic/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	//连接mysql
	conn := mysql.ConnectMysql(config.MySQL, "default")
	defer mysql.DisconnectMysql()

	//迁移数据表
	conn.AutoMigrate(&model.AppDevice{})
	conn.AutoMigrate(&model.Cate{})
	conn.AutoMigrate(&model.Article{})
	conn.AutoMigrate(&model.Banner{})

	<-done
}
