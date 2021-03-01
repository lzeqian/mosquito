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
	PressHome     string `gorm:"not null"`
	AppPath       string `gorm:"unique;not null"` //虚拟目录
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
	db.Model(&UserLink{}).Update(vp)
}
func GetAllVuePress() {
	var vuePress []VuePress
	db.Where("1=1").Find(&vuePress)
}
