package controllers

import (
	"fmt"
	"github.com/kataras/iris/mvc"
	"iris-exam/repositories"
	"iris-exam/services"
)

type DemoController struct {
}

func (c *DemoController) Get() mvc.View {
	demoRepository := repositories.NewDemoManager()
	demoService := services.NewDemoServiceManager(demoRepository)
	result := demoService.ShowDemoName()
	fmt.Println(result)

	return mvc.View{
		Name: "demo/index.html",
		Data: result,
	}
}
