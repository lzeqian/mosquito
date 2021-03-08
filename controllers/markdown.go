package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/database"
	"gpm/models"
	"gpm/service"
	"gpm/tools"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type MarkDownController struct {
	beego.Controller
}

func (c *MarkDownController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	PubInit(c.Controller, ctx, controllerName, actionName, app)
}

/**
  创建vuepress
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *MarkDownController) CreateVuePress() {
	markdown := models.Markdown{}
	data := c.Ctx.Input.RequestBody
	json.Unmarshal(data, &markdown)
	fileDir := markdown.FileDir
	fileName := markdown.FileName
	destPath := fileDir + tools.PathSeparator + fileName
	//创建目录
	err := GetFileSystem(c.Ctx).Mkdir(fileDir, fileName)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	//创建.vuepress目录
	GetFileSystem(c.Ctx).Mkdir(destPath, ".vuepress")
	//创建文件README.md
	fileErr := GetFileSystem(c.Ctx).CreateFile(destPath, "README.md")
	if fileErr != nil {
		ServeJSON(c.Controller, fileErr)
		return
	}
	GetFileSystem(c.Ctx).SaveTextFile(destPath, "README.md", "#测试", 0777)
	//创建.vuepress/config.js
	GetFileSystem(c.Ctx).CreateFile(destPath+tools.PathSeparator+".vuepress", "config.js")
	configJsTemplateByte, e := ioutil.ReadFile("files/vuepress/config.js")
	if e != nil {
		ServeJSON(c.Controller, e)
		return
	}
	configJsTemplate := strings.ReplaceAll(string(configJsTemplateByte), "${base}", "/"+fileName+"/")
	GetFileSystem(c.Ctx).SaveTextFile(destPath+tools.PathSeparator+".vuepress", "config.js", configJsTemplate, 0777)
	ServeJSON(c.Controller, "")
}
func copyRemoteToLocal(fileSystem service.FileSystem, remoteDir string, localDir string) {
	nodes, _ := fileSystem.ListDir(remoteDir, "")
	for _, nodeTmp := range nodes {
		if nodeTmp.IsDir {
			os.Mkdir(localDir+tools.PathSeparator+nodeTmp.Title, os.ModePerm)
			copyRemoteToLocal(fileSystem, remoteDir+tools.PathSeparator+nodeTmp.Title, localDir+tools.PathSeparator+nodeTmp.Title)
		} else {
			allBytes, _ := fileSystem.ReadByte(remoteDir, nodeTmp.Title)
			ioutil.WriteFile(localDir+tools.PathSeparator+nodeTmp.Title, allBytes, 0777)
		}
	}
}
func (c *MarkDownController) CancelVuePress() {
	markdown := database.VuePress{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &markdown)
	workspaceArr := GetWorkSpace(c.Ctx)
	workspace := 0
	//如果是个人空间需要在个人路径上加上用户名
	if len(workspaceArr) > 0 {
		workspace, _ = strconv.Atoi(workspaceArr[0])
	}
	markdown.Workspace = workspace
	markdown = database.SearchVuePress(markdown)
	if markdown.ID == 0 {
		ServeJSON(c.Controller, errors.New("未构建的vp项目，无需取消"))
		return
	}
	database.DeleteVuePress(markdown)
	beego.DelStaticPath(markdown.AppPath)
	ServeJSON(c.Controller, "")
}

/**
  搜索分享文件
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *MarkDownController) SearchVuePress() {
	keyword := c.GetString("keyword")
	token := c.Ctx.Input.Header("Authorization")
	clwas, err := tools.GetTokenInfo(token)
	if err != nil {
		ServeJSON(c.Controller, errors.New("token错误"))
		return
	}
	ServeJSON(c.Controller, database.SearchUserVuePress(keyword, clwas.Name))
}
func checkIfVp(fileSystem service.FileSystem, remoteDir string) bool {
	nodes, _ := fileSystem.ListDir(remoteDir, "")
	for _, nodeTmp := range nodes {
		if nodeTmp.FileName == ".vuepress" {
			return true
		}
	}
	return false
}

/**
  构建vuepress
*/
func (c *MarkDownController) BuildVuePress() {
	markdown := models.Markdown{}
	requestBody := c.Ctx.Input.RequestBody
	json.Unmarshal(requestBody, &markdown)
	fileDir := markdown.FileDir
	fileName := markdown.FileName
	user, _ := user.Current()
	homeDir := user.HomeDir
	//将当前项目copy到本地用户目录.vphome
	os.Mkdir(homeDir+tools.PathSeparator+".vphome", os.ModePerm)
	packageJsonByte, err := ioutil.ReadFile("files/vuepress/package.json")
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	err = ioutil.WriteFile(homeDir+tools.PathSeparator+".vphome"+tools.PathSeparator+"package.json", packageJsonByte, 0777)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	token := c.Ctx.Input.Header("Authorization")
	clwas, err := tools.GetTokenInfo(token)
	workspaceArr := GetWorkSpace(c.Ctx)
	workspace := 0
	//如果是个人空间需要在个人路径上加上用户名
	if len(workspaceArr) > 0 {
		workspace, _ = strconv.Atoi(workspaceArr[0])
	}
	readFileDir := fileDir + fileName + "/.vuepress"
	readFileName := "config.js"
	configJsContentStr1, err := GetFileSystem(c.Ctx).ReadText(readFileDir, readFileName)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	reg := regexp.MustCompile(".*base: '(?P<mapping>.+)'.*")
	match := reg.FindStringSubmatch(configJsContentStr1)
	groupNames := reg.SubexpNames()
	mapping := ""
	for i, name := range groupNames {
		if name == "mapping" { // 第一个分组为空（也就是整个匹配）
			mapping = match[i]
		}
	}
	if mapping == "" {
		ServeJSON(c.Controller, errors.New("必须设置base上下文路径,请确认配置满足 base: '上下文路径',"))
		return
	}
	filePrefix := "public_"
	if workspace == 1 {
		filePrefix = clwas.Name + "_"
	}
	curFileName := filePrefix + strings.ReplaceAll(mapping, "/", "")
	//re3, _ := regexp.Compile(`[/|\\]+`)
	//curFileName := re3.ReplaceAllString(fileDir+tools.PathSeparator+fileName, "_")
	//if strings.HasPrefix(curFileName, "_") {
	//	curFileName = strings.TrimPrefix(curFileName, "_")
	//}
	//同时创建目录名称，比如 a目录下的b目录 目录名称为 a_b
	targetLocalDir := homeDir + tools.PathSeparator + ".vphome" + tools.PathSeparator + curFileName
	os.Mkdir(homeDir+tools.PathSeparator+".vphome"+tools.PathSeparator+curFileName, os.ModePerm)
	if !checkIfVp(GetFileSystem(c.Ctx), fileDir+tools.PathSeparator+fileName) {
		ServeJSON(c.Controller, errors.New("当前目录非vuepress项目无法构建"))
		return
	}
	copyRemoteToLocal(GetFileSystem(c.Ctx), fileDir+tools.PathSeparator+fileName, targetLocalDir)
	cmdfileDir := homeDir + tools.PathSeparator + ".vphome"
	destPath := targetLocalDir
	configJs := destPath + tools.PathSeparator + ".vuepress" + tools.PathSeparator + "config.js"
	//,err:=ioutil.ReadFile(configJs)
	//configJsContentStr1 := ""
	//f, err := os.OpenFile(configJs, os.O_RDONLY, 0600)
	//defer f.Close()
	//if err != nil {
	//	ServeJSON(c.Controller, err)
	//	return
	//} else {
	//	configJsContent, _ := ioutil.ReadAll(f)
	//	configJsContentStr1 = string(configJsContent)
	//	fmt.Println(configJsContentStr1)
	//}
	//if err != nil {
	//	ServeJSON(c.Controller, err)
	//	return
	//}

	realMapping := service.GetVuePressMappingV1(mapping, workspace, clwas.Name)
	//检查mapping是否已经映射
	queryVp := database.FindVuePress(realMapping)
	if queryVp.ID != 0 && (queryVp.FileDir != fileDir || queryVp.FileName != fileName || queryVp.Workspace != workspace) {
		ServeJSON(c.Controller, errors.New("映射路径"+realMapping+"已经配置,请选择其他的路径"))
		return
	}
	//将新生成映射路径写入临时配置文件,比如个人的映射路径为/admin/mydoc/

	regStr := reg.ReplaceAllString(configJsContentStr1, "	base: '"+realMapping+"',")
	ioutil.WriteFile(configJs, []byte(regStr), 0777)
	//vuepress build docs
	osType := runtime.GOOS
	cmd := ""
	param1 := ""
	cdDir := ""
	cdCom := ""
	if osType == "windows" {
		cmd = "cmd"
		param1 = "/C"
		cdDir = strings.ReplaceAll(cmdfileDir, "/", "\\")
		cdCom = "pushd"
	} else {
		cmd = "/bin/bash"
		param1 = "-c"
		cdDir = strings.ReplaceAll(cmdfileDir, "\\", "/")
		cdCom = "cd"
	}

	vuepress := database.VuePress{
		PressHome:     curFileName,
		AppPath:       realMapping,
		ShareUserName: clwas.Name,
		FileDir:       fileDir,
		FileName:      fileName,
		Workspace:     workspace,
	}
	searchVp := database.SearchVuePress(vuepress)
	if searchVp.ID == 0 {
		database.InsertVuePress(vuepress)
	} else {
		vuepress.ID = searchVp.ID
		database.UpdateVuePress(vuepress)
	}
	fmt.Println(cmd + " " + param1 + " " + cdCom + " " + cdDir + " " + "&& npm i && vuepress build " + curFileName)
	if execCommand(cmd, []string{param1, cdCom + " " + cdDir + " " + "&& npm i && vuepress build " + curFileName}) {
		beego.SetStaticPath(realMapping, destPath+tools.PathSeparator+".vuepress"+tools.PathSeparator+"dist")
		ServeJSON(c.Controller, realMapping)
	}

}
