package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/service"
)

type Result struct{
	Code int `json:"code"`
	Data interface{} `json:"data"`
}
func ServeJSON(controller beego.Controller,data interface{}) {
	result:=Result{};
	result.Data=data;
	result.Code=0;
	controller.Data["json"] = &result
	switch data.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
	case error:
		result.Code = 1;
		result.Data=data.(error).Error();
		break;
	}
	controller.ServeJSON()
}
var fileSystem service.FileSystem;
func initFileSystem(){
	fsaf := service.FileSystemAbstractFactory{}
	fsaf.InitFactory()
	fileSystemLocal, _ := fsaf.ConstructFactory()
	fileSystem=fileSystemLocal
}
func PubInit(controller beego.Controller,ctx *context.Context, controllerName, actionName string, app interface{}){
	controller.Init(ctx,controllerName,actionName,app)
	if fileSystem==nil {
		initFileSystem()
	}else{
		if(fileSystem.Ping()!=nil){
			initFileSystem()
		}
	}
}