package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
  收藏列表
*/
type Favorite struct {
	ID        uint64 `json:"ID,string";gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	FileDir   string
	FileName  string
	Workspace int    //0表示公共文档库，1表示个人文档库
	UserName  string //构建用户
}

func (_ *Favorite) InitDatabase(gdb *gorm.DB) {
}
func InsertFavorite(favorite Favorite) {
	favorite.ID, _ = snowFake.NextID()
	favorite.CreatedAt = time.Now()
	db.Model(&Favorite{}).Create(favorite)
}
func UpdateFavorite(favorite Favorite) {
	db.Model(&Favorite{}).Update(favorite)
}
func GetAllFavorite() []Favorite {
	var favorites []Favorite
	db.Where("1=1").Find(&favorites)
	return favorites
}
func FindFavorite(id string) (rpress Favorite) {
	var favorite Favorite
	db.Where("id=?", id).Find(&favorite)
	return favorite
}
func DeleteFavorite(ifavorite Favorite) {
	var favorite Favorite
	db.Where("file_dir=?", ifavorite.FileDir).Where("file_name=?", ifavorite.FileName).Where("workspace=?", ifavorite.Workspace).Delete(&favorite)
}
func DeleteFavoriteById(id string) {
	var favorite Favorite
	db.Where("id=?", id).Delete(&favorite)
}
func SearchFavorite(ifavorite Favorite) (rpress Favorite) {
	var favorite Favorite
	db.Where("file_dir=?", ifavorite.FileDir).Where("file_name=?", ifavorite.FileName).Where("workspace=?", ifavorite.Workspace).Find(&favorite)
	return favorite
}
func SearchUserFavorite(fileName string, shareUser string) (ruserLink []Favorite) {
	var favorite []Favorite
	db.Or(db.Where("app_path like ?", "%"+fileName+"%"), db.Where("file_name like ?", "%"+fileName+"%")).Where("user_name=?", shareUser).Find(&favorite)
	return favorite
}
