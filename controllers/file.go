package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/satori/go.uuid"
	"gpm/models"
	"gpm/service"
	"gpm/tools"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	PubInit(c.Controller, ctx, controllerName, actionName, app)
}

/**
  获取子目录结构
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *FileController) DownloadFile() {
	fileDir := this.GetString("fileDir")
	fileName := this.GetString("fileName")
	readBytes, _ := GetFileSystem(this.Ctx).ReadByte(fileDir, fileName)
	this.Ctx.Output.Header("Content-type", "application/force-download")
	this.Ctx.Output.Header("Content-Disposition", "attachment;filename="+fileName)
	this.Ctx.Output.Header("Pragma", "No-cache")
	this.Ctx.Output.Header("Cache-Control", "No-cache")
	this.Ctx.Output.Header("Expires", "0")
	this.Ctx.ResponseWriter.Write(readBytes)
}

/**
  删除file
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FileController) DeleteFile() {
	fileDir := c.GetString("fileDir")
	fileName := c.GetString("fileName")
	destPath := fileDir + tools.PathSeparator + fileName
	isDir, _ := GetFileSystem(c.Ctx).IsDir(destPath)
	if !isDir {
		err := GetFileSystem(c.Ctx).DeleteFile(fileDir, fileName)
		ServeJSON(c.Controller, err)
		return
	}
	ServeJSON(c.Controller, "")
}

/**
  上传文件
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *FileController) UploadFile() {
	fileDir := this.GetString("fileDir")
	f, h, _ := this.GetFile("myfile") //获取上传的文件
	fileName := h.Filename
	fileByte, _ := ioutil.ReadAll(f)
	err := GetFileSystem(this.Ctx).SaveByte(fileDir, fileName, fileByte, os.ModePerm)
	if err != nil {
		ServeJSON(this.Controller, err)
		return
	}
	fmt.Print(fileDir, fileName)
	ServeJSON(this.Controller, "")
}
func (this *FileController) CopyFileTo() {
	data := this.Ctx.Input.RequestBody
	paramData := make(map[string]string)
	json.Unmarshal(data, &paramData)
	//原始空间类型 0表示公共 1表示个人
	sourceWorkspace := this.Ctx.Input.Header("Workspace")
	//原始空间文件目录和文件名
	fileDir := paramData["fileDir"]
	fileName := paramData["fileName"]
	//目标空间的目录名称
	targetDir := paramData["targetDir"]
	globalFileSystem, personFileSystem := InitFileSystem()
	var sourFileSystem service.FileSystem = personFileSystem
	var targetFileSystem service.FileSystem = globalFileSystem
	token := this.Ctx.Input.Header("Authorization")
	clwas, _ := tools.GetTokenInfo(token)
	//从公共空间复制到个人空间，个人空间路径需带上用户名
	if sourceWorkspace == "0" {
		sourFileSystem = globalFileSystem
		targetFileSystem = personFileSystem
		userInfo := service.GetUser(clwas.Name)
		targetDir = tools.PathSeparator + userInfo["userFullName"].(string) + targetDir
	}

	//公共目录需要检查是否有读的权限
	if sourceWorkspace == "0" && !service.CheckUserMulAct(clwas.Name, fileDir, []string{service.ActRead}) {
		ServeJSON(this.Controller, errors.New("请确保当前用户拥有以下权限："+service.ActRead))
		return
	}

	//从个人空间复制到公共空间需要检测公共空间目录是否有权限
	if sourceWorkspace == "1" && !service.CheckUserMulAct(clwas.Name, targetDir, []string{service.ActCreateFile, service.ActWrite}) {
		ServeJSON(this.Controller, errors.New("请确保当前用户拥有以下权限："+strings.Join([]string{service.ActCreateFile, service.ActWrite}, ",")))
		return
	}
	srcByte, err := sourFileSystem.ReadByte(fileDir, fileName)
	if err != nil {
		ServeJSON(this.Controller, err)
		return
	}
	err = targetFileSystem.SaveByte(targetDir, fileName, srcByte, os.ModePerm)
	if err != nil {
		ServeJSON(this.Controller, err)
		return
	}
	ServeJSON(this.Controller, "")
}

/**
  将文件从当前位置拖动到到另外一个位置
*/
func (this *FileController) MoveFile() {
	data := this.Ctx.Input.RequestBody
	paramData := make(map[string]string)
	json.Unmarshal(data, &paramData)
	fileDir := paramData["fileDir"]
	fileName := paramData["fileName"]
	targetDir := paramData["targetDir"]
	token := this.Ctx.Input.Header("Authorization")
	workspace := this.Ctx.Input.Header("Workspace")
	clwas, _ := tools.GetTokenInfo(token)
	//个人不需要检查权限，公共文档库需要检查
	if workspace == "0" {
		if !service.CheckUserMulAct(clwas.Name, fileDir, []string{service.ActRead}) {
			ServeJSON(this.Controller, errors.New("请确保当前用户拥有以下权限："+service.ActRead))
			return
		}
		if !service.CheckUserMulAct(clwas.Name, targetDir, []string{service.ActCreateFile, service.ActWrite}) {
			ServeJSON(this.Controller, errors.New("请确保当前用户拥有以下权限："+strings.Join([]string{service.ActCreateFile, service.ActWrite}, ",")))
			return
		}
	} else {
		userInfo := service.GetUser(clwas.Name)
		targetDir = tools.PathSeparator + userInfo["userFullName"].(string) + targetDir
	}
	srcByte, err := GetFileSystem(this.Ctx).ReadByte(fileDir, fileName)
	if err != nil {
		ServeJSON(this.Controller, err)
		return
	}
	err = GetFileSystem(this.Ctx).SaveByte(targetDir, fileName, srcByte, os.ModePerm)
	if err != nil {
		ServeJSON(this.Controller, err)
		return
	}
	ServeJSON(this.Controller, "")
}

/**
  上传文件
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *FileController) UploadOfficeFile() {
	fileDir := this.GetString("fileDir")
	fileName := this.GetString("fileName")
	data := this.Ctx.Input.RequestBody
	paramData := make(map[string]string)
	json.Unmarshal(data, &paramData)
	resultJson := make(map[string]interface{})
	resultJson["error"] = 0
	//保存逻辑，保存需要检验是否是共享保存
	if paramData["url"] != "" {
		if this.Ctx.Request.Form.Get("sharing") == "1" {
			shareKeyString := this.Ctx.Request.Form.Get("shareKey")
			checkResult := CheckSharePrivileges(this.Ctx, shareKeyString)
			if checkResult.Code != 0 {
				fmt.Println(this.Ctx.Request.URL.Path + "%%%%%%%%" + strconv.Itoa(checkResult.Code))
				resultJson["error"] = 1
				this.Data["json"] = &resultJson
				this.ServeJSON()
				return
			}
		}
		url := paramData["url"]
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		if paramData["fileDir"] != "" {
			fileDir = paramData["fileDir"]
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		bufBytes := buf.Bytes()
		if len(bufBytes) > 0 {
			err1 := GetFileSystem(this.Ctx).SaveByte(fileDir, fileName, bufBytes, 0777)
			if err1 != nil {
				resultJson["error"] = 1
				this.ServeJSON()
			}
		}
	}
	this.Data["json"] = &resultJson
	this.ServeJSON()

}

/**
  获取子目录结构
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FileController) SaveFile() {
	markdown := models.Markdown{}
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &markdown)
	if err != nil {
		ServeJSON(c.Controller, err)
	}
	err1 := GetFileSystem(c.Ctx).SaveTextFile(markdown.FileDir, markdown.FileName, string(markdown.Value), 0777)
	if err1 != nil {
		ServeJSON(c.Controller, err1)
		return
	}
	ServeJSON(c.Controller, string(""))
}

/**
  获取子目录结构
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FileController) QueryFile() {
	fileDir := c.GetString("fileDir")
	fileName := c.GetString("fileName")
	text, _ := GetFileSystem(c.Ctx).ReadText(fileDir, fileName)
	ServeJSON(c.Controller, text)
}
func createFile(fileSystem service.FileSystem, markdown models.Markdown) (err error) {
	filesuffix := path.Ext(markdown.FileName)
	var rerr error
	rerr = fileSystem.CreateFile(markdown.FileDir, markdown.FileName)
	ifCopy := false
	copySrc := ""
	if ".xlsx" == filesuffix {
		ifCopy = true
		copySrc = "null.xlsx"

	} else if ".pptx" == filesuffix {
		ifCopy = true
		copySrc = "null.pptx"
	}
	if ifCopy {
		f, err := os.Open("files/" + copySrc)
		if err != nil {
			fmt.Println("read file fail", err)
			return err
		}
		defer f.Close()
		fd, err := ioutil.ReadAll(f)
		fileSystem.SaveByte(markdown.FileDir, markdown.FileName, fd, os.ModePerm)
	}
	return rerr
}

/**
  创建markdown
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FileController) CreateFile() {
	markdown := models.Markdown{}
	data := c.Ctx.Input.RequestBody
	json.Unmarshal(data, &markdown)
	exist, _ := GetFileSystem(c.Ctx).ExistFile(markdown.FileDir, markdown.FileName)
	if exist {
		ServeJSON(c.Controller, errors.New("文件已存在"))
		return
	}
	//err := fileSystem.CreateFile(markdown.FileDir, markdown.FileName)
	err := createFile(GetFileSystem(c.Ctx), markdown)

	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	ServeJSON(c.Controller, "")
}

/**
  copy文件
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FileController) CopyFile() {
	markdown := models.Markdown{}
	data := c.Ctx.Input.RequestBody
	json.Unmarshal(data, &markdown)
	readByte, err := GetFileSystem(c.Ctx).ReadByte(markdown.FileDir, markdown.FileName)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	err = GetFileSystem(c.Ctx).SaveByte(markdown.FileDir, markdown.NewFileName, readByte, os.ModePerm)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	ServeJSON(c.Controller, "")
}

/**
  创建markdown
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FileController) CreateDir() {
	markdown := models.Markdown{}
	data := c.Ctx.Input.RequestBody
	json.Unmarshal(data, &markdown)
	err := GetFileSystem(c.Ctx).Mkdir(markdown.FileDir, markdown.FileName)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	ServeJSON(c.Controller, "")
}

/**
  删除目录
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FileController) DeleteDir() {
	fileDir := c.GetString("fileDir")
	fileName := c.GetString("fileName")
	workspace := c.Ctx.Input.Header("Workspace")
	//公共文档库存在子目录不允许删除
	if workspace == "0" {
		listFileDir := ""
		if strings.HasSuffix(fileDir, tools.PathSeparator) {
			listFileDir = fileDir + fileName
		} else {
			listFileDir = fileDir + tools.PathSeparator + fileName
		}
		nodeList, _ := GetFileSystem(c.Ctx).ListDir(listFileDir, "")
		if len(nodeList) > 0 {
			ServeJSON(c.Controller, errors.New("存在多个子元素，请删除后继续"))
			return
		}
	}
	err := GetFileSystem(c.Ctx).RmDir(fileDir, fileName)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	ServeJSON(c.Controller, "")
}

/**
  创建markdown
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (c *FileController) RenameFile() {
	markdown := models.Markdown{}
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &markdown)
	exist, _ := GetFileSystem(c.Ctx).ExistFile(markdown.FileDir, markdown.NewFileName)
	if exist {
		ServeJSON(c.Controller, errors.New("文件已存在"))
		return
	}
	err = GetFileSystem(c.Ctx).Rename(markdown.FileDir, markdown.FileName, markdown.NewFileName)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	ServeJSON(c.Controller, "")
}

/**
  上传文件
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *FileController) UploadToBase64Img() {
	f, _, _ := this.GetFile("myfile") //返回文件，文件信息头，错误信息
	bytesImg, _ := ioutil.ReadAll(f)
	result := make(map[string]interface{})
	result["errno"] = 0
	result["data"] = []string{
		"data:image/png;base64," + base64.StdEncoding.EncodeToString(bytesImg),
	}
	this.Data["json"] = &result
	this.ServeJSON()
}

/**
  上传文件
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *FileController) UploadToServer() {
	_, fh, _ := this.GetFile("myfile") //返回文件，文件信息头，错误信息
	projectName := this.GetString("projectName")
	uploadDir := beego.AppConfig.String("uploadDir")
	uploadAccessAdress := beego.AppConfig.String("uploadAccessAdress")
	os.Mkdir(uploadDir+tools.PathSeparator+projectName, os.ModePerm)
	uuid := uuid.NewV4()
	cdirPath := uuid.String() + "_" + projectName + fh.Filename
	err := this.SaveToFile("myfile", uploadDir+tools.PathSeparator+projectName+tools.PathSeparator+cdirPath)
	if err != nil {
		ServeJSON(this.Controller, err)
		return
	}
	result := make(map[string]interface{})
	result["errno"] = 0
	result["data"] = []string{uploadAccessAdress + "/file/viewerFromServer?filePath=" + tools.PathSeparator + projectName + tools.PathSeparator + cdirPath}
	this.Data["json"] = &result
	this.ServeJSON()
}

/**
  上传文件
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *FileController) ViewerFromServer() {
	filePath := this.GetString("filePath")
	uploadDir := beego.AppConfig.String("uploadDir")
	readBytes, _ := ioutil.ReadFile(uploadDir + tools.PathSeparator + filePath)
	this.Ctx.Output.Header("Content-type", "application/force-download")
	this.Ctx.Output.Header("Content-Disposition", "attachment;filename=")
	this.Ctx.Output.Header("Pragma", "No-cache")
	this.Ctx.Output.Header("Cache-Control", "No-cache")
	this.Ctx.Output.Header("Expires", "0")
	this.Ctx.ResponseWriter.Write(readBytes)

}

/**
  获取子目录结构
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *FileController) TranslateDoc() {
	fileDir := this.GetString("fileDir")
	fileName := this.GetString("fileName")
	srcFileBytes, _ := GetFileSystem(this.Ctx).ReadByte(fileDir, fileName)
	destPath := os.TempDir() + tools.PathSeparator + fileName
	//拷贝文件内容到服务器中。
	ioutil.WriteFile(destPath, srcFileBytes, 0777)
	defer os.Remove(destPath)
	this.Ctx.Output.Header("Content-type", "application/force-download")
	this.Ctx.Output.Header("Content-Disposition", "attachment;filename="+fileName+".doc")
	this.Ctx.Output.Header("Pragma", "No-cache")
	this.Ctx.Output.Header("Cache-Control", "No-cache")
	this.Ctx.Output.Header("Expires", "0")
	genDestFileName := os.TempDir() + tools.PathSeparator + fileName + ".doc"
	defer os.Remove(genDestFileName)
	os := runtime.GOOS
	cmd := ""
	param1 := ""
	if os == "windows" {
		cmd = "cmd"
		param1 = "/C"
	} else {
		cmd = "/bin/bash"
		param1 = "-c"
	}
	//pandoc -f markdown -t docx ./test.md -o test.docx
	pandocPath := beego.AppConfig.String("pandocpath")
	commandParam := pandocPath + " " + destPath + " -o " + genDestFileName
	if tools.ExecCommand(cmd, []string{param1, commandParam}) {
		readBytes, _ := ioutil.ReadFile(genDestFileName)
		this.Ctx.ResponseWriter.Write(readBytes)
	}
}

/**
  获取子目录结构
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *FileController) TranslateToMarkdown() {
	_, fs, _ := this.GetFile("myfile") //返回文件，文件信息头，错误信息
	uploadFileName := fs.Filename
	uploadExt := strings.TrimLeft(path.Ext(uploadFileName), ".")
	uploadFileNamePath := os.TempDir() + tools.PathSeparator + uuid.NewV4().String()
	this.SaveToFile("myfile", uploadFileNamePath)
	genDestFileNamePath := os.TempDir() + tools.PathSeparator + uuid.NewV4().String() + ".markdown"
	defer os.Remove(uploadFileNamePath)
	defer os.Remove(genDestFileNamePath)
	os := runtime.GOOS
	cmd := ""
	param1 := ""
	if os == "windows" {
		cmd = "cmd"
		param1 = "/C"
	} else {
		cmd = "/bin/bash"
		param1 = "-c"
	}
	//pandoc -f markdown -t docx ./test.md -o test.docx
	pandocPath := beego.AppConfig.String("pandocpath")
	commandParam := pandocPath + " -f " + uploadExt + " -t markdown " + uploadFileNamePath + " -o " + genDestFileNamePath
	if tools.ExecCommand(cmd, []string{param1, commandParam}) {
		readBytes, _ := ioutil.ReadFile(genDestFileNamePath)
		ServeJSON(this.Controller, string(readBytes))
	}
}

/**
  转换文件为pdf，并输出到响应流
   :param fileDir 当前文件目录。
   :param fileName 当前文件名。
*/
func (this *FileController) TranslatePdf() {
	fileDir := this.GetString("fileDir")
	fileName := this.GetString("fileName")
	//PthSep := string(os.PathSeparator)
	srcFileBytes, _ := GetFileSystem(this.Ctx).ReadByte(fileDir, fileName)
	destPath := os.TempDir() + tools.PathSeparator + fileName
	//拷贝文件内容到服务器中。
	ioutil.WriteFile(destPath, srcFileBytes, 0777)
	defer os.Remove(destPath)
	this.Ctx.Output.Header("Content-type", "application/force-download")
	this.Ctx.Output.Header("Content-Disposition", "attachment;filename="+fileName+".doc")
	this.Ctx.Output.Header("Pragma", "No-cache")
	this.Ctx.Output.Header("Cache-Control", "No-cache")
	this.Ctx.Output.Header("Expires", "0")

	defer os.Remove(fileName)
	fileBase := strings.Split(filepath.Base(fileName), ".")[0]
	genDestFileName := os.TempDir() + tools.PathSeparator + fileBase + ".pdf"
	osType := runtime.GOOS
	cmd := ""
	param1 := ""
	libreofficePath := beego.AppConfig.String("libreofficepath")
	if osType == "windows" {
		cmd = "cmd"
		param1 = "/C"
	} else {
		cmd = "/bin/sh/"
		param1 = "-c"
	}
	//执行后是异步的，无法感知最终生成的文件，只能定时去扫描该文件是否存在
	cmdExec := exec.Command(cmd, []string{param1, libreofficePath, "--headless", "--convert-to", "pdf", "--outdir", os.TempDir(), destPath}...)
	fmt.Println(cmdExec.Args)
	cmdExec.Output()
	//尝试15s后未处理完成直接失败
	i := 1
	//最大尝试次数是30次
	maxCount := 30
	//每次休眠时间为 500ms
	unit := 500
	for {
		if i > maxCount {
			break
		}
		_, err := os.Stat(genDestFileName)
		if os.IsNotExist(err) {
			time.Sleep(time.Duration(unit) * time.Millisecond)
			i = i + 1
			continue
		}
		defer os.Remove(genDestFileName)
		readBytes, _ := ioutil.ReadFile(genDestFileName)
		if len(readBytes) <= 0 {
			time.Sleep(time.Duration(unit) * time.Millisecond)
			i = i + 1
			continue
		}
		this.Ctx.ResponseWriter.Write(readBytes)
		break
	}

}
