package config

import "time"

// socket配置
var (
	ConnectTCPListenPort = "50000"
	LogicHTTPListenIP    = ":90"
	AppSign              = "yim hello world 2"
	Env                  = "develop"
	AppDomain            = "https://test-yim-api.yidejia.com"
	TimeZone, _          = time.LoadLocation("Asia/Shanghai")
)
