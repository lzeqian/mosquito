package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/database"
	"gpm/tools"
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

	token := c.Ctx.Input.Header("Authorization")
	clwas, err := tools.GetTokenInfo(token)
	if err != nil {
		ServeJSON(c.Controller, errors.New("token错误"))
		return
	}
	json.Unmarshal(data, &link)
	link.ShareUserName = clwas.User.Name
	userLink := database.FindLink(link.FileDir, link.FileName)
	if userLink.ID != 0 {
		ServeJSON(c.Controller, errors.New("已经分享过key："+userLink.ShareKey))
		return
	}
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
