package response

import "github.com/kataras/iris"

const OK = 200

// 请求返回值
const (
    CodeSuccess   = 0 // 成功返回
    CodeFail      = 1
    CodeFailRetry = 2 //需要改参数重试
)

// Response 用户响应数据
type RespData struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

//请求成功，返回
func Success(data interface{}) iris.Map {
    return iris.Map{
        "code":    CodeSuccess,
        "message": "成功",
        "data":    data,
    }
}

//请求失败
func Fail(msg string) iris.Map {
    return iris.Map{
        "code":    1,
        "message": msg,
        "data":    nil,
    }
}

//失败
func FailCode(msg string, code int) iris.Map {
    return iris.Map{
        "code":    code,
        "message": msg,
        "data":    nil,
    }
}
