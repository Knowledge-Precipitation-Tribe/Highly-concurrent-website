package main

import (
	"Highly-concurrent-website/common"
	"Highly-concurrent-website/fronted/middleware"
	"Highly-concurrent-website/fronted/web/controllers"
	"Highly-concurrent-website/repositories"
	"Highly-concurrent-website/services"
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	//1.创建iris 实例
	app := iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	tmplate := iris.HTML("./fronted/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	//4.设置模板
	app.StaticWeb("/public", "./fronted/web/public")
	//访问生成好的html静态文件
	app.StaticWeb("/html", "./fronted/web/htmlProductShow")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	//连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {

	}
	//sess := sessions.New(sessions.Config{
	//	Cookie:"AdminCookie",
	//	Expires:600*time.Minute,
	//})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	user := repositories.NewUserRepository("user", db)
	userService := services.NewService(user)
	userPro := mvc.New(app.Party("/user"))
	//userPro.Register(userService, ctx,sess.Start)
	userPro.Register(userService, ctx)
	userPro.Handle(new(controllers.UserController))

	//注册product控制器
	product := repositories.NewProductManager("product", db)
	productService := services.NewProductService(product)
	order := repositories.NewOrderMangerRepository("website.order", db)
	orderService := services.NewOrderService(order)
	proProduct := app.Party("/product")
	pro := mvc.New(proProduct)
	proProduct.Use(middleware.AuthConProduct)
	pro.Register(productService, orderService)
	pro.Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr("0.0.0.0:8082"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
