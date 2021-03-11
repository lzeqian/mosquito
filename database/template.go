package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
模板信息
*/
type FileTemplateGroup struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	GroupName string //模板分组
}
type FileTemplate struct {
	ID           uint64 `gorm:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	GroupId      uint64 //模板分组
	TemplatePath string //模板路径
	TemplateName string //模板名称
	TemplateExt  string //模板拓展信息，比如参数
}

// 创建数据库自动插入周报模板分组
func (_ *FileTemplateGroup) InitDatabase(gdb *gorm.DB) {
	var initFileTemplateGroup = func(groupName string, groupId uint64) {
		var groupTemplate = FileTemplateGroup{}
		gdb.Where("group_name=?", groupName).First(&groupTemplate)
		if groupTemplate.ID == 0 {
			groupTemplate.ID = groupId
			groupTemplate.CreatedAt = time.Now()
			groupTemplate.GroupName = groupName
			gdb.Model(&groupTemplate).Create(groupTemplate)
		}
	}
	initFileTemplateGroup("工作类", 1)
	initFileTemplateGroup("学习类", 2)
}

// 创建数据库自动插入周报模板
func (_ *FileTemplate) InitDatabase(gdb *gorm.DB) {
	var initFileTemplate = func(templatePath string, templateName string, groupId uint64) {
		var fileTemplate = FileTemplate{}
		gdb.Where("template_path=?", templatePath).First(&fileTemplate)
		if fileTemplate.ID == 0 {
			fileTemplate.ID, _ = snowFake.NextID()
			fileTemplate.CreatedAt = time.Now()
			fileTemplate.TemplatePath = templatePath
			fileTemplate.TemplateName = templateName
			fileTemplate.GroupId = groupId
			gdb.Model(&fileTemplate).Create(fileTemplate)
		}
	}
	initFileTemplate("files/templates/weekreport.xlsx", "周报模板", 1)
	initFileTemplate("files/templates/weekmeeting.docx", "周会议模板", 1)
	initFileTemplate("files/templates/mdlearn.md", "markdown学习模板", 2)

}
func GetAllFileTemplateGroup() []FileTemplateGroup {
	var group []FileTemplateGroup
	db.Where(`1=1`).Find(&group)
	return group
}
func GetAllFileTemplate(groupId string) []FileTemplate {
	var fileTemplate []FileTemplate
	sql := "1=1"
	var args []interface{}
	if groupId != "" {
		sql += " and group_id=?"
		args = append(args, groupId)
	}
	db.Where(sql, args...).Find(&fileTemplate)
	return fileTemplate
}
func GetFileTemplate(id string) (rfileTemplate FileTemplate) {
	var fileTemplate FileTemplate
	db.First(&fileTemplate, "id=?", id)
	return fileTemplate
}
