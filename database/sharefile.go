package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
  分享文件信息
*/
type UserLink struct {
	ID             uint64 `gorm:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	FileDir        string
	FileName       string
	ShareUserName  string //当前分享用户
	ShareMode      *int   //0表示仅我自己 1表示所有用户只读，2表示所有用户可编辑，3仅仅我分享的好友
	AssignUserMode *int   // 当ShareMode=3时分享的类型，0表示可查看 1表示可编辑
	ShareUser      string //当状态为ShareMode=3时指定可编辑的用户，如张三,李四
	Status         *int   //0表示禁用，1表示启用
	ShareKey       string //默认地址 http://ip/doc/Sfymd3D
	JoinKey        string //默认加入key http://ip/docJoin/Sfymd3D
	IsPublic       *int   //0表示内网使用，1表示公网使用
}

func (_ *UserLink) InitDatabase(gdb *gorm.DB) {
}
func InsertLink(link UserLink) {
	link.ID, _ = snowFake.NextID()
	link.CreatedAt = time.Now()
	status := 1
	link.Status = &status
	db.Model(&UserLink{}).Create(link)
}
func UpdateLink(link UserLink) {
	db.Model(&UserLink{}).Update(link)
}
func DeleteLink(link UserLink) {
	db.Model(&UserLink{}).Delete(link)
}
func GetLink(shareKey string) (ruserLink UserLink) {
	var userLink UserLink
	db.First(&userLink, "share_key=?", shareKey)
	return userLink
}

func FindLink(fileDir string, fileName string) (ruserLink UserLink) {
	var userLink UserLink
	db.Where("file_dir=?", fileDir).Where("file_name=?", fileName).Find(&userLink)
	return userLink
}
func SearchLink(fileName string, shareUser string) (ruserLink []UserLink) {
	var userLinks []UserLink
	db.Where("file_name like ?", "%"+fileName+"%").Where("share_user_name=?", shareUser).Find(&userLinks)
	return userLinks
}
func CancelLink(shareKey string) {
	tx := db.Begin()
	tx.Where("share_key=?", shareKey).Delete(&UserLink{})
	tx.Commit()
}
