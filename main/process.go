package main
import (
    "os"
    "os/signal"
    "syscall"
    "media-matrix/process"
    "media-matrix/config"
    "media-matrix/lib/mysql"
     
)


func init(){
    
}

func main() {
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

    mysql.ConnectMysql(config.MySQL ,"default")
    
 
   
    process.Register.Start()
   //  process.test()

    <-done
   // logger.Sugar.Info("process exit right now")


}