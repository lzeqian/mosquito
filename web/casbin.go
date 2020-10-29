package web

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/chenhg5/collection"
	"gpm/models"
	"gpm/service"
	"gpm/tools"
)

var IgnoreList = []string{
	"/file/viewerFromServer",
	"/file/download",
	"/file/transDoc",
	"/file/transPdf",
}
var FilterUser = func(ctx *context.Context) {
	if ctx.Request.Method == "OPTIONS" || collection.Collect(IgnoreList).Contains(ctx.Request.URL.Path) {
		return
	}
	if ctx.Request.URL.Path != "/login" {
		token := ctx.Request.Header["Authorization"]
		if token == nil {
			ctx.Input.RunController = nil
			result := models.Result{
				Code: 2,
				Data: "请求中缺失安全信息，无法访问",
			}
			byteJson, _ := json.Marshal(result)
			ctx.WriteString(string(byteJson))
		} else {
			//验证token是否有效
			returnClaims, err := tools.ValidateToken(token[0])
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
			}

		}
	}
}
