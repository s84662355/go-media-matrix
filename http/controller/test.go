package controller

import (
      "github.com/kataras/iris"
)

var Test = new(test)

type test struct{

}

func (l *test) GetTest(ctx iris.Context){
   ctx.Writef("Hellofdgdfgdfg from %s", ctx.Path())
}