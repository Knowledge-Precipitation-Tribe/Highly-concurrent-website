package main

import (
	"Highly-concurrent-website/backend/web/controllers"
	"Highly-concurrent-website/common"
	"Highly-concurrent-website/repositories"
	"Highly-concurrent-website/services"
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
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
	//连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//product 注册控制器
	productRepository := repositories.NewProductManager(
		"product",
		db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	//order 注册控制器
	orderRepository := repositories.NewOrderMangerRepository("website.table", db)
	orderService := services.NewOrderService(orderRepository)
	orderParty := app.Party("/order")
	order := mvc.New(orderParty)
	order.Register(ctx, orderService)
	order.Handle(new(controllers.OrderController))

	//启动服务
	_ = app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations)
}
