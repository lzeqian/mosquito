package web

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/context"
	"github.com/chenhg5/collection"
	"gpm/models"
	"gpm/service"
	"gpm/tools"
	"strings"
)

var IgnoreList = []string{
	"/login",
	"/file/viewerFromServer",
}

/**
  检查当前用户当前路径下是否有权限操作文件系统权限
  @userName 当前用户名
*/
func checkFileSystemPerm(ctx *context.Context, userName string) error {
	requestPath := ctx.Request.URL.Path
	if requestPath == "/home/tree" {
		if !service.CheckUserAct(userName, tools.PathSeparator, service.ActListDir) {
			return errors.New("请确保当前用户拥有以下权限：" + service.ActListDir)
		} else {
			return nil
		}
	}
	var dirPath string
	if "GET" == ctx.Request.Method || "DELETE" == ctx.Request.Method {
		dirPath = ctx.Request.FormValue("fileDir")
	} else {
		mapParam := make(map[string]interface{})
		json.Unmarshal(ctx.Input.RequestBody, &mapParam)
		dirPath = mapParam["fileDir"].(string)
	}
	//获取当前路径需要验证的用户权限
	actList := service.GetPathRequirePerm(requestPath)
	if actList.Len() == 0 {
		return nil
	}
	actArray := tools.ListToArray(actList)
	if !service.CheckUserMulAct(userName, dirPath, actArray) {
		return errors.New("请确保当前用户拥有以下权限：" + strings.Join(actArray, ","))
	}
	return nil
}

var FilterUser = func(ctx *context.Context) {
	if ctx.Request.Method == "OPTIONS" || collection.Collect(IgnoreList).Contains(ctx.Request.URL.Path) {
		return
	}
	//获取请求头的授权头token，未获取到则获取token参数
	token := ctx.Request.Header["Authorization"]
	var tokenString string = ""
	if token != nil && len(token) > 0 {
		tokenString = token[0]
	}
	if tokenString == "" {
		tokenString = ctx.Request.FormValue("token")
	}
	if tokenString == "" {
		ctx.Input.RunController = nil
		result := models.Result{
			Code: 2,
			Data: "请求中缺失安全信息，无法访问",
		}
		byteJson, _ := json.Marshal(result)
		ctx.WriteString(string(byteJson))
		return
	}
	//验证token是否有效
	returnClaims, err := tools.ValidateToken(tokenString)
	if err != nil {
		ctx.Input.RunController = nil
		result := models.Result{
			Code: 3,
			Data: "token验证错误:" + err.Error(),
		}
		byteJson, _ := json.Marshal(result)
		ctx.WriteString(string(byteJson))
		return
	}
	//验证token中用户是否在当前数据库中
	userName := returnClaims.Name
	if service.GetUser(userName) == nil {
		result := models.Result{
			Code: 4,
			Data: "用户" + userName + "不存在或者被禁用",
		}
		byteJson, _ := json.Marshal(result)
		ctx.WriteString(string(byteJson))
		return
	}
	//验证用户是否有操作文件系统权限
	e := checkFileSystemPerm(ctx, userName)
	if e != nil {
		result := models.Result{
			Code: 5,
			Data: "用户" + userName + "无权限操作文件系统," + e.Error(),
		}
		byteJson, _ := json.Marshal(result)
		ctx.WriteString(string(byteJson))
		return
	}
}
