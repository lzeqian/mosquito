package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/database"
	"gpm/tools"
	"strconv"
)

type FavoriteController struct {
	beego.Controller
}

func (c *FavoriteController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	PubInit(c.Controller, ctx, controllerName, actionName, app)
}

/**
  收藏file
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FavoriteController) CollectFile() {
	data := c.Ctx.Input.RequestBody
	fav := database.Favorite{}

	token := c.Ctx.Input.Header("Authorization")
	clwas, err := tools.GetTokenInfo(token)
	if err != nil {
		ServeJSON(c.Controller, errors.New("token错误"))
		return
	}
	json.Unmarshal(data, &fav)
	workspaceArr := GetWorkSpace(c.Ctx)
	workspace := 0
	//如果是个人空间需要在个人路径上加上用户名
	if len(workspaceArr) > 0 {
		workspace, _ = strconv.Atoi(workspaceArr[0])
	}
	fav.Workspace = workspace
	fav.UserName = clwas.User.Name
	userLink := database.SearchFavorite(fav)
	if userLink.ID != 0 {
		ServeJSON(c.Controller, errors.New("已经收藏，无需收藏"))
		return
	}
	database.InsertFavorite(fav)
	ServeJSON(c.Controller, "")
}

/**
  取消分享file
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FavoriteController) CancelFavFile() {
	id := c.GetString("id")
	database.DeleteFavoriteById(id)
	ServeJSON(c.Controller, "")
}

/**
  搜索分享文件
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FavoriteController) SearchFavFile() {
	keyword := c.GetString("keyword")
	token := c.Ctx.Input.Header("Authorization")
	clwas, err := tools.GetTokenInfo(token)
	if err != nil {
		ServeJSON(c.Controller, errors.New("token错误"))
		return
	}
	ServeJSON(c.Controller, database.SearchUserFavorite(keyword, clwas.Name))
}
