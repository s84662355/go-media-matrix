package main

import (
    "os"
    "os/signal"
    "syscall"

    _ "github.com/jinzhu/gorm/dialects/mysql"
    "media-matrix/config"
    "media-matrix/lib/mysql"
    "media-matrix/logic/model"
)

func main() {
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
    //连接mysql
    conn := mysql.ConnectMysql(config.MySQL, "default")
    defer mysql.DisconnectMysql()

    //迁移数据表
    conn.AutoMigrate(&model.AppDevice{})
        <-done
}
