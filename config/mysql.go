package config  
import (
    "fmt"   
)
var  MySQL = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local","root","123456","127.0.0.1",3306,"test")

 