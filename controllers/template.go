package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gpm/database"
	"io/ioutil"
	"os"
	"strconv"
)

type TemplateController struct {
	beego.Controller
}

func (c *TemplateController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	PubInit(c.Controller, ctx, controllerName, actionName, app)
}

/**
获取当前模板分组
*/
func (c *TemplateController) GetTemplateGroup() {
	groups := database.GetAllFileTemplateGroup()
	length := len(groups)
	mapKeyValue := make([]map[string]string, length, length)
	for i, gp := range groups {
		tempMap := make(map[string]string)
		tempMap["label"] = gp.GroupName
		tempMap["value"] = strconv.FormatUint(gp.ID, 10)
		mapKeyValue[i] = tempMap
	}
	ServeJSON(c.Controller, mapKeyValue)
}

/**
获取当前模板列表
*/
func (c *TemplateController) GetTemplateList() {
	groupId := c.GetString("groupId")
	fileTemplates := database.GetAllFileTemplate(groupId)
	length := len(fileTemplates)
	mapKeyValue := make([]map[string]string, length, length)
	for i, gp := range fileTemplates {
		tempMap := make(map[string]string)
		tempMap["label"] = gp.TemplateName
		tempMap["value"] = strconv.FormatUint(gp.ID, 10)
		tempMap["templatePath"] = gp.TemplatePath
		mapKeyValue[i] = tempMap
	}
	ServeJSON(c.Controller, mapKeyValue)
}

/**
  从模板文件创建文件
*/
func (c *TemplateController) GenerateFileFromTemplate() {
	data := c.Ctx.Input.RequestBody
	inputMap := make(map[string]string)
	json.Unmarshal(data, &inputMap)
	fileDir := inputMap["fileDir"]
	fileName := inputMap["fileName"]
	templateId := inputMap["templateId"]
	fileTemplate := database.GetFileTemplate(templateId)
	existFile, err := fileSystem.ExistFile(fileDir, fileName)
	if existFile || err != nil {
		ServeJSON(c.Controller, errors.New("文件"+fileName+"已存在无法创建,请重新输入文件名"))
		return
	}
	//读取模板文件
	btCache, err := ioutil.ReadFile(fileTemplate.TemplatePath)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	err = fileSystem.SaveByte(fileDir, fileName, btCache, os.ModePerm)
	if err != nil {
		ServeJSON(c.Controller, err)
		return
	}
	ServeJSON(c.Controller, "")
}
