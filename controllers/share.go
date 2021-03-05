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
  更新分享file
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *ShareController) UpdateShareFile() {
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
	if userLink.ID == 0 {
		ServeJSON(c.Controller, errors.New("该key未分享，无法修改"))
		return
	}
	link.ID = userLink.ID
	database.UpdateLink(link)
	ServeJSON(c.Controller, "")
}

/**
  取消分享file
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *ShareController) CancelShareFile() {
	shareKey := c.GetString("preShareKey")
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

/**
  搜索分享文件
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *ShareController) SearchShareFile() {
	keyword := c.GetString("keyword")
	token := c.Ctx.Input.Header("Authorization")
	clwas, err := tools.GetTokenInfo(token)
	if err != nil {
		ServeJSON(c.Controller, errors.New("token错误"))
		return
	}
	ServeJSON(c.Controller, database.SearchLink(keyword, clwas.Name))
}

/**
  该文件是否已经分享
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *ShareController) IsShareFile() {
	fileDir := c.GetString("fileDir")
	fileName := c.GetString("fileName")
	userLink := database.FindLink(fileDir, fileName)
	ServeJSON(c.Controller, userLink)
}

/**
  查看静态HTML
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *ShareController) ShareStatic() {
	shareKey := c.Ctx.Input.Param(":shareKey")
	goServer := beego.AppConfig.String("front.goServer")
	documentServer := beego.AppConfig.String("front.documentServer")
	injectJs := "<script type=\"text/javascript\">window.goServer='" + goServer + "'</script>"
	injectJs += "<script type=\"text/javascript\">window.documentServer='" + documentServer + "'</script>"
	injectJs += "<script type=\"text/javascript\">window.shareKey='" + shareKey + "'</script>"
	c.Data["InjectScript"] = injectJs
	c.TplName = "share.html"
}
