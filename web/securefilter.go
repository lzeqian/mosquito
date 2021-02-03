package web

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/context"
	"github.com/chenhg5/collection"
	"gpm/controllers"
	"gpm/models"
	"gpm/service"
	"gpm/tools"
	"strings"
)

var IgnoreList = []string{
	"/login",
	"/file/viewerFromServer",
}

func getDirPath(ctx *context.Context) string {
	//获取父亲目录路径
	var dirPath string
	if "GET" == ctx.Request.Method || "DELETE" == ctx.Request.Method {
		dirPath = ctx.Request.FormValue("fileDir")
	} else {
		contentType := ctx.Request.Header["Content-Type"]
		if len(contentType) > 0 && strings.HasPrefix(contentType[0], "multipart/form-data;") {
			dirPath = ctx.Request.Form.Get("fileDir")
		} else {
			mapParam := make(map[string]interface{})
			json.Unmarshal(ctx.Input.RequestBody, &mapParam)
			if mapParam["fileDir"] != nil {
				dirPath = mapParam["fileDir"].(string)
			} else {
				dirPath = ctx.Request.FormValue("fileDir")
			}

		}
	}
	return dirPath
}
func setDirPath(ctx *context.Context, insertValue string) string {
	//获取父亲目录路径
	var dirPath string
	if "GET" == ctx.Request.Method || "DELETE" == ctx.Request.Method {
		ctx.Request.Form.Set("fileDir", insertValue)
	} else {
		contentType := ctx.Request.Header["Content-Type"]
		if len(contentType) > 0 && strings.HasPrefix(contentType[0], "multipart/form-data;") {
			ctx.Request.Form.Set("fileDir", insertValue)
		} else {
			mapParam := make(map[string]interface{})
			json.Unmarshal(ctx.Input.RequestBody, &mapParam)
			mapParam["fileDir"] = insertValue
			requestBody, _ := json.Marshal(mapParam)
			ctx.Input.RequestBody = requestBody
		}
	}
	return dirPath
}
func getAuthorization(ctx *context.Context) string {
	//获取请求头的授权头token，未获取到则获取token参数
	token := ctx.Request.Header["Authorization"]
	var tokenString string = ""
	if token != nil && len(token) > 0 {
		tokenString = token[0]
	}
	if tokenString == "" {
		tokenString = ctx.Request.FormValue("token")
	}
	return tokenString
}

/**
  检查当前用户当前路径下是否有权限操作文件系统权限
  @userName 当前用户名
*/
func checkFileSystemPerm(ctx *context.Context, userName string) error {
	requestPath := ctx.Request.URL.Path
	dirPath := getDirPath(ctx)
	//初始化fileSystem，workspace确定是个人空间还是公共空间
	workspace := ctx.Request.Header["Workspace"]
	//如果请求头中不存在，判断参数中是否存在
	if len(workspace) == 0 {
		workspaceParam := ctx.Request.FormValue("Workspace")
		if workspaceParam != "" {
			workspace = []string{workspaceParam}
		}
	}
	//如果是个人空间需要在个人路径上加上用户名
	if len(workspace) > 0 && workspace[0] == "1" {
		controllers.RequestFileSystem("1")
		if dirPath != "" {
			//if requestPath == "/home/listSub" || dirPath==tools.PathSeparator {
			authorization := getAuthorization(ctx)
			if authorization != "" {
				myCustomClaims, _ := tools.GetTokenInfo(authorization)
				userInfo := service.GetUser(myCustomClaims.Name)
				//if !strings.HasSuffix(dirPath, tools.PathSeparator) {
				//	dirPath = dirPath + tools.PathSeparator
				//}

				if !strings.HasPrefix(dirPath, tools.PathSeparator+userInfo["userFullName"].(string)) {
					targetDirPath := dirPath
					//if(dirPath==tools.PathSeparator){
					//	targetDirPath="";
					//}
					setDirPath(ctx, tools.PathSeparator+userInfo["userFullName"].(string)+(targetDirPath))
					_, err := controllers.GetFileSystem().IsDir(tools.PathSeparator + userInfo["userFullName"].(string))
					if err != nil {
						controllers.GetFileSystem().Mkdir(tools.PathSeparator, userInfo["userFullName"].(string))
					}
				}

			}
			//}
		}
		//个人不需要权限控制
		return nil
	} else {
		controllers.RequestFileSystem("0")
	}
	//if requestPath == "/home/tree" {
	//	if !service.CheckUserAct(userName, tools.PathSeparator, service.ActListDir) {
	//		return errors.New("请确保当前用户拥有以下权限：" + service.ActListDir)
	//	} else {
	//		return nil
	//	}
	//}

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

/**
  检验用户是否登录过滤器
*/
var FilterUser = func(ctx *context.Context) {
	if ctx.Request.Method == "OPTIONS" || collection.Collect(IgnoreList).Contains(ctx.Request.URL.Path) {
		return
	}
	//获取请求头的授权头token，未获取到则获取token参数
	var tokenString string = getAuthorization(ctx)
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
