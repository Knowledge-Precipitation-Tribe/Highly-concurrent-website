package controllers

import (
	"Highly-concurrent-website/common"
	"Highly-concurrent-website/datamodels"
	"Highly-concurrent-website/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type ProductController struct {
	Ctx iris.Context
	ProductService services.IProductService
}

//以Get开头的方法代表Get请求
func (p *ProductController) GetAll() mvc.View{
	productArray, _ := p.ProductService.GetAllProduct()
	return mvc.View{
		Name:"product/view.html",
		Data:iris.Map{
			"productArray":productArray,
		},
	}
}

//修改商品信息
func (p *ProductController) PostUpdate(){
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{
		TagName:"conSite",
	})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.ProductService.UpdateProduct(product)
	if err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}