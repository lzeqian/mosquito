package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/database"
	"gpm/models"
	"gpm/service"
	"gpm/tools"
	"strings"
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

const FileSystemKey = "fileSystem"

func InitFileSystem() (globalFileSystem service.FileSystem, personFileSystem service.FileSystem) {
	fsaf := service.FileSystemAbstractFactory{}
	fsaf.InitFactory()
	fileSystemLocal, _ := fsaf.ConstructFactory()
	fileSystemPerson, _ := fsaf.ConstructFactoryCustom("person")
	return fileSystemLocal, fileSystemPerson

}
func GetFileSystem(ctx *context.Context) service.FileSystem {
	fileSystem := ctx.Input.GetData(FileSystemKey)
	if fileSystem != nil {
		return fileSystem.(service.FileSystem)
	}
	return nil
}
func RequestFileSystem(ctx *context.Context, tp string) {
	globalFileSystem, personFileSystem := InitFileSystem()
	if "" == tp || "0" == tp {
		ctx.Input.SetData(FileSystemKey, globalFileSystem)
		return
	}
	ctx.Input.SetData(FileSystemKey, personFileSystem)
}
func PubInit(controller beego.Controller, ctx *context.Context, controllerName, actionName string, app interface{}) {
	controller.Init(ctx, controllerName, actionName, app)
	fileSystem := GetFileSystem(ctx)
	if fileSystem != nil && fileSystem.Ping() != nil {
		InitFileSystem()
	}

}
func GetAuthorization(ctx *context.Context) string {
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
func GetWorkSpace(ctx *context.Context) []string {
	workspace := ctx.Request.Header["Workspace"]
	//如果请求头中不存在，判断参数中是否存在
	if len(workspace) == 0 {
		workspaceParam := ctx.Request.FormValue("Workspace")
		if workspaceParam != "" {
			workspace = []string{workspaceParam}
		}
	}
	return workspace
}
func CheckSharePrivileges(ctx *context.Context, shareKeyString string) models.Result {
	requestPath := ctx.Request.URL.Path
	result := models.Result{Code: 2, Data: "您无权限执行该操作333"}
	if requestPath == "/file/download" {
		fmt.Println("hello")
	}
	//如果没有shareKey参数，则需要往后验证
	if len(shareKeyString) == 0 {
		return result
	}
	userLink := database.GetLink(shareKeyString)
	if userLink.ID != 0 {
		//检查权限，获取当前路径需要权限

		//仅仅允许自己访问，找到token和共享用户比较
		if userLink.ShareMode == 0 {
			authorization := GetAuthorization(ctx)
			if authorization != "" {
				myCustomClaims, _ := tools.GetTokenInfo(authorization)
				userInfo := service.GetUser(myCustomClaims.Name)
				if userLink.ShareUserName == userInfo["userName"] {
					result.Code = 0
					return result
				}
			}
		} else if userLink.ShareMode == 1 {
			actList := service.GetPathRequirePerm(requestPath)
			if actList.Len() != 0 {
				//1表示所有用户只读
				actArray := tools.ListToArray(actList)
				//判断接口中是否只需要只读权限，才允许通过
				if len(actArray) == 1 && actArray[0] == service.ActRead {
					result.Code = 0
					return result
				}
			}
		} else if userLink.ShareMode == 2 {
			//表示所有用户可编辑
			result.Code = 0
			return result
		} else if userLink.ShareMode == 3 {
			//0表示可查看
			if userLink.AssignUserMode == 0 {
				actList := service.GetPathRequirePerm(requestPath)
				if actList.Len() != 0 {
					//1表示所有用户只读
					actArray := tools.ListToArray(actList)
					//判断接口中是否只需要只读权限，才允许通过
					if len(actArray) == 1 && actArray[0] == service.ActRead {
						//判断当前用户是否已经登录并且在被分享用户中
						authorization := GetAuthorization(ctx)
						if authorization == "" {
							myCustomClaims, _ := tools.GetTokenInfo(authorization)
							userInfo := service.GetUser(myCustomClaims.Name)
							if tools.In(strings.Split(userLink.ShareUser, ","), userInfo["userName"].(string)) {
								result.Code = 0
								return result
							}
						}
						result.Code = 0
						return result
					}
				}
			} else {
				//1表示可编辑
				authorization := GetAuthorization(ctx)
				if authorization != "" {
					myCustomClaims, _ := tools.GetTokenInfo(authorization)
					userInfo := service.GetUser(myCustomClaims.Name)
					if tools.In(strings.Split(userLink.ShareUser, ","), userInfo["userName"].(string)) {
						result.Code = 0
						return result
					}
				}
			}
		}
	}
	result.Code = 1
	return result
}
