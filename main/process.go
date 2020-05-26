package main
import (
    "os"
    "os/signal"
    "syscall"
    "media-matrix/process"
    "media-matrix/config"
    "media-matrix/lib/mysql"
  ///  "media-matrix/lib/helper"
  ///  "fmt"
     "media-matrix/lib/log"
     
)


func init(){
     log.Logger.Init("log/process",true)

/// str , _ :=  os.Getwd()
 

}

func main() {
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

    mysql.ConnectMysql(config.MySQL ,"default")
    
 
   log.Logger.Info("384954398589")
    process.Register.Start()
   //  process.test()

    <-done
   // logger.Sugar.Info("process exit right now")


}