package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/models"
	"gpm/service"
)

func ServeJSON(controller beego.Controller, data interface{}) {
	result := models.Result{}
	result.Code = 0
	switch data.(type) { //这里是通过i.(type)来判断是什么类型  下面的case分支匹配到了 则执行相关的分支
	case error:
		result.Code = 1
		result.Data = data.(error).Error()
	case models.Result:
		result = data.(models.Result)
	default:
		result.Data = data
	}
	controller.Data["json"] = &result
	controller.ServeJSON()
}

var fileSystem service.FileSystem
var globalFileSystem service.FileSystem
var personFileSystem service.FileSystem

func initFileSystem() {
	fsaf := service.FileSystemAbstractFactory{}
	fsaf.InitFactory()
	fileSystemLocal, _ := fsaf.ConstructFactory()
	globalFileSystem = fileSystemLocal
	fileSystemPerson, _ := fsaf.ConstructFactoryCustom("person")
	personFileSystem = fileSystemPerson
}
func GetFileSystem() service.FileSystem {
	return fileSystem
}
func RequestFileSystem(tp string) {
	if globalFileSystem == nil {
		initFileSystem()
	}
	if "" == tp || "0" == tp {
		fileSystem = globalFileSystem
		return
	}
	fileSystem = personFileSystem
}
func PubInit(controller beego.Controller, ctx *context.Context, controllerName, actionName string, app interface{}) {
	controller.Init(ctx, controllerName, actionName, app)
	if globalFileSystem == nil {
		initFileSystem()
	} else {
		if globalFileSystem.Ping() != nil {
			initFileSystem()
		}
	}
}
