package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/database"
)

type ShareController struct {
	beego.Controller
}

func (c *ShareController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	PubInit(c.Controller, ctx, controllerName, actionName, app)
}

/**
  分享file
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *ShareController) ShareFile() {
	data := c.Ctx.Input.RequestBody
	link := database.UserLink{}
	json.Unmarshal(data, &link)
	database.InsertLink(link)
	ServeJSON(c.Controller, "")
}

/**
  取消分享file
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *ShareController) CancelShareFile() {
	shareKey := c.GetString("shareKey")
	database.CancelLink(shareKey)
	ServeJSON(c.Controller, "")
}

/**
  查看分享file
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *ShareController) GetShareFile() {
	shareKey := c.GetString("shareKey")
	ServeJSON(c.Controller, database.GetLink(shareKey))
}
