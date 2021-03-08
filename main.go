package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"gpm/database"
	_ "gpm/routers"
	"gpm/web"
	"io"
	"os"
	"os/user"
)

func initEnv() {
	user, _ := user.Current()
	PthSep := string(os.PathSeparator)
	homeDir := user.HomeDir
	vpArray := database.GetAllVuePress()
	for _, vp := range vpArray {
		beego.SetStaticPath(vp.AppPath, homeDir+PthSep+".vphome"+PthSep+vp.PressHome+PthSep+"/.vuepress/dist")
	}
	//初始化需要权限目录
	_, err := os.Stat("rbac")
	if os.IsNotExist(err) {
		os.Mkdir("rbac", os.ModePerm)
	}
	//判断目录下是否存在rbac.yml文件
	destFile := "rbac/rbac.yml"
	_, errd := os.Stat(destFile)
	if os.IsNotExist(errd) {
		source, err1 := os.Open("files/default_rbac.yml")
		if err1 != nil {
			os.Exit(1)
		}
		defer source.Close()

		destination, err2 := os.Create(destFile)
		if err2 != nil {
			os.Exit(1)
		}
		defer destination.Close()
		io.Copy(destination, source)
	}
}
func main() {
	initEnv()
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/literallycanvas", "static/literallycanvas")
	beego.SetStaticPath("/pdf", "static/pdf")
	beego.SetStaticPath("/src", "static/src")
	beego.SetStaticPath("/static", "static/static")
	beego.InsertFilter("/*", beego.AfterExec, web.FilterLast, false, false)
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
