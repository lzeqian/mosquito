package database

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sony/sonyflake"
)

type Table interface {
	InitDatabase(gdb *gorm.DB)
}

var snowFakeSetting = sonyflake.Settings{
	MachineID: func() (u uint16, e error) {
		return 1, nil
	},
	CheckMachineID: func(u uint16) bool {
		return true
	},
}
var snowFake = sonyflake.NewSonyflake(snowFakeSetting)

func getDb() (rdb *sql.DB) {
	dbPath := beego.AppConfig.String("db.path")
	dbMaxConns, _ := beego.AppConfig.Int("db.maxOpenConns")
	dbMaxIdleConns, _ := beego.AppConfig.Int("db.maxIdleConns")
	db, _ := sql.Open("sqlite3", dbPath)
	db.SetMaxOpenConns(dbMaxConns) //最大连接数；注意：当执行完sql，连接转移到rows对象上，如果rows不关闭，这条连接不会被放回池里，其他并发获取不到连接会被阻塞住。
	db.SetMaxIdleConns(dbMaxIdleConns)
	return db
}

var ifInitDb bool
var gdb *gorm.DB
var db = getOrmDb()

func getOrmDb() (rdb *gorm.DB) {
	if !ifInitDb {
		dbPath := beego.AppConfig.String("db.conn")
		dbDialet := beego.AppConfig.String("db.dialet")
		dbMaxConns, _ := beego.AppConfig.Int("db.maxOpenConns")
		dbMaxIdleConns, _ := beego.AppConfig.Int("db.maxIdleConns")
		db, _ := gorm.Open(dbDialet, dbPath)
		db.DB().SetMaxIdleConns(dbMaxIdleConns)
		db.DB().SetMaxOpenConns(dbMaxConns)
		gdb = db
		tableList := []Table{
			&UserLink{},
			&FileTemplateGroup{},
			&FileTemplate{},
			&VuePress{},
		}
		for _, table := range tableList {
			db.AutoMigrate(table)
			table.InitDatabase(gdb)
		}
		ifInitDb = true
	}
	return gdb
}
