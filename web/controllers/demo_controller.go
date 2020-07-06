package controllers

import (
	"github.com/hqd888/iris-example/repositories"
	"github.com/hqd888/iris-example/services"
	"github.com/kataras/iris/mvc"
)

type DemoController struct {
}

func (c *DemoController) Get() mvc.View {
	demoRepository := repositories.NewDemoManager()
	demoService := services.NewDemoServiceManager(demoRepository)
	result := demoService.ShowDemoName()

	return mvc.View{
		Name: "demo/index.html",
		Data: result,
	}
}
