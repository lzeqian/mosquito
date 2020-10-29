package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"
	"gpm/models"
	"gpm/tools"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
	"runtime"
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
	err := fileSystem.Mkdir(fileDir, fileName)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	//创建.vuepress目录
	fileSystem.Mkdir(destPath, ".vuepress")
	//创建文件README.md
	fileErr := fileSystem.CreateFile(destPath, "README.md")
	if fileErr != nil {
		ServeJSON(c.Controller, fileErr)
		return
	}
	fileSystem.SaveTextFile(destPath, "README.md", "#测试", 0777)
	//创建.vuepress/config.js
	fileSystem.CreateFile(destPath+tools.PathSeparator+".vuepress", "config.js")
	fileSystem.SaveTextFile(destPath+tools.PathSeparator+".vuepress", "config.js", "module.exports = {\r\n"+
		"	title: '测试',\r\n"+
		"	description: '测试',\r\n"+
		"	base: '/"+fileName+"/',\r\n"+
		"	themeConfig: {\r\n"+
		"		sidebar: \"auto\",\r\n"+
		"		displayAllHeaders: true,\r\n"+
		"		nav: [\r\n"+
		"			{ text: 'Home', link: '/' },\r\n"+
		"			{ text: 'baidu', link: 'https://www.baidu.com' }\r\n"+
		"		]\r\n"+
		"	}\r\n"+
		"}", 0777)
	ServeJSON(c.Controller, "")
}
func copyRemoteToLocal(remoteDir string, localDir string) {
	nodes, _ := fileSystem.ListDir(remoteDir)
	for _, nodeTmp := range nodes {
		if nodeTmp.IsDir {
			os.Mkdir(localDir+tools.PathSeparator+nodeTmp.Title, os.ModePerm)
			copyRemoteToLocal(remoteDir+tools.PathSeparator+nodeTmp.Title, localDir+tools.PathSeparator+nodeTmp.Title)
		} else {
			allBytes, _ := fileSystem.ReadByte(remoteDir, nodeTmp.Title)
			ioutil.WriteFile(localDir+tools.PathSeparator+nodeTmp.Title, allBytes, 0777)
		}
	}
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
	re3, _ := regexp.Compile(`[/|\\]+`)
	curFileName := re3.ReplaceAllString(fileDir+tools.PathSeparator+fileName, "_")
	if strings.HasPrefix(curFileName, "_") {
		curFileName = strings.TrimPrefix(curFileName, "_")
	}
	//同时创建目录名称，比如 a目录下的b目录 目录名称为 a_b
	targetLocalDir := homeDir + tools.PathSeparator + ".vphome" + tools.PathSeparator + curFileName
	os.Mkdir(homeDir+tools.PathSeparator+".vphome"+tools.PathSeparator+curFileName, os.ModePerm)
	copyRemoteToLocal(fileDir+tools.PathSeparator+fileName, targetLocalDir)
	cmdfileDir := homeDir + tools.PathSeparator + ".vphome"
	destPath := targetLocalDir
	configJs := destPath + tools.PathSeparator + ".vuepress" + tools.PathSeparator + "config.js"
	//,err:=ioutil.ReadFile(configJs)
	configJsContentStr1 := ""
	f, err := os.OpenFile(configJs, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		configJsContent, _ := ioutil.ReadAll(f)
		configJsContentStr1 = string(configJsContent)
		fmt.Println(configJsContentStr1)
	}
	if err != nil {
		fmt.Println(err)
	}
	configJsContentStrSplit := strings.Split(configJsContentStr1, "= ")
	configJsonStr := configJsContentStrSplit[1]
	reg := regexp.MustCompile("([a-zA-Z]\\w*):")
	regStr := reg.ReplaceAllString(configJsonStr, `"$1":`)
	regStr = strings.ReplaceAll(regStr, "'", "\"")
	regStr = strings.ReplaceAll(regStr, "`", "\"")
	regStr = strings.Replace(regStr, `"http":`, "http:", -1)
	regStr = strings.Replace(regStr, `"https":`, "https:", -1)
	data := make(map[string]interface{})
	fmt.Println(regStr)
	jerr := json.Unmarshal([]byte(regStr), &data)
	if jerr != nil {
		ServeJSON(c.Controller, jerr)
	}
	if data["base"] == nil || data["base"] == "" {
		ServeJSON(c.Controller, errors.New("必须设置base上下文路径"))
		return
	}
	mapping := data["base"].(string)
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
		cmd = "/bin/sh/"
		param1 = "-c"
		cdDir = strings.ReplaceAll(cmdfileDir, "\\", "/")
		cdCom = "cd"
	}
	if execCommand(cmd, []string{param1, cdCom, cdDir, "&&", "vuepress", "build", curFileName}) {
		/**
		  写入配置，配置格式:
		  {
		     "/test/":"d:/abcd/ddd"
		      "/ggg/":"d:/abcd/ddd"
		  }
		*/

		configFile := homeDir + tools.PathSeparator + ".vpcofig"
		if !utils.FileExists(configFile) {
			os.Create(configFile)
		}
		bytes, _ := ioutil.ReadFile(configFile)
		maps := make(map[string]interface{})
		json.Unmarshal(bytes, &maps)
		maps[mapping] = destPath + tools.PathSeparator + ".vuepress" + tools.PathSeparator + "dist"

		resultByte, _ := json.Marshal(maps)
		ioutil.WriteFile(configFile, resultByte, 0777)

		beego.SetStaticPath(mapping, destPath+tools.PathSeparator+".vuepress"+tools.PathSeparator+"dist")
		ServeJSON(c.Controller, mapping)
	}

}
