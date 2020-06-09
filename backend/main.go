package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	template := iris.HTML("./backend/web/views", ".html").Layout(
		"shared/layout.html").Reload(true)
	app.RegisterView(template)
	app.StaticWeb("/assets", "./backend/web/assets")
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault(
			"message", "访问错误"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	//注册控制器

	//启动服务
	_ = app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, )
}