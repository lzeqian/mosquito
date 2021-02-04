package controllers

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/satori/go.uuid"
	"gpm/models"
	"gpm/tools"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
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
	readBytes, _ := fileSystem.ReadByte(fileDir, fileName)

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
	isDir, _ := fileSystem.IsDir(destPath)
	if !isDir {
		err := fileSystem.DeleteFile(fileDir, fileName)
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
	err := fileSystem.SaveByte(fileDir, fileName, fileByte, 0777)
	if err != nil {
		ServeJSON(this.Controller, err)
		return
	}
	fmt.Print(fileDir, fileName)
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
	if paramData["url"] != "" {
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
		err1 := fileSystem.SaveByte(fileDir, fileName, buf.Bytes(), 0777)
		if err1 != nil {
			resultJson["error"] = 1
			this.ServeJSON()
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
	err1 := fileSystem.SaveTextFile(markdown.FileDir, markdown.FileName, string(markdown.Value), 0777)
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
	text, _ := fileSystem.ReadText(fileDir, fileName)
	ServeJSON(c.Controller, text)
}
func createFile(markdown models.Markdown) (err error) {
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
	exist, _ := fileSystem.ExistFile(markdown.FileDir, markdown.FileName)
	if exist {
		ServeJSON(c.Controller, errors.New("文件已存在"))
		return
	}
	//err := fileSystem.CreateFile(markdown.FileDir, markdown.FileName)
	err := createFile(markdown)

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
	readByte, err := fileSystem.ReadByte(markdown.FileDir, markdown.FileName)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	err = fileSystem.SaveByte(markdown.FileDir, markdown.NewFileName, readByte, os.ModePerm)
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
	err := fileSystem.Mkdir(markdown.FileDir, markdown.FileName)
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
	listFileDir := ""
	if strings.HasSuffix(fileDir, tools.PathSeparator) {
		listFileDir = fileDir + fileName
	} else {
		listFileDir = fileDir + tools.PathSeparator + fileName
	}
	nodeList, _ := fileSystem.ListDir(listFileDir, "")
	if len(nodeList) > 0 {
		ServeJSON(c.Controller, errors.New("存在多个子元素，请删除后继续"))
		return
	}
	err := fileSystem.RmDir(fileDir, fileName)
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
	exist, _ := fileSystem.ExistFile(markdown.FileDir, markdown.NewFileName)
	if exist {
		ServeJSON(c.Controller, errors.New("文件已存在"))
		return
	}
	err = fileSystem.Rename(markdown.FileDir, markdown.FileName, markdown.NewFileName)
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
	this.SaveToFile("myfile", uploadDir+tools.PathSeparator+projectName+tools.PathSeparator+cdirPath)
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
	srcFileBytes, _ := fileSystem.ReadByte(fileDir, fileName)
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
		cmd = "/bin/sh/"
		param1 = "-c"
	}

	pandocPath := beego.AppConfig.String("pandocpath")
	if execCommand(cmd, []string{param1, pandocPath, destPath, "-o", genDestFileName}) {
		readBytes, _ := ioutil.ReadFile(genDestFileName)
		this.Ctx.ResponseWriter.Write(readBytes)
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
	srcFileBytes, _ := fileSystem.ReadByte(fileDir, fileName)
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

func execCommand(commandName string, params []string) bool {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	err1 := cmd.Start()
	if err1 != nil {
		io.Copy(cmd.Stderr, bytes.NewBufferString(err1.Error()))
		//fmt.Println(err)
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return true
}
