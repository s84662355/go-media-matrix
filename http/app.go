package http

import (
	"media-matrix/http/controller"

	"github.com/kataras/iris"
)

func init() {
	HttpApp.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	HttpApp.Get("/mypath1", controller.Test.GetTest)
	HttpApp.Get("/home/banner", controller.Index.Banner)
	HttpApp.Get("/home/cate", controller.Index.Cate)
	HttpApp.Get("/home/recommend", controller.Index.Recommend)
	HttpApp.Get("/home/article", controller.Index.Article)

	//

}
