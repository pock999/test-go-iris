package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"

    "app/controllers"
)

func main() {
    app := iris.New()
    
    app.Use(iris.Compression)

    app.Logger().SetLevel("debug")

    app.Get("/pong", pong);

    app.Get("/console", console);

    app.Get("/jjson", jjson);

    testAPI := app.Party("/test")

    testMVC := mvc.New(testAPI)

    testMVC.Handle(new(controllers.TestController))

    // controllers.TestGet();

    app.Run(
        iris.Addr("localhost:8080"),
        // iris.WithOptimizations,
    )
}

func pong(ctx iris.Context) {
    ctx.WriteString("pong")
}

func console(ctx iris.Context) {
    fmt.Println(ctx);
    ctx.WriteString("console")
}

func jjson(ctx iris.Context) {
    ctx.JSON(iris.Map{
        "data": "jjson",
    })
}
