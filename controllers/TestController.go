package controllers

import (
	"fmt"
	// "github.com/kataras/iris/v12"
)

type TestController struct{
	
}

func (t TestController) TestGet() {
	fmt.Println("TestGet");
}