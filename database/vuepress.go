package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
  vuepress构建目录信息
*/
type VuePress struct {
	ID            uint64 `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	FileDir       string
	FileName      string
	Workspace     int    //0表示公共文档库，1表示个人文档库
	PressHome     string `gorm:"not null"`        //存储的目录
	AppPath       string `gorm:"unique;not null"` //虚拟目录 比如/hello/
	ShareUserName string //构建用户
}

func (_ *VuePress) InitDatabase(gdb *gorm.DB) {
}
func InsertVuePress(vp VuePress) {
	vp.ID, _ = snowFake.NextID()
	vp.CreatedAt = time.Now()
	db.Model(&VuePress{}).Create(vp)
}
func UpdateVuePress(vp VuePress) {
	db.Model(&VuePress{}).Update(vp)
}
func GetAllVuePress() []VuePress {
	var vuePress []VuePress
	db.Where("1=1").Find(&vuePress)
	return vuePress
}
func FindVuePress(appPath string) (rpress VuePress) {
	var tpress VuePress
	db.Where("app_path=?", appPath).Find(&tpress)
	return tpress
}
func DeleteVuePress(ipress VuePress) {
	var tpress VuePress
	db.Where("file_dir=?", ipress.FileDir).Where("file_name=?", ipress.FileName).Where("workspace=?", ipress.Workspace).Delete(&tpress)
}
func SearchVuePress(ipress VuePress) (rpress VuePress) {
	var tpress VuePress
	db.Where("file_dir=?", ipress.FileDir).Where("file_name=?", ipress.FileName).Where("workspace=?", ipress.Workspace).Find(&tpress)
	return tpress
}
func SearchUserVuePress(fileName string, shareUser string) (ruserLink []VuePress) {
	var vuePress []VuePress
	db.Where("(app_path like ? or file_name like ?) and share_user_name=?", "%"+fileName+"%", "%"+fileName+"%", shareUser).Find(&vuePress)
	return vuePress
}
