package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/tools"
)

type TreeController struct {
	beego.Controller
}
func (c *TreeController) Init(ctx *context.Context, controllerName, actionName string, app interface{}){
	c.Controller.Init(ctx,controllerName,actionName,app)
	PubInit(c.Controller,ctx,controllerName,actionName,app)
}
/**
	获取根目录结构
 */
func (c *TreeController) Get() {
	files,e:=fileSystem.ListRoot()
	if(e!=nil){
		ServeJSON(c.Controller,e)
	}
	ServeJSON(c.Controller,files)
}
/**
  获取子目录结构
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
 */
func (c *TreeController) ListSubTree() {
	fileDir := c.GetString("fileDir")
	fileName := c.GetString("fileName")
	root,_ := c.GetBool("root")
	var destPath string;
	if root {
		destPath=tools.PathSeparator
	}else{
		if(fileDir==tools.PathSeparator){
			destPath=fileDir+fileName
		}else {
			destPath = fileDir + tools.PathSeparator + fileName
		}
	}
	if(fileSystem.IsDir(destPath)){
		files,_:=fileSystem.ListDir(destPath)
		ServeJSON(c.Controller,files)
	}else {
		c.Data["json"] = &Result{}
		c.ServeJSON()
	}
}


