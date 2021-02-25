package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/utils"
	_ "gpm/routers"
	"gpm/web"
	"io/ioutil"
	"os"
	"os/user"
)

func main() {
	user, _ := user.Current()
	PthSep := string(os.PathSeparator)
	homeDir := user.HomeDir
	configFile := homeDir + PthSep + ".vpcofig"
	if utils.FileExists(configFile) {
		bytes, _ := ioutil.ReadFile(configFile)
		maps := make(map[string]interface{})
		json.Unmarshal(bytes, &maps)
		for mapping := range maps {
			beego.SetStaticPath(mapping, maps[mapping].(string))
		}
	}
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/literallycanvas", "static/literallycanvas")
	beego.SetStaticPath("/pdf", "static/pdf")
	beego.SetStaticPath("/src", "static/src")
	beego.SetStaticPath("/static", "static/static")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Workspace", "Share-Key"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("*", beego.BeforeRouter, web.FilterUser)
	beego.Run()
}
