package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/chenhg5/collection"
	"gpm/controllers"
	"gpm/models"
	"gpm/service"
	"gpm/tools"
	"regexp"
	"strconv"
	"strings"
)

var IgnoreList = []string{
	"/console",
	"/login",
	"/file/viewerFromServer",
	"/share/getShareFile",
	"/docs/.*",
	".*/.*\\.[css|js|png|PNG|jpg|JPG|html|xml|txt|woff|woff2|ttf|eot|svg|map]",
	"/favicon.ico",
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

/**
  检查当前用户当前路径下是否有权限操作文件系统权限
  @userName 当前用户名
*/
func checkFileSystemPerm(ctx *context.Context, userName string) error {
	requestPath := ctx.Request.URL.Path
	dirPath := getDirPath(ctx)
	//初始化fileSystem，workspace确定是个人空间还是公共空间
	workspace := controllers.GetWorkSpace(ctx)
	//如果是个人空间需要在个人路径上加上用户名
	if len(workspace) > 0 && workspace[0] == "1" {
		controllers.RequestFileSystem(ctx, "1")
		if dirPath != "" {
			//if requestPath == "/home/listSub" || dirPath==tools.PathSeparator {
			authorization := controllers.GetAuthorization(ctx)
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
					_, err := controllers.GetFileSystem(ctx).IsDir(tools.PathSeparator + userInfo["userFullName"].(string))
					if err != nil {
						controllers.GetFileSystem(ctx).Mkdir(tools.PathSeparator, userInfo["userFullName"].(string))
					}
				}

			}
			//}
		}
		//个人不需要权限控制
		return nil
	} else {
		controllers.RequestFileSystem(ctx, "0")
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

/**
  检验用户是否登录过滤器
  0表示验证通过,允许共享通过调用后端接口。
  1表示验证失败，需要告诉用户没有权限访问。
  2表示非共享接口调用，直接越过共享检测，进行其他权限校验。
*/
func checkShare(ctx *context.Context) models.Result {
	requestPath := ctx.Request.URL.Path
	result := models.Result{Code: 2, Data: "您无权限执行该操作444"}
	//获取请求头的授权头token，未获取到则获取token参数
	shareKey := ctx.Request.Header["Share-Key"]
	var shareKeyString string = ""
	if shareKey != nil && len(shareKey) > 0 {
		shareKeyString = shareKey[0]
	}
	if len(shareKeyString) == 0 {
		shareKeyParam := ctx.Request.FormValue("shareKey")
		if shareKeyParam != "" {
			shareKeyString = shareKeyParam
		}
	}
	//应该该接口在office插件启动时就需要检查状态，必须先放过到实际接口中去验证
	if requestPath == "/file/uploadOfficeFile" && len(shareKeyString) > 0 {
		ctx.Input.Params()["sharing"] = "1"
		ctx.Request.Form.Set("sharing", "1")
		ctx.Request.Form.Set("shareKey", shareKeyString)
		result.Code = 0
		return result
	}
	return controllers.CheckSharePrivileges(ctx, shareKeyString)
}

/**
  检验用户是否登录过滤器
*/
var FilterUser = func(ctx *context.Context) {
	fmt.Println("&&&&&&&&&&&&&&&" + ctx.Request.URL.Path)
	if ctx.Request.Method == "OPTIONS" || collection.Collect(IgnoreList).Contains(ctx.Request.URL.Path) {
		return
	} else {
		//考虑使用正则表达式匹配
		for _, ignore := range IgnoreList {
			re := regexp.MustCompile(ignore)
			if re.MatchString(ctx.Request.URL.Path) {
				return
			}
		}
	}
	//0表示验证通过,允许共享通过调用后端接口。
	//1表示验证失败，需要告诉用户没有权限访问。
	//2表示非共享接口调用，直接越过共享检测，进行其他权限校验。
	checkResult := checkShare(ctx)
	fmt.Println(ctx.Request.URL.Path + "----" + strconv.Itoa(checkResult.Code))
	if checkResult.Code == 0 {
		controllers.RequestFileSystem(ctx, "1")
		return
	}
	if checkResult.Code == 1 {
		checkResult.Data = "您无权限执行该操作11"
		byteJson, _ := json.Marshal(checkResult)
		ctx.WriteString(string(byteJson))
		return
	}
	//获取请求头的授权头token，未获取到则获取token参数
	var tokenString string = controllers.GetAuthorization(ctx)
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
var FilterLast = func(ctx *context.Context) {
	fileSystem := controllers.GetFileSystem(ctx)
	if fileSystem != nil {
		fileSystem.Close()
	}
}
