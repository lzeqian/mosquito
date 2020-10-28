package tools

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"gpm/models"
)

var FilterUser = func(ctx *context.Context) {
	if ctx.Request.Method == "OPTIONS" {
		return
	}
	if ctx.Request.RequestURI != "/login" {
		token := ctx.Request.Header["Authorization"]
		if token == nil {
			ctx.Input.RunController = nil
			result := models.Result{
				Code: 2,
				Data: "请求中缺失安全信息，无法访问",
			}
			byteJson, _ := json.Marshal(result)
			ctx.WriteString(string(byteJson))
		}
	}
}
