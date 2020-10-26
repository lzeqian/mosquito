package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/utils"
	_ "gpm/routers"
	"gpm/service"
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
	beego.InsertFilter("/file/*",beego.BeforeRouter,service.FilterUser)
	beego.InsertFilter("/md/*",beego.BeforeRouter,service.FilterUser)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run(":8089")
}

