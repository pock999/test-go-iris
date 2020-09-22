package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

type TestController struct{
}

func (c *TestController) Get(ctx iris.Context) {
	fmt.Println("fasfas");
	ctx.JSON(iris.Map{
        "get": "test controller",
    })
}

func (c *TestController) Post(ctx iris.Context) {
	ctx.JSON(iris.Map{
        "post": "test controller",
    })
}