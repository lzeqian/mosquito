package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"gpm/service"
	"gpm/tools"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
	userInfo := make(map[string]string)
	data := c.Ctx.Input.RequestBody
	json.Unmarshal(data, &userInfo)
	userName := userInfo["userName"]
	password := userInfo["password"]
	if service.CheckUserPassword(userName, password) {
		token := tools.GenerateToken(tools.User{Id: userName, Name: userName}, 10*24*60)
		userMap := make(map[string]string)
		userMap["token"] = token
		userMap["userName"] = userName
		ServeJSON(c.Controller, userMap)
	} else {
		ServeJSON(c.Controller, errors.New("用户验证失败，请重试"))
	}
}
