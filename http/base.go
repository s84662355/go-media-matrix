package http

import (
    "github.com/kataras/iris"
    nethttp "net/http"
)

var HttpApp = initRouter()

func initRouter() *iris.Application {
    router := iris.New()
    router.Use(func(ctx iris.Context) {
        ctx.Header("Access-Control-Allow-Origin", "*")
        ctx.Header("Access-Control-Allow-Credentials", "true")
        ctx.Header("Access-Control-Allow-Headers", "*")
        ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin, X-Requested-With, Content-Type, Accept, X-Access-Token,Token")
        ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if ctx.Method() == "OPTIONS" {
            ctx.StatusCode( nethttp.StatusOK)
            return
        }

        ctx.Next()
    })
    router.AllowMethods(iris.MethodOptions)
    return router
}
